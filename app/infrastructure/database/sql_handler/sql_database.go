package sql_handler

import (
	"database/sql"
	"go-kafka-clean-architecture/app/interfaces/database"
)

func NewSQLDatabase(driver, dataSource string) (database.SQLHandler, error) {
	sqlDB, err := sql.Open(driver, dataSource)
	if err != nil {
		return nil, err
	}
	err = sqlDB.Ping()
	if err != nil {
		return nil, err
	}

	return sqlDB, nil
}
