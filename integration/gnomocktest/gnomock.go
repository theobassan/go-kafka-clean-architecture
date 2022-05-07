package gnomocktest

import (
	"fmt"
	"go-kafka-clean-architecture/app/infrastructure/api/event_api"
	"go-kafka-clean-architecture/app/infrastructure/database/sql_gorm"
	"go-kafka-clean-architecture/app/infrastructure/database/sql_handler"
	"go-kafka-clean-architecture/app/interfaces/api"
	"go-kafka-clean-architecture/app/interfaces/database"
	"os"

	"github.com/go-errors/errors"

	"github.com/orlangure/gnomock"
	"github.com/orlangure/gnomock/preset/kafka"
	"github.com/orlangure/gnomock/preset/mysql"
	"github.com/orlangure/gnomock/preset/postgres"
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
		return nil, errors.Wrap(err, 1)
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
		return nil, errors.Wrap(err, 1)
	}

	return kafkaC, nil
}

func SetupSqlHandlerMySql() (database.SqlHandler, *gnomock.Container, error) {
	mySqlC, err := SetupMySql()
	if !errors.Is(err, nil) {
		return nil, nil, errors.Wrap(err, 1)
	}

	addr := mySqlC.DefaultAddress()
	connectionString := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s",
		dbSqlUsername, dbSqlPassword, addr, dbSqlName,
	)

	mySqlDb, err := sql_handler.NewSqlDatabase("mysql", connectionString)
	if !errors.Is(err, nil) {
		gnomock.Stop(mySqlC)
		return nil, nil, errors.Wrap(err, 1)
	}

	return mySqlDb, mySqlC, nil
}

func SetupPostgres() (*gnomock.Container, error) {
	seedDataPath, err := os.Getwd()
	if !errors.Is(err, nil) {
		return nil, errors.Wrap(err, 1)
	}
	schema := seedDataPath + "/db-data/postgres-schema.sql"

	postgresC, err := gnomock.Start(
		postgres.Preset(
			postgres.WithUser(dbSqlUsername, dbSqlPassword),
			postgres.WithDatabase(dbSqlName),
			postgres.WithQueriesFile(schema),
		),
		//gnomock.WithDebugMode(),
		//gnomock.WithLogWriter(os.Stdout),
	)
	if !errors.Is(err, nil) {
		return nil, errors.Wrap(err, 1)
	}

	return postgresC, nil
}

func SetupSqlHandlerPostgres() (database.SqlHandler, *gnomock.Container, error) {
	postgresC, err := SetupPostgres()
	if !errors.Is(err, nil) {
		return nil, nil, errors.Wrap(err, 1)
	}
	connectionString := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s  dbname=%s sslmode=disable",
		postgresC.Host, postgresC.DefaultPort(), dbSqlUsername, dbSqlPassword, dbSqlName,
	)

	postgresDb, err := sql_handler.NewSqlDatabase("postgres", connectionString)
	if !errors.Is(err, nil) {
		gnomock.Stop(postgresC)
		return nil, nil, errors.Wrap(err, 1)
	}

	return postgresDb, postgresC, nil
}

func SetupSqlGormPostgres() (database.SqlGorm, *gnomock.Container, error) {
	postgresC, err := SetupPostgres()
	if !errors.Is(err, nil) {
		return nil, nil, errors.Wrap(err, 1)
	}
	connectionString := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s  dbname=%s sslmode=disable",
		postgresC.Host, postgresC.DefaultPort(), dbSqlUsername, dbSqlPassword, dbSqlName,
	)

	postgresDb, err := sql_gorm.NewGormDatabase("postgres", connectionString)
	if !errors.Is(err, nil) {
		gnomock.Stop(postgresC)
		return nil, nil, errors.Wrap(err, 1)
	}

	return postgresDb, postgresC, nil
}

func SetupKafka() (*gnomock.Container, error) {
	kafkaC, err := gnomock.Start(
		kafka.Preset(kafka.WithTopics("product")),
		//gnomock.WithDebugMode(),
		//gnomock.WithLogWriter(os.Stdout),
	)
	if !errors.Is(err, nil) {
		return nil, errors.Wrap(err, 1)
	}

	return kafkaC, nil
}

func SetupEventApi() (api.EventApi, *gnomock.Container, error) {
	kafkaC, err := SetupKafka()
	if !errors.Is(err, nil) {
		return nil, nil, errors.Wrap(err, 1)
	}

	connectionString := kafkaC.Address(kafka.BrokerPort)
	eventApi := event_api.NewKafkaApi(connectionString)

	return eventApi, kafkaC, nil
}
