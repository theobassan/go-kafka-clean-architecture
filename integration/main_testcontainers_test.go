package integration

import (
	"context"
	"go-kafka-clean-architecture/app/command/controller/event_context"
	"go-kafka-clean-architecture/app/command/controller/http_context"
	"go-kafka-clean-architecture/app/infrastructure/api/rest_api"
	"go-kafka-clean-architecture/app/infrastructure/logger"
	event_context_infrastructure "go-kafka-clean-architecture/app/infrastructure/router/event_context"
	http_context_infrastructure "go-kafka-clean-architecture/app/infrastructure/router/http_context"
	"go-kafka-clean-architecture/app/registry"
	"go-kafka-clean-architecture/integration/test"
	"go-kafka-clean-architecture/integration/testcontainers"
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

	mySqlDb, mySqlC, err := testcontainers.SetupSqlHandlerMySql(ctx)
	require.NoError(t, err)
	//defer mySqlC.Terminate(ctx)

	kafkaApi, zookeeperC, kafkaC, err := testcontainers.SetupEventApi(ctx)
	require.NoError(t, err)
	//defer zookeeperC.Terminate(ctx)
	//defer kafkaC.Terminate(ctx)

	serverURL := "http://localhost:" + strconv.Itoa(port.Int())

	restApi := rest_api.NewHttpApi(serverURL)

	r := registry.NewRegistry()
	httpContextRestApiSqlHandlerMySqlEventApiBrasilProductController := r.NewHttpContextRestApiSqlHandlerMySqlEventApiBrasilProductController(restApi, mySqlDb, kafkaApi)
	newHttpContextSqlHandlerMySqlProductTranslatedController := r.NewHttpContextSqlHandlerMySqlProductTranslatedController(mySqlDb)
	httpAppController := http_context.NewAppController(httpContextRestApiSqlHandlerMySqlEventApiBrasilProductController, newHttpContextSqlHandlerMySqlProductTranslatedController)

	eventContextSqlHandlerMySqlProductTranslatedController := r.NewEventContextSqlHandlerMySqlProductTranslatedController(mySqlDb)
	eventAppController := event_context.NewAppController(eventContextSqlHandlerMySqlProductTranslatedController)

	logger := logger.NewDebugLogger()

	kafkaConnectionString := "localhost:9092"
	go event_context_infrastructure.StartKafkaRouter(eventAppController, kafkaConnectionString, logger)
	go http_context_infrastructure.StartEchoRouter(httpAppController, port.Int(), logger)

	test.TestCreate(t, serverURL, int64(123))

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

	mySqlDb, mySqlC, err := testcontainers.SetupSqlHandlerMySql(ctx)
	require.NoError(t, err)
	//defer mySqlC.Terminate(ctx)

	serverURL := "http://localhost:" + strconv.Itoa(port.Int())

	r := registry.NewRegistry()
	httpContextRestApiSqlHandlerMySqlEventApiBrasilProductController := r.NewHttpContextRestApiSqlHandlerMySqlEventApiBrasilProductController(nil, mySqlDb, nil)
	newHttpContextSqlHandlerMySqlProductTranslatedController := r.NewHttpContextSqlHandlerMySqlProductTranslatedController(mySqlDb)
	httpAppController := http_context.NewAppController(httpContextRestApiSqlHandlerMySqlEventApiBrasilProductController, newHttpContextSqlHandlerMySqlProductTranslatedController)

	logger := logger.NewDebugLogger()

	go http_context_infrastructure.StartEchoRouter(httpAppController, port.Int(), logger)

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

	postgresDb, postgresC, err := testcontainers.SetupSqlHandlerPostgres(ctx)
	require.NoError(t, err)
	//defer postgresC.Terminate(ctx)

	serverURL := "http://localhost:" + strconv.Itoa(port.Int())

	r := registry.NewRegistry()
	httpContextRestApiSqlHandlerMySqlEventApiBrasilProductController := r.NewHttpContextRestApiSqlHandlerMySqlEventApiBrasilProductController(nil, postgresDb, nil)
	newHttpContextSqlHandlerMySqlProductTranslatedController := r.NewHttpContextSqlHandlerMySqlProductTranslatedController(postgresDb)
	httpAppController := http_context.NewAppController(httpContextRestApiSqlHandlerMySqlEventApiBrasilProductController, newHttpContextSqlHandlerMySqlProductTranslatedController)

	logger := logger.NewDebugLogger()

	go http_context_infrastructure.StartEchoRouter(httpAppController, port.Int(), logger)

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
