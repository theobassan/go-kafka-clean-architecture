package database

import (
	"github.com/go-errors/errors"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
)

func NewSqlGormMock(dialect string) (*gorm.DB, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()
	if !errors.Is(err, nil) {
		return nil, nil, errors.Wrap(err, 1)
	}

	gormDB, err := gorm.Open(dialect, db)
	if !errors.Is(err, nil) {
		return nil, nil, errors.Wrap(err, 1)
	}

	return gormDB, mock, nil
}
