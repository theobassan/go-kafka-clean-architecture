package sql_handler

import (
	"database/sql"
	"go-kafka-clean-architecture/app/interfaces/database"

	"github.com/go-errors/errors"
)

func NewSQLDatabase(driver, dataSource string) (database.SQLHandler, error) {
	sqlDB, err := sql.Open(driver, dataSource)
	if !errors.Is(err, nil) {
		return nil, errors.Wrap(err, 1)
	}
	err = sqlDB.Ping()
	if !errors.Is(err, nil) {
		return nil, errors.Wrap(err, 1)
	}

	return sqlDB, nil
}
