package database

import (
	"github.com/go-errors/errors"

	"github.com/DATA-DOG/go-sqlmock"
)

type NewProductRepositoryMock struct {
}

func NewSQLHandlerMock() (SQLHandler, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()
	if !errors.Is(err, nil) {
		return nil, nil, errors.Wrap(err, 1)
	}

	return db, mock, nil
}
