package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

func InitializeDB() *sql.DB {
	db_config := make(map[string]string)
	db_config["hostname"] = os.Getenv("DB_HOST")
	db_config["dbname"] = os.Getenv("DB_NAME")
	db_config["user"] = os.Getenv("DB_USER")
	db_config["password"] = os.Getenv("DB_PASS")
	db, err := sql.Open("postgres", fmt.Sprintf("hostname=localhost dbname=system_design user=kokurou password=hello@User123 sslmode=disable", db_config["hostname"], db_config["dbname"], db_config["user"], db_config["password"]))
	if err != nil {
		logrus.Error("DB connection failed with error", err)
		return nil
	}
	return db
}
