package integration

import (
	"context"
	"go-kafka-clean-architecture/app/infrastructure/api/rest_api"
	"go-kafka-clean-architecture/app/infrastructure/logger"
	event_context_infrastructure "go-kafka-clean-architecture/app/infrastructure/router/event_context"
	http_context_infrastructure "go-kafka-clean-architecture/app/infrastructure/router/http_context"
	"go-kafka-clean-architecture/integration/test"
	"go-kafka-clean-architecture/integration/testcontainers"
	"go-kafka-clean-architecture/registry"
	"strconv"
	"testing"

	"github.com/docker/go-connections/nat"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"
)

func TestCreate_testcontainers_mysql(t *testing.T) {
	port, _ := nat.NewPort("", strconv.Itoa(8080))
	ctx := context.Background()

	mySqlDb, mySqlC, err := testcontainers.SetupSQLHandlerMySql(ctx)
	require.NoError(t, err)
	//defer mySqlC.Terminate(ctx)

	kafkaAPI, zookeeperC, kafkaC, err := testcontainers.SetupEventAPI(ctx)
	require.NoError(t, err)
	//defer zookeeperC.Terminate(ctx)
	//defer kafkaC.Terminate(ctx)

	serverURL := "http://localhost:" + strconv.Itoa(port.Int())

	restAPI := rest_api.NewHttpAPI(serverURL)

	r := registry.NewRegistry()
	httpContextAppController := r.NewHttpContextRestSqlEventAppControllerMySql(restAPI, mySqlDb, kafkaAPI)
	eventContextAppController := r.NewEventContextRestSqlEventAppControllerMySql(restAPI, mySqlDb, kafkaAPI)
	logger := logger.NewDebugLogger()

	go event_context_infrastructure.StartKafkaRouter(eventContextAppController, "localhost:9092", logger)
	go http_context_infrastructure.StartEchoRouter(httpContextAppController, port.Int(), logger)

	test.TestCreate(t, serverURL)

	err = mySqlC.Terminate(ctx)
	require.NoError(t, err)

	err = kafkaC.Terminate(ctx)
	require.NoError(t, err)

	err = zookeeperC.Terminate(ctx)
	require.NoError(t, err)
}

func TestFindAll_testcontainers_mysql(t *testing.T) {
	port, _ := nat.NewPort("", strconv.Itoa(8080))
	ctx := context.Background()

	mySqlDb, mySqlC, err := testcontainers.SetupSQLHandlerMySql(ctx)
	require.NoError(t, err)
	//defer mySqlC.Terminate(ctx)

	serverURL := "http://localhost:" + strconv.Itoa(port.Int())

	r := registry.NewRegistry()
	httpContextAppController := r.NewHttpContextRestSqlEventAppControllerMySql(nil, mySqlDb, nil)
	logger := logger.NewDebugLogger()

	go http_context_infrastructure.StartEchoRouter(httpContextAppController, port.Int(), logger)

	productID := int64(123)
	productType := "Type"
	productName := "Name"

	_, err = mySqlDb.Exec(`
		INSERT INTO
			products(external_id, type, name)
		VALUES
			(?, ?, ?)
	`, productID, productType, productName)
	require.NoError(t, err)

	_, err = mySqlDb.Exec(`
		INSERT INTO
			products_translated(external_id, type, name)
		VALUES
		(?, ?, ?)
	`, productID, productType, productName)
	require.NoError(t, err)

	test.TestFindAll(t, serverURL, &productID, &productType, &productName, &productID, &productType, &productName)

	err = mySqlC.Terminate(ctx)
	require.NoError(t, err)
}

func TestFindAll_testcontainers_postgres(t *testing.T) {
	port, _ := nat.NewPort("", strconv.Itoa(8080))
	ctx := context.Background()

	postgresDb, postgresC, err := testcontainers.SetupSQLHandlerPostgres(ctx)
	require.NoError(t, err)
	//defer postgresC.Terminate(ctx)

	serverURL := "http://localhost:" + strconv.Itoa(port.Int())

	r := registry.NewRegistry()
	httpContextAppController := r.NewHttpContextRestSqlEventAppControllerPostgres(nil, postgresDb, nil)
	logger := logger.NewDebugLogger()

	go http_context_infrastructure.StartEchoRouter(httpContextAppController, port.Int(), logger)

	productID := int64(123)
	productType := "Type"
	productName := "Name"

	_, err = postgresDb.Exec(`
		INSERT INTO
			"products" ("external_id", "type", "name")
		VALUES
			($1, $2, $3)
	`, productID, productType, productName)
	require.NoError(t, err)

	_, err = postgresDb.Exec(`
		INSERT INTO
			"products_translated" ("external_id", "type", "name")
		VALUES
			($1, $2, $3)
	`, productID, productType, productName)
	require.NoError(t, err)

	test.TestFindAll(t, serverURL, &productID, &productType, &productName, &productID, &productType, &productName)

	err = postgresC.Terminate(ctx)
	require.NoError(t, err)
}
