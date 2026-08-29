package db

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

func InitializeDB() *sql.DB {
	db_config := make(map[string]string)
	db_config["hostname"] = os.Getenv("DB_HOST")
	db, err := sql.Open("postgres", "hostname=localhost dbname=system_design user=kokurou password=hello@User123 sslmode=disable")
	if err != nil {
		logrus.Error("DB connection failed with error", err)
		return nil
	}
	return db
}
