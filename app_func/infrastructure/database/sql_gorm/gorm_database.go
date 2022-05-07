package sql_gorm

import (
	"go-kafka-clean-architecture/app_func/interfaces/database"

	"github.com/go-errors/errors"

	"github.com/jinzhu/gorm"
)

func NewGormDatabase(dialect string, source string) (database.SqlGormFind, database.SqlGormCreate, error) {

	gormDb, err := gorm.Open(dialect, source)
	if !errors.Is(err, nil) {
		return nil, nil, errors.Wrap(err, 1)
	}
	err = gormDb.DB().Ping()
	if !errors.Is(err, nil) {
		return nil, nil, errors.Wrap(err, 1)
	}

	return gormDb.Find, gormDb.Create, nil
}
