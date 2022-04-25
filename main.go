package main

import (
	"fmt"
	"go-kafka-clean-architecture/app/infrastructure/api/rest_api"
	"go-kafka-clean-architecture/app/infrastructure/broker/kafka"
	"go-kafka-clean-architecture/app/infrastructure/database/sql_gorm"
	"go-kafka-clean-architecture/app/infrastructure/database/sql_handler"
	infrastructure "go-kafka-clean-architecture/app/infrastructure/router/rest_context"
	interfaces "go-kafka-clean-architecture/app/interfaces/controller/rest_context"
	"go-kafka-clean-architecture/registry"

	"github.com/go-sql-driver/mysql"
	_ "gorm.io/driver/postgres"
	//_ "github.com/go-sql-driver/mysql"
	//"github.com/jinzhu/gorm/dialects/mysql"
)

func main() {

	sqlKafkaAppController := registerSQLHandlerSQL()
	//sqlKafkaAppController := registerSQLHandlerGorm()
	//sqlKafkaAppController := registerGormHandler()

	//infrastructure.NewEchoRouter(sqlKafkaAppController, 8080)
	infrastructure.NewGinRouter(sqlKafkaAppController, 8080)
}

func registerSQLHandlerSQL() *interfaces.AppController {

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", "cleanarchitecture", "cleanarchitecture", "localhost", "3306", "cleanarchitecture")
	mySqlDb, err := sql_handler.NewSQLDatabase("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}

	kafkaProductWriter := kafka.NewKafkaWriter("localhost:9092", "product")
	kafkaProductReader := kafka.NewKafkaReader("localhost:9092", "product", "clean-architecture")

	restAPI := rest_api.NewRestAPI()

	r := registry.NewRegistry()
	return r.NewSqlBrokerAppController(restAPI, mySqlDb, kafkaProductWriter, kafkaProductReader)
}

func registerSQLHandlerGorm() *interfaces.AppController {
	mySqlConfig := &mysql.Config{
		User:   "cleanarchitecture",
		Passwd: "cleanarchitecture",
		Net:    "tcp",
		Addr:   "localhost:3306",
		DBName: "cleanarchitecture",
	}

	mySqlDb, err := sql_handler.NewSQLGormDatabase("mysql", mySqlConfig.FormatDSN())
	if err != nil {
		panic(err)
	}

	kafkaProductWriter := kafka.NewKafkaWriter("localhost:9092", "product")
	kafkaProductReader := kafka.NewKafkaReader("localhost:9092", "product", "clean-architecture")

	restAPI := rest_api.NewRestAPI()

	r := registry.NewRegistry()
	return r.NewSqlBrokerAppController(restAPI, mySqlDb, kafkaProductWriter, kafkaProductReader)
}

func registerGormHandler() *interfaces.AppController {
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
	if err != nil {
		panic(err)
	}

	kafkaProductWriter := kafka.NewKafkaWriter("localhost:9092", "product")
	kafkaProductReader := kafka.NewKafkaReader("localhost:9092", "product", "clean-architecture")

	restAPI := rest_api.NewRestAPI()

	r := registry.NewRegistry()
	return r.NewGormBrokerAppController(restAPI, mySqlDb, kafkaProductWriter, kafkaProductReader)
}
