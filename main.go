package main

import (
	"go-kafka-clean-architecture/golang/pointers"

	_ "github.com/go-sql-driver/mysql"
	_ "gorm.io/driver/postgres"
)

func main() {
	pointers.Run()
	//recursion.Run()
	//concurrency.Run()
	//goroutine.Run()
	//goroutinemanager.Run()
	//app.StartEventContextHttpContextRestApiSqlGormMySqlEventApiBrasil()
}
