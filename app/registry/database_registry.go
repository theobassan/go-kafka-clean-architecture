package registry

import (
	"errors"
	"fmt"
	"go-kafka-clean-architecture/app/infrastructure/database/sql_gorm"
	"go-kafka-clean-architecture/app/infrastructure/database/sql_handler"
	"go-kafka-clean-architecture/app/interfaces/database"

	"github.com/go-sql-driver/mysql"
)

func (r *Registry) NewSqlGormMySql() database.SqlGorm {
	mySqlConfig := &mysql.Config{
		User:   "cleanarchitecture",
		Passwd: "cleanarchitecture",
		Net:    "tcp",
		Addr:   "localhost:3306",
		DBName: "cleanarchitecture",
	}

	mySqlDb, err := sql_gorm.NewGormDatabase("mysql", mySqlConfig.FormatDSN())
	if !errors.Is(err, nil) {
		panic(err)
	}
	return mySqlDb
}

func (r *Registry) NewSqlGormPostgres() database.SqlGorm {
	dataSourceName := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", "localhost", "cleanarchitecture", "cleanarchitecture", "cleanarchitecture", "5432")
	postgresDb, err := sql_gorm.NewGormDatabase("postgres", dataSourceName)
	if !errors.Is(err, nil) {
		panic(err)
	}
	return postgresDb
}

func (r *Registry) NewSqlHandlerMySql() database.SqlHandler {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", "cleanarchitecture", "cleanarchitecture", "localhost", "3306", "cleanarchitecture")
	mySqlDb, err := sql_handler.NewSqlDatabase("mysql", dataSourceName)
	if !errors.Is(err, nil) {
		panic(err)
	}

	return mySqlDb
}

func (r *Registry) NewSqlHandlerPostgres() database.SqlHandler {
	dataSourceName := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", "localhost", "cleanarchitecture", "cleanarchitecture", "cleanarchitecture", "5432")
	postgresDb, err := sql_handler.NewSqlDatabase("postgres", dataSourceName)
	if !errors.Is(err, nil) {
		panic(err)
	}
	return postgresDb
}
