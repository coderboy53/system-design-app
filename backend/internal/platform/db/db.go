package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

// initializes the connection with local postgres reading env variables
func InitializeDB() *sql.DB {
	db_config := make(map[string]string)
	db_config["host"] = os.Getenv("DB_HOST")
	db_config["dbname"] = os.Getenv("DB_NAME")
	db_config["user"] = os.Getenv("DB_USER")
	db_config["password"] = os.Getenv("DB_PASS")
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s dbname=%s user=%s password=%s sslmode=verify-full sslrootcert=./internal/platform/db/global-bundle.pem", db_config["host"], db_config["dbname"], db_config["user"], db_config["password"]))
	if err != nil {
		logrus.Fatal("DB string configuration not valid")
		return nil
	}
	if err = db.Ping(); err != nil {
		logrus.Fatal("DB connection failed with error: ", err)
		return nil
	}
	logrus.Info("DB connection initiated successfully")
	return db
}
