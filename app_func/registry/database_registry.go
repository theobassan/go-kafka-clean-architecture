package registry

import (
	"errors"
	"fmt"
	"go-kafka-clean-architecture/app_func/infrastructure/database/sql_gorm"
	"go-kafka-clean-architecture/app_func/infrastructure/database/sql_handler"
	"go-kafka-clean-architecture/app_func/interfaces/database"

	"github.com/go-sql-driver/mysql"
)

func SqlGormMySql() (database.SqlGormCreate, database.SqlGormFind) {
	mySqlConfig := &mysql.Config{
		User:   "cleanarchitecture",
		Passwd: "cleanarchitecture",
		Net:    "tcp",
		Addr:   "localhost:3306",
		DBName: "cleanarchitecture",
	}

	find, create, err := sql_gorm.NewGormDatabase("mysql", mySqlConfig.FormatDSN())
	if !errors.Is(err, nil) {
		panic(err)
	}
	return create, find
}

func SqlGormPostgres() (database.SqlGormCreate, database.SqlGormFind) {
	dataSourceName := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", "localhost", "cleanarchitecture", "cleanarchitecture", "cleanarchitecture", "5432")
	find, create, err := sql_gorm.NewGormDatabase("postgres", dataSourceName)
	if !errors.Is(err, nil) {
		panic(err)
	}
	return create, find
}

func SqlHandlerMySql() (database.SqlHandlerExec, database.SqlHandlerQuery, database.SqlHandlerQueryRow) {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", "cleanarchitecture", "cleanarchitecture", "localhost", "3306", "cleanarchitecture")
	exec, query, queryRow, err := sql_handler.NewSqlDatabase("mysql", dataSourceName)
	if !errors.Is(err, nil) {
		panic(err)
	}

	return exec, query, queryRow
}

func SqlHandlerPostgres() (database.SqlHandlerExec, database.SqlHandlerQuery, database.SqlHandlerQueryRow) {
	dataSourceName := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", "localhost", "cleanarchitecture", "cleanarchitecture", "cleanarchitecture", "5432")
	exec, query, queryRow, err := sql_handler.NewSqlDatabase("postgres", dataSourceName)
	if !errors.Is(err, nil) {
		panic(err)
	}
	return exec, query, queryRow
}
