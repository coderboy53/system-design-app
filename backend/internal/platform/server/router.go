package server

import (
	"net/http"

	"github.com/coderboy53/system-design-app/internal/auth"
	"github.com/coderboy53/system-design-app/internal/modules"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func (app *App) handlerRoutes() {
	h := modules.NewHandlers(app.Db)
	// temp root endpoint for now
	app.Router.GET(
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
	app.Router.GET(
		"/api/login",
		auth.LoginHandler,
	)
	app.Router.GET(
		"/api/callback",
		auth.CallbackHandler,
	)
	app.Router.POST(
		"/api/logout",
		auth.LogoutHandler,
	)
	app.Router.GET(
		"/api/module",
		h.GetModules,
	)
	app.Router.GET(
		"/api/module/:id",
		h.GetModules,
	)

		
}
