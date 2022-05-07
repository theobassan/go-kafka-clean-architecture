package integration

import (
	"go-kafka-clean-architecture/app/command/controller/event_context"
	"go-kafka-clean-architecture/app/command/controller/http_context"
	"go-kafka-clean-architecture/app/infrastructure/api/rest_api"
	"go-kafka-clean-architecture/app/infrastructure/logger"
	event_context_infrastructure "go-kafka-clean-architecture/app/infrastructure/router/event_context"
	http_context_infrastructure "go-kafka-clean-architecture/app/infrastructure/router/http_context"
	"go-kafka-clean-architecture/app/interfaces/repository/sql_gorm/model"
	"go-kafka-clean-architecture/app/registry"
	"go-kafka-clean-architecture/integration/gnomocktest"
	"go-kafka-clean-architecture/integration/test"
	"strconv"
	"testing"

	"github.com/docker/go-connections/nat"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"github.com/orlangure/gnomock/preset/kafka"
	"github.com/stretchr/testify/require"
)

func TestCreate_gnomock_mysql_sql_handler(t *testing.T) {
	port, _ := nat.NewPort("", strconv.Itoa(8080))

	mySqlDb, mySqlC, err := gnomocktest.SetupSqlHandlerMySql()
	require.NoError(t, err)
	//defer gnomocktest.Stop(mySqlC)

	kafkaApi, kafkaC, err := gnomocktest.SetupEventApi()
	require.NoError(t, err)
	//defer gnomocktest.Stop(kafkaC)

	serverURL := "http://localhost:" + strconv.Itoa(port.Int())

	restApi := rest_api.NewHttpApi(serverURL)

	r := registry.NewRegistry()
	httpContextRestApiSqlHandlerMySqlEventApiBrasilProductController := r.NewHttpContextRestApiSqlHandlerMySqlEventApiBrasilProductController(restApi, mySqlDb, kafkaApi)
	newHttpContextSqlHandlerMySqlProductTranslatedController := r.NewHttpContextSqlHandlerMySqlProductTranslatedController(mySqlDb)
	httpAppController := http_context.NewAppController(httpContextRestApiSqlHandlerMySqlEventApiBrasilProductController, newHttpContextSqlHandlerMySqlProductTranslatedController)

	eventContextSqlHandlerMySqlProductTranslatedController := r.NewEventContextSqlHandlerMySqlProductTranslatedController(mySqlDb)
	eventAppController := event_context.NewAppController(eventContextSqlHandlerMySqlProductTranslatedController)

	logger := logger.NewDebugLogger()

	kafkaConnectionString := kafkaC.Address(kafka.BrokerPort)
	go event_context_infrastructure.StartKafkaRouter(eventAppController, kafkaConnectionString, logger)
	go http_context_infrastructure.StartEchoRouter(httpAppController, port.Int(), logger)

	test.TestCreate(t, serverURL, int64(123))

	err = gnomocktest.Stop(mySqlC)
	require.NoError(t, err)

	err = gnomocktest.Stop(kafkaC)
	require.NoError(t, err)
}

func TestCreate_gnomock_postgres_sql_handler(t *testing.T) {
	port, _ := nat.NewPort("", strconv.Itoa(8080))

	postgresDb, postgresC, err := gnomocktest.SetupSqlHandlerPostgres()
	require.NoError(t, err)
	//defer gnomocktest.Stop(postgresC)

	kafkaApi, kafkaC, err := gnomocktest.SetupEventApi()
	require.NoError(t, err)
	//defer gnomocktest.Stop(kafkaC)

	serverURL := "http://localhost:" + strconv.Itoa(port.Int())

	restApi := rest_api.NewHttpApi(serverURL)

	r := registry.NewRegistry()
	httpContextRestApiSqlHandlerPostgresEventApiBrasilProductController := r.NewHttpContextRestApiSqlHandlerPostgresEventApiBrasilProductController(restApi, postgresDb, kafkaApi)
	newHttpContextSqlHandlerPostgresProductTranslatedController := r.NewHttpContextSqlHandlerPostgresProductTranslatedController(postgresDb)
	httpAppController := http_context.NewAppController(httpContextRestApiSqlHandlerPostgresEventApiBrasilProductController, newHttpContextSqlHandlerPostgresProductTranslatedController)

	eventContextSqlHandlerPostgresProductTranslatedController := r.NewEventContextSqlHandlerPostgresProductTranslatedController(postgresDb)
	eventAppController := event_context.NewAppController(eventContextSqlHandlerPostgresProductTranslatedController)

	logger := logger.NewDebugLogger()

	kafkaConnectionString := kafkaC.Address(kafka.BrokerPort)
	go event_context_infrastructure.StartKafkaRouter(eventAppController, kafkaConnectionString, logger)
	go http_context_infrastructure.StartEchoRouter(httpAppController, port.Int(), logger)

	test.TestCreate(t, serverURL, int64(123))

	err = gnomocktest.Stop(postgresC)
	require.NoError(t, err)

	err = gnomocktest.Stop(kafkaC)
	require.NoError(t, err)
}

