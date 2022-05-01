package sql_handler

import (
	"errors"
	"go-kafka-clean-architecture/app/interfaces/database"

	"github.com/jinzhu/gorm"
)

func NewSQLGormDatabase(dialect string, source string) (database.SQLHandler, error) {

	gormDB, err := gorm.Open(dialect, source)
	if !errors.Is(err, nil) {
		return nil, err
	}
	err = gormDB.DB().Ping()
	if !errors.Is(err, nil) {
		return nil, err
	}

	return gormDB.DB(), nil
}
