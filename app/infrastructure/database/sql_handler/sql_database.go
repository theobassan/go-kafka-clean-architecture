package sql_handler

import (
	"database/sql"
	"errors"
	"go-kafka-clean-architecture/app/interfaces/database"
)

func NewSQLDatabase(driver, dataSource string) (database.SQLHandler, error) {
	sqlDB, err := sql.Open(driver, dataSource)
	if !errors.Is(err, nil) {
		return nil, err
	}
	err = sqlDB.Ping()
	if !errors.Is(err, nil) {
		return nil, err
	}

	return sqlDB, nil
}
