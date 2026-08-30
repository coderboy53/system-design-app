package server

import (
	"net/http"

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
}
