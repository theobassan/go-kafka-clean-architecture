package sql_gorm

import (
	"go-kafka-clean-architecture/app/interfaces/database"

	"github.com/go-errors/errors"

	"github.com/jinzhu/gorm"
)

func NewGormDatabase(dialect string, source string) (database.SqlGorm, error) {

	gormDB, err := gorm.Open(dialect, source)
	if !errors.Is(err, nil) {
		return nil, errors.Wrap(err, 1)
	}
	err = gormDB.DB().Ping()
	if !errors.Is(err, nil) {
		return nil, errors.Wrap(err, 1)
	}

	return gormDB, nil
}
