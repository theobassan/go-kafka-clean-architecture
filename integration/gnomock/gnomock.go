package gnomock

import (
	"errors"
	"fmt"
	"go-kafka-clean-architecture/app/infrastructure/api/event_api"
	"go-kafka-clean-architecture/app/infrastructure/database/sql_handler"
	"go-kafka-clean-architecture/app/interfaces/api"
	"go-kafka-clean-architecture/app/interfaces/database"
	"os"

	"github.com/orlangure/gnomock"
	"github.com/orlangure/gnomock/preset/kafka"
	"github.com/orlangure/gnomock/preset/mysql"
)

const (
	dbSqlName     string = "cleanarchitecture"
	dbSqlUsername string = "cleanarchitecture"
	dbSqlPassword string = "cleanarchitecture"
)

func Stop(container *gnomock.Container) error {
	return gnomock.Stop(container)
}

func SetupMySql() (*gnomock.Container, error) {
	seedDataPath, err := os.Getwd()
	if !errors.Is(err, nil) {
		return nil, err
	}
	schema := seedDataPath + "/db-data/mysql-schema.sql"

	kafkaC, err := gnomock.Start(
		mysql.Preset(
			mysql.WithUser(dbSqlUsername, dbSqlPassword),
			mysql.WithDatabase(dbSqlName),
			mysql.WithQueriesFile(schema),
		),
		//gnomock.WithDebugMode(),
		//gnomock.WithLogWriter(os.Stdout),
	)
	if !errors.Is(err, nil) {
		return nil, err
	}

	return kafkaC, nil
}

func SetupSQLHandlerMySQL() (database.SQLHandler, *gnomock.Container, error) {
	mySqlC, err := SetupMySql()
	if !errors.Is(err, nil) {
		return nil, nil, err
	}
	//

	addr := mySqlC.DefaultAddress()
	connectionString := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s",
		dbSqlUsername, dbSqlPassword, addr, dbSqlName,
	)

	mySqlDb, err := sql_handler.NewSQLDatabase("mysql", connectionString)
	if !errors.Is(err, nil) {
		return nil, nil, err
	}

	return mySqlDb, mySqlC, nil
}

func SetupKafka() (*gnomock.Container, error) {
	kafkaC, err := gnomock.Start(
		kafka.Preset(kafka.WithTopics("product")),
		//gnomock.WithDebugMode(),
		//gnomock.WithLogWriter(os.Stdout),
	)
	if !errors.Is(err, nil) {
		return nil, err
	}

	return kafkaC, nil
}

func SetupEventAPI() (api.EventAPI, *gnomock.Container, error) {
	kafkaC, err := SetupKafka()
	if !errors.Is(err, nil) {
		return nil, nil, err
	}

	connectionString := kafkaC.Address(kafka.BrokerPort)
	eventAPI := event_api.NewKafkaAPI(connectionString)

	return eventAPI, kafkaC, nil
}
