package server

import (
	"net/http"

	"github.com/coderboy53/system-design-app/internal/auth"
	"github.com/gin-gonic/gin"
)

func handlerRoutes(r *gin.Engine) {
	// temp root endpoint for now
	r.GET(
		"/",
		func(c *gin.Context) {
			c.JSON(
				http.StatusOK,
				gin.H{
					"message": "Welcome to app",
				},
			)
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
}
