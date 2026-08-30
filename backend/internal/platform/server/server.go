package server

import (
	"github.com/coderboy53/system-design-app/internal/platform/db"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// initializes the gin Engine, then sets up the routes and initializes connection to DB
func StartServer() {
	r := gin.Default()
	handlerRoutes(r)
	db.InitializeDB()
	err := r.Run("localhost:9000")
	if err != nil {
		logrus.Fatal("Failed to start gin engine with error: ", err)
	}
}
