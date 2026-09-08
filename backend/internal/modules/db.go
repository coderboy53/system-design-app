package modules

import (
	"database/sql"

	"github.com/sirupsen/logrus"
)

// tmp functions for populating data
func AddModules(db *sql.DB, body Module) {
	

}

func AddTopics(db *sql.DB, body Topic) {

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
