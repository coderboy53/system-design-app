package modules

import (
	"database/sql"

	"github.com/sirupsen/logrus"
)

// tmp functions for populating data
func AddModules(db *sql.DB, module Module) error {
	_, err := db.Exec("INSERT INTO modules VALUES ($1, $2, $3, $4, $5)", module.Id, module.Title, module.ModuleOrder, module.TopicCount, module.Overview)
	if err != nil {
		logrus.Error("Error inserting module data: ", err)
		return err
	}
	return nil
}

func AddTopics(db *sql.DB, topic Topic) error {
	_, err := db.Exec("INSERT INTO topics VALUES ($1, $2, $3, $4, $5q)", topic.Id, topic.Title, topic.ModuleId, topic.Order, topic.Minutes)
	if err != nil {
		logrus.Error("Error inserting module data: ", err)
		return err
	}
	return nil
}

func GetModules(db *sql.DB) ([]Module, error) {
	modules := make([]Module, 9)
	rows, err := db.Query("SELECT * FROM modules;")
	if err != nil {
		logrus.Error("Error retrieving rows: ", err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var module Module
		if err := rows.Scan(&module.Id, &module.Title, &module.ModuleOrder, &module.TopicCount, &module.Overview); err != nil {
			logrus.Error("Error scanning rows: ", err)
			return nil, err
		}
		modules = append(modules, module)
	}
	if !rows.NextResultSet() {
		logrus.Error("Expected more rows: ", rows.Err())
		return nil, err
	}
	return modules, nil
}
