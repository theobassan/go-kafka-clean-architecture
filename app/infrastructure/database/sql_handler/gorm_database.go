package sql_handler

import (
	"go-kafka-clean-architecture/app/interfaces/database"

	"github.com/go-errors/errors"

	"github.com/jinzhu/gorm"
)

func NewSQLGormDatabase(dialect string, source string) (database.SQLHandler, error) {

	gormDB, err := gorm.Open(dialect, source)
	if !errors.Is(err, nil) {
		return nil, errors.Wrap(err, 1)
	}
	err = gormDB.DB().Ping()
	if !errors.Is(err, nil) {
		return nil, errors.Wrap(err, 1)
	}

	return gormDB.DB(), nil
}
