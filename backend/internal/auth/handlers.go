package auth

import (
	"crypto/rand"
	"io"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
)

var cfg *oauth2.Config = newOAuth2Config()

func LoginHandler(c *gin.Context) {
	state := rand.Text()
	session := sessions.Default(c)
	if session.Get("state") == nil {
		authUrl := cfg.AuthCodeURL(state)
		session.Set("state", state)
		session.Save()
		c.Redirect(http.StatusFound, authUrl)
	}
}

func CallbackHandler(c *gin.Context) {
	sessionState := sessions.Default(c).Get("state")
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
		return
	}
	token, err := cfg.Exchange(c, code)
	if err != nil {
		logrus.Error("Token exchange failed with error: ", err)
		return
	}
	client := cfg.Client(c, token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
	if err != nil {
		logrus.Error("Failed to fetch user info with error: ", err)
		return
	}
	defer resp.Body.Close()
	userInfo, err := io.ReadAll(resp.Body)
	if err != nil {
		logrus.Error("Failed to read response body with error: ", err)
		return
	}
	logrus.Info(string(userInfo))

}
