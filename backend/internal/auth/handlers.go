package auth

import (
	"crypto/rand"
	"encoding/json"
	"io"
	"net/http"
	"net/url"

	"github.com/coderboy53/system-design-app/internal/user"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
)

var cfg *oauth2.Config = newOAuth2Config()

func LoginHandler(c *gin.Context) {
	state := rand.Text()
	session := sessions.Default(c)
	authUrl := cfg.AuthCodeURL(state)
	session.Set("state", state)
	err := session.Save()
	if err != nil {
		logrus.Error("Failed to save session with error: ", err)
		c.Status(http.StatusInternalServerError)
	}
	c.Redirect(http.StatusFound, authUrl)
}

func CallbackHandler(c *gin.Context) {
	currentSession := sessions.Default(c)
	sessionState := currentSession.Get("state")
	state := c.Query("state")
	if state != sessionState {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"message": "Invalid session state. Login again",
			},
		)
		logrus.Error("Invalid session state")
		return
	}
	code := c.Query("code")
	if code == "" {
		logrus.Error("Missing authorization code")
		c.Status(http.StatusBadRequest)
		return
	}
	token, err := cfg.Exchange(c, code)
	if err != nil {
		logrus.Error("Token exchange failed with error: ", err)
		c.Status(http.StatusBadGateway)
		return
	}
	client := cfg.Client(c, token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
	if err != nil || resp.StatusCode != 200 {
		logrus.Error("Failed to fetch user info with error: ", err)
		c.Status(http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()
	userInfo, err := io.ReadAll(resp.Body)
	if err != nil {
		logrus.Error("Failed to read response body with error: ", err)
		c.Status(http.StatusInternalServerError)
		return
	}
	var userDetails user.User
	err = json.Unmarshal(userInfo, &userDetails)
	if err != nil {
		logrus.Error("Failed to decode respones body into JSON with error: ", err)
	}
	logrus.Info(userDetails)
	currentSession.Set("userid", userDetails.Subject)
	currentSession.Set("token", token.AccessToken)
	err = currentSession.Save()
	if err != nil {
		logrus.Error("Failed to save session with error: ", err)
		c.Status(http.StatusInternalServerError)
	}
	c.Redirect(http.StatusFound, "/")
}

func LogoutHandler(c *gin.Context) {
	currentSession := sessions.Default(c)
	token, ok := currentSession.Get("token").(string)
	if !ok {
		logrus.Error("Failed to convert token value to string")
	}
	currentSession.Clear()
	err := currentSession.Save()
	if err != nil {
		logrus.Error("Failed to clear session with error: ", err)
	}
	http.PostForm("https://oauth2.googleapis.com/revoke", url.Values{"token": {token}})
	c.Redirect(http.StatusFound, "/")
}
