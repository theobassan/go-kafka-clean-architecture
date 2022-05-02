package main

import (
	"fmt"
	event_context_interfaces "go-kafka-clean-architecture/app/command/controller/event_context"
	http_context_interfaces "go-kafka-clean-architecture/app/command/controller/http_context"
	"go-kafka-clean-architecture/app/infrastructure/api/event_api"
	"go-kafka-clean-architecture/app/infrastructure/api/rest_api"
	"go-kafka-clean-architecture/app/infrastructure/database/sql_gorm"
	"go-kafka-clean-architecture/app/infrastructure/database/sql_handler"
	event_context_infrastructure "go-kafka-clean-architecture/app/infrastructure/router/event_context"
	http_context_infrastructure "go-kafka-clean-architecture/app/infrastructure/router/http_context"
	"go-kafka-clean-architecture/registry"

	"github.com/go-errors/errors"

	"github.com/go-sql-driver/mysql"
	_ "gorm.io/driver/postgres"
	//"github.com/jinzhu/gorm/dialects/mysql"
)

func main() {

	restAppController, eventAppController := registerSQLHandlerSQL()
	//sqlKafkaAppController := registerSQLHandlerGorm()
	//sqlKafkaAppController := registerGormHandler()

	go event_context_infrastructure.StartKafkaRouter(eventAppController, "localhost:9092")
	http_context_infrastructure.StartEchoRouter(restAppController, 8080)
}

func registerSQLHandlerSQL() (*http_context_interfaces.AppController, *event_context_interfaces.AppController) {

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", "cleanarchitecture", "cleanarchitecture", "localhost", "3306", "cleanarchitecture")
	mySqlDb, err := sql_handler.NewSQLDatabase("mysql", dataSourceName)
	if !errors.Is(err, nil) {
		panic(err)
	}

	restAPI := rest_api.NewHttpAPI("http://localhost:8080")
	eventAPI := event_api.NewKafkaAPI("localhost:9092")

	r := registry.NewRegistry()
	return r.NewHttpContextRestSqlEventAppController(restAPI, mySqlDb, eventAPI), r.NewEventContextRestSqlEventAppController(restAPI, mySqlDb, eventAPI)
}

func registerSQLHandlerGorm() (*http_context_interfaces.AppController, *event_context_interfaces.AppController) {
	mySqlConfig := &mysql.Config{
		User:   "cleanarchitecture",
		Passwd: "cleanarchitecture",
		Net:    "tcp",
		Addr:   "localhost:3306",
		DBName: "cleanarchitecture",
	}

	mySqlDb, err := sql_handler.NewSQLGormDatabase("mysql", mySqlConfig.FormatDSN())
	if !errors.Is(err, nil) {
		panic(err)
	}

	restAPI := rest_api.NewHttpAPI("http://localhost:8080")
	eventAPI := event_api.NewKafkaAPI("localhost:9092")

	r := registry.NewRegistry()
	return r.NewHttpContextRestSqlEventAppController(restAPI, mySqlDb, eventAPI), r.NewEventContextRestSqlEventAppController(restAPI, mySqlDb, eventAPI)
}

func registerGormHandler() (*http_context_interfaces.AppController, *event_context_interfaces.AppController) {
	mySqlConfig := &mysql.Config{
		User:   "cleanarchitecture",
		Passwd: "cleanarchitecture",
		Net:    "tcp",
		Addr:   "localhost:3306",
		DBName: "cleanarchitecture",
	}

	mySqlDb, err := sql_gorm.NewGormDatabase("mysql", mySqlConfig.FormatDSN())
	//dataSourceName := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", "localhost", "cleanarchitecture", "cleanarchitecture", "cleanarchitecture", "5432")
	//postgreSqlDb, err := sql_gorm.NewGormDatabase("pgx", dataSourceName)
	if !errors.Is(err, nil) {
		panic(err)
	}

	restAPI := rest_api.NewHttpAPI("http://localhost:8080")
	eventAPI := event_api.NewKafkaAPI("localhost:9092")

	r := registry.NewRegistry()
	return r.NewHttpContextRestGormEventAppController(restAPI, mySqlDb, eventAPI), r.NewEventContextRestGormEventAppController(restAPI, mySqlDb, eventAPI)
}
