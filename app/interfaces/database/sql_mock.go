package database

import (
	"errors"

	"github.com/DATA-DOG/go-sqlmock"
)

type NewProductRepositoryMock struct {
}

func NewSQLHandlerMock() (SQLHandler, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()
	if !errors.Is(err, nil) {
		return nil, nil, err
	}

	return db, mock, nil
}
