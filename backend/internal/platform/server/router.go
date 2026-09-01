package server

import (
	"net/http"

	"github.com/coderboy53/system-design-app/internal/auth"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func handlerRoutes(r *gin.Engine) {
	// temp root endpoint for now
	r.GET(
		"/",
		func(c *gin.Context) {
			currentSession := sessions.Default(c)
			if currentSession.Get("userid") == nil {
				c.Status(http.StatusUnauthorized)
			} else {
				c.Status(http.StatusOK)
			}
		},
	)
	r.GET(
		"/api/login",
		auth.LoginHandler,
	)
	r.GET(
		"/api/callback",
		auth.CallbackHandler,
	)
	r.POST(
		"/api/logout",
		auth.LogoutHandler,
	)
}
