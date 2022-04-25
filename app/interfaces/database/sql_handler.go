package database

import "database/sql"

// A SQLHandler belong to the inteface layer.
type SQLHandler interface {
	Query(query string, args ...any) (*sql.Rows, error)
	Exec(query string, args ...any) (sql.Result, error)
}
