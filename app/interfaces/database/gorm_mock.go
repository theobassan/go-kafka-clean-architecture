package database

import (
	"errors"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
)

func NewSqlGormMock(dialect string) (SQLGorm, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()
	if !errors.Is(err, nil) {
		return nil, nil, err
	}

	gormDB, err := gorm.Open(dialect, db)
	if !errors.Is(err, nil) {
		return nil, nil, err
	}

	return gormDB, mock, nil
}
