package server

import (
	"github.com/coderboy53/system-design-app/internal/platform/db"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// initializes the gin Engine, then sets up the routes and initializes connection to DB
func StartServer() {
	r := gin.Default()
	handlerRoutes(r)
	db.InitializeDB()
	redisStore := initializeRedis()
	r.Use(sessions.Sessions("appsessions", redisStore))
	err := r.Run("localhost:9000")
	if err != nil {
		logrus.Fatal("Failed to start gin engine with error: ", err)
	}
}

func initializeRedis() redis.Store {
	store, err := redis.NewStoreWithDB(20, "tcp", "localhost:6379", "", "", "")
	if err != nil {
		logrus.Error("Failed to create session store with error: ", err)
		return nil
	}
	return store
}
