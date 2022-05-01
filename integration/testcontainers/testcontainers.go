package testcontainers

import (
	"context"
	"errors"
	"fmt"
	"go-kafka-clean-architecture/app/infrastructure/api/event_api"
	"go-kafka-clean-architecture/app/infrastructure/database/sql_handler"
	"go-kafka-clean-architecture/app/interfaces/api"
	"go-kafka-clean-architecture/app/interfaces/database"
	"os"
	"strconv"
	"strings"

	"github.com/docker/go-connections/nat"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

const (
	dbSqlName     string = "cleanarchitecture"
	dbSqlUsername string = "cleanarchitecture"
	dbSqlPassword string = "cleanarchitecture"
)

func SetupMySql(ctx context.Context) (testcontainers.Container, error) {

	seedDataPath, err := os.Getwd()
	if !errors.Is(err, nil) {
		return nil, err
	}
	schema := seedDataPath + "/db-data/mysql-schema.sql"
	config := seedDataPath + "/db-data/mysql.cnf"

	req := testcontainers.ContainerRequest{
		Image:        "mysql:latest",
		ExposedPorts: []string{"3306/tcp"},
		Env: map[string]string{
			"MYSQL_DATABASE":      dbSqlName,
			"MYSQL_USER":          dbSqlUsername,
			"MYSQL_PASSWORD":      dbSqlPassword,
			"MYSQL_ROOT_PASSWORD": dbSqlPassword,
		},
		Mounts: testcontainers.ContainerMounts{
			{Source: testcontainers.GenericBindMountSource{HostPath: schema}, Target: "/docker-entrypoint-initdb.d/mysql-schema.sql"},
			{Source: testcontainers.GenericBindMountSource{HostPath: config}, Target: "/etc/mysql/conf.d/my.cnf"},
		},
		WaitingFor: wait.ForLog("port: 3306  MySQL Community Server - GPL"),
		SkipReaper: true,
	}

	mySqlC, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if !errors.Is(err, nil) {
		return nil, err
	}

	return mySqlC, nil
}

func SetupSQLHandlerMySql(ctx context.Context) (database.SQLHandler, testcontainers.Container, error) {

	mySqlC, err := SetupMySql(ctx)
	if !errors.Is(err, nil) {
		return nil, nil, err
	}

	host, _ := mySqlC.Host(ctx)
	p, _ := mySqlC.MappedPort(ctx, nat.Port("3306/tcp"))
	port := strconv.Itoa(p.Int())
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbSqlUsername, dbSqlPassword, host, port, dbSqlName)

	mySqlDb, err := sql_handler.NewSQLDatabase("mysql", connectionString)
	if !errors.Is(err, nil) {
		return nil, nil, err
	}

	return mySqlDb, mySqlC, nil
}

func SetupPostgres(ctx context.Context) (testcontainers.Container, error) {

	seedDataPath, err := os.Getwd()
	if !errors.Is(err, nil) {
		return nil, err
	}
	schema := seedDataPath + "/db-data/postgres-schema.sql"

	req := testcontainers.ContainerRequest{
		Image:        "postgres:latest",
		ExposedPorts: []string{"5432/tcp"},
		Cmd:          []string{"postgres", "-c", "fsync=off"},
		Env: map[string]string{
			"POSTGRES_DB":       dbSqlName,
			"POSTGRES_USER":     dbSqlUsername,
			"POSTGRES_PASSWORD": dbSqlPassword,
		},
		Mounts: testcontainers.ContainerMounts{
			{Source: testcontainers.GenericBindMountSource{HostPath: schema}, Target: "/docker-entrypoint-initdb.d/postgres-schema.sql"},
		},
		WaitingFor: wait.ForLog("database system is ready to accept connections"),
		SkipReaper: true,
	}

	postgresC, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if !errors.Is(err, nil) {
		return nil, err
	}

	return postgresC, nil
}

func SetupSQLHandlerPostgres(ctx context.Context) (database.SQLHandler, testcontainers.Container, error) {

	postgresC, err := SetupPostgres(ctx)
	if !errors.Is(err, nil) {
		return nil, nil, err
	}

	host, _ := postgresC.Host(ctx)
	p, _ := postgresC.MappedPort(ctx, nat.Port("5432/tcp"))
	port := strconv.Itoa(p.Int())

	//connectionString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbSqlUsername, dbSqlPassword, host, port, dbSqlName)
	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, dbSqlUsername, dbSqlPassword, dbSqlName)

	postgresDb, err := sql_handler.NewSQLDatabase("postgres", connectionString)
	if !errors.Is(err, nil) {
		return nil, nil, err
	}

	return postgresDb, postgresC, nil
}

func SetupZookeeper(ctx context.Context, networkName string) (testcontainers.Container, error) {

	aliases := make(map[string][]string)
	aliases[networkName] = []string{"zookeeper"}

	req := testcontainers.ContainerRequest{
		Image:        "confluentinc/cp-zookeeper:latest",
		ExposedPorts: []string{"2181/tcp"},
		Env: map[string]string{
			"ZOOKEEPER_CLIENT_PORT": "2181",
			"ZOOKEEPER_TICK_TIME":   "2000",
		},
		Networks:       []string{networkName},
		NetworkAliases: aliases,
		Hostname:       "zookeeper",
		Name:           "zookeeper",
		WaitingFor:     wait.ForLog("binding to port 0.0.0.0/0.0.0.0:2181"),
		SkipReaper:     true,
	}

	zookeeperC, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if !errors.Is(err, nil) {
		return nil, err
	}

	return zookeeperC, nil
}

