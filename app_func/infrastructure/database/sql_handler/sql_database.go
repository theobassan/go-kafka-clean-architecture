package sql_handler

import (
	"database/sql"
	"go-kafka-clean-architecture/app_func/interfaces/database"

	"github.com/go-errors/errors"
)

func NewSqlDatabase(driver, dataSource string) (database.SqlHandlerExec, database.SqlHandlerQuery, database.SqlHandlerQueryRow, error) {
	sqlDb, err := sql.Open(driver, dataSource)
	if !errors.Is(err, nil) {
		return nil, nil, nil, errors.Wrap(err, 1)
	}
	err = sqlDb.Ping()
	if !errors.Is(err, nil) {
		return nil, nil, nil, errors.Wrap(err, 1)
	}

	return sqlDb.Exec, sqlDb.Query, sqlDb.QueryRow, nil
}
