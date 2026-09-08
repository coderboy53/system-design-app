package server

import (
	"database/sql"
	"net/http"
	"os"

	"github.com/coderboy53/system-design-app/internal/platform/db"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type App struct {
	Db     *sql.DB
	Router *gin.Engine
}

// initializes the gin Engine, then sets up the routes and initializes connection to DB
func StartServer() {
	app := App{}
	app.Router = gin.Default()
	app.Db = db.InitializeDB()
	redisStore := initializeRedis()
	app.Router.Use(sessions.Sessions("appsessions", redisStore))
	app.handlerRoutes()
	err := app.Router.Run("localhost:9000")
	if err != nil {
		logrus.Fatal("Failed to start gin engine with error: ", err)
	}
}

func initializeRedis() redis.Store {
	store, err := redis.NewStoreWithDB(20, "tcp", "localhost:6379", "", "", "", []byte(os.Getenv("REDIS_AUTH_KEY")))
	if err != nil {
		logrus.Fatal("Failed to create session store with error: ", err)
		return nil
	}
	store.Options(sessions.Options{
		Path:     "/",
		MaxAge:   21600,
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	})
	logrus.Info("Cache connection initiated successfully")
	return store
}
