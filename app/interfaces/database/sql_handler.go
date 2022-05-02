package database

import "database/sql"

// A SQLHandler belong to the inteface layer.
type SQLHandler interface {
	Query(query string, args ...any) (*sql.Rows, error)
	QueryRow(query string, args ...any) *sql.Row
	Exec(query string, args ...any) (sql.Result, error)
}
