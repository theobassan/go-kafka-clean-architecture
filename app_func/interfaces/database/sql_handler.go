package database

import "database/sql"

type SqlHandlerQuery func(query string, args ...any) (*sql.Rows, error)
type SqlHandlerExec func(query string, args ...any) (sql.Result, error)
type SqlHandlerQueryRow func(query string, args ...any) *sql.Row