func TestFindAll_gnomock_mqsql_sql_handler(t *testing.T) {
	port, _ := nat.NewPort("", strconv.Itoa(8080))

	mySqlDb, mySqlC, err := gnomocktest.SetupSqlHandlerMySql()
	require.NoError(t, err)
	//defer gnomocktest.Stop(mySqlC)

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

	err = gnomocktest.Stop(mySqlC)
	require.NoError(t, err)
}

func TestFindAll_gnomock_postgres_sql_handler(t *testing.T) {
	port, _ := nat.NewPort("", strconv.Itoa(8080))

	postgresDb, postgresC, err := gnomocktest.SetupSqlHandlerPostgres()
	require.NoError(t, err)
	//defer gnomocktest.Stop(postgresC)

	serverURL := "http://localhost:" + strconv.Itoa(port.Int())

	r := registry.NewRegistry()
	httpContextRestApiSqlHandlerPostgresEventApiBrasilProductController := r.NewHttpContextRestApiSqlHandlerPostgresEventApiBrasilProductController(nil, postgresDb, nil)
	newHttpContextSqlHandlerPostgresProductTranslatedController := r.NewHttpContextSqlHandlerPostgresProductTranslatedController(postgresDb)
	httpAppController := http_context.NewAppController(httpContextRestApiSqlHandlerPostgresEventApiBrasilProductController, newHttpContextSqlHandlerPostgresProductTranslatedController)

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

	err = gnomocktest.Stop(postgresC)
	require.NoError(t, err)
}

func TestFindAll_gnomock_postgres_sql_gorm(t *testing.T) {
	port, _ := nat.NewPort("", strconv.Itoa(8080))

	postgresDb, postgresC, err := gnomocktest.SetupSqlGormPostgres()
	require.NoError(t, err)
	//defer gnomocktest.Stop(postgresC)

	serverURL := "http://localhost:" + strconv.Itoa(port.Int())

	r := registry.NewRegistry()
	httpContextRestApiSqlGormEventApiBrasilProductController := r.NewHttpContextRestApiSqlGormEventApiBrasilProductController(nil, postgresDb, nil)
	newHttpContextSqlGormProductTranslatedController := r.NewHttpContextSqlGormProductTranslatedController(postgresDb)
	httpAppController := http_context.NewAppController(httpContextRestApiSqlGormEventApiBrasilProductController, newHttpContextSqlGormProductTranslatedController)

	logger := logger.NewDebugLogger()

	go http_context_infrastructure.StartEchoRouter(httpAppController, port.Int(), logger)

	productID := int64(123)
	productType := "Type"
	productName := "Name"

	modelProduct := model.Product{
		ExternalID: &productID,
		Type:       &productType,
		Name:       &productName,
	}
	err = postgresDb.Create(&modelProduct).Error
	require.NoError(t, err)

	modelTranslatedProduct := model.ProductTranslated{
		ExternalID: &productID,
		Type:       &productType,
		Name:       &productName,
	}
	err = postgresDb.Create(&modelTranslatedProduct).Error
	require.NoError(t, err)

	test.TestFindAll(t, serverURL, &productID, &productType, &productName, &productID, &productType, &productName)

	err = gnomocktest.Stop(postgresC)
	require.NoError(t, err)
}