func SetupKafka(ctx context.Context, zookeeperHost string, zookeeperPort string, networkName string) (testcontainers.Container, error) {
	req := testcontainers.ContainerRequest{
		Image:        "confluentinc/cp-server:latest",
		ExposedPorts: []string{"9092/tcp"},
		Env: map[string]string{
			"KAFKA_BROKER_ID":                                   "1",
			"KAFKA_ZOOKEEPER_CONNECT":                           fmt.Sprintf("%s:%s", "zookeeper", "2181"),
			"KAFKA_LISTENER_SECURITY_PROTOCOL_MAP":              "PLAINTEXT:PLAINTEXT,PLAINTEXT_HOST:PLAINTEXT",
			"KAFKA_ADVERTISED_LISTENERS":                        "PLAINTEXT://broker:29092,PLAINTEXT_HOST://localhost:9092",
			"KAFKA_METRIC_REPORTERS":                            "io.confluent.metrics.reporter.ConfluentMetricsReporter",
			"KAFKA_OFFSETS_TOPIC_REPLICATION_FACTOR":            "1",
			"KAFKA_GROUP_INITIAL_REBALANCE_DELAY_MS":            "0",
			"KAFKA_CONFLUENT_LICENSE_TOPIC_REPLICATION_FACTOR":  "1",
			"KAFKA_CONFLUENT_BALANCER_TOPIC_REPLICATION_FACTOR": "1",
			"KAFKA_TRANSACTION_STATE_LOG_MIN_ISR":               "1",
			"KAFKA_TRANSACTION_STATE_LOG_REPLICATION_FACTOR":    "1",
			"KAFKA_JMX_PORT":                                    "9101",
			"KAFKA_JMX_HOSTNAME":                                "localhost",
			"KAFKA_CONFLUENT_SCHEMA_REGISTRY_URL":               "http://schema-registry:8081",
			"CONFLUENT_METRICS_REPORTER_BOOTSTRAP_SERVERS":      "broker:29092",
			"CONFLUENT_METRICS_REPORTER_TOPIC_REPLICAS":         "1",
			"CONFLUENT_METRICS_ENABLE":                          "true",
			"CONFLUENT_SUPPORT_CUSTOMER_ID":                     "anonymous",
			"KAFKA_CREATE_TOPICS":                               "product:1:1",
		},
		Networks:   []string{networkName},
		Hostname:   "broker",
		Name:       "broker",
		WaitingFor: wait.ForListeningPort(nat.Port("9092/tcp")),
		SkipReaper: true,
	}

	kafkaC, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if !errors.Is(err, nil) {
		return nil, err
	}

	return kafkaC, nil
}

func SetupEventAPI(ctx context.Context) (api.EventAPI, testcontainers.Container, error) {
	networkName := "kafka-network"
	network, err := testcontainers.GenericNetwork(ctx, testcontainers.GenericNetworkRequest{
		NetworkRequest: testcontainers.NetworkRequest{
			Driver:         "bridge",
			Name:           networkName,
			Attachable:     true,
			CheckDuplicate: true,
			SkipReaper:     true,
		},
	})
	if !errors.Is(err, nil) {
		return nil, nil, err
	}
	defer network.Remove(ctx)

	zookeeperC, err := SetupZookeeper(ctx, networkName)
	if !errors.Is(err, nil) {
		return nil, nil, err
	}
	defer zookeeperC.Terminate(ctx)

	zookeeperHost, _ := zookeeperC.Host(ctx)
	zookeeperP, _ := zookeeperC.MappedPort(ctx, nat.Port("2181/tcp"))
	zookeeperPort := strconv.Itoa(zookeeperP.Int())

	kafkaC, err := SetupKafka(ctx, zookeeperHost, zookeeperPort, networkName)
	if !errors.Is(err, nil) {
		return nil, nil, err
	}
	defer kafkaC.Terminate(ctx)

	kafkaHost, _ := kafkaC.Host(ctx)
	kafkaP, _ := zookeeperC.MappedPort(ctx, nat.Port("9092/tcp"))
	kafkaPort := strconv.Itoa(kafkaP.Int())
	connectionString := fmt.Sprintf("%s:%s", kafkaHost, kafkaPort)

	eventAPI := event_api.NewKafkaAPI(connectionString)

	return eventAPI, kafkaC, nil
}

func SetupKafkaCompose(ctx context.Context) (*testcontainers.LocalDockerCompose, error) {
	seedDataPath, err := os.Getwd()
	if !errors.Is(err, nil) {
		return nil, err
	}

	/*port, _ := nat.NewPort("", strconv.Itoa(9092))
	if !errors.Is(err, nil) {
		return nil, err
	}*/

	kafkaC := testcontainers.NewLocalDockerCompose(
		[]string{seedDataPath + "/testcontainers/docker-compose.yml"},
		strings.ToLower(uuid.New().String()),
	)
	//kafkaC.WaitForService("broker", wait.ForListeningPort(port))

	err = kafkaC.WithCommand([]string{"up", "-d", "broker"}).Invoke().Error
	if !errors.Is(err, nil) {
		return nil, err
	}

	return kafkaC, nil
}

func SetupEventAPICompose(ctx context.Context) (api.EventAPI, *testcontainers.LocalDockerCompose, error) {

	kafkaC, err := SetupKafkaCompose(ctx)
	if !errors.Is(err, nil) {
		return nil, nil, err
	}
	//defer kafkaC.Down(ctx)

	eventAPI := event_api.NewKafkaAPI("localhost:9092")

	return eventAPI, kafkaC, nil
}
