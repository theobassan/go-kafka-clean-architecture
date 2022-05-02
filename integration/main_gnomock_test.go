package integration

import (
	"go-kafka-clean-architecture/app/infrastructure/api/rest_api"
	"go-kafka-clean-architecture/app/infrastructure/logger"
	event_context_infrastructure "go-kafka-clean-architecture/app/infrastructure/router/event_context"
	http_context_infrastructure "go-kafka-clean-architecture/app/infrastructure/router/http_context"
	"go-kafka-clean-architecture/app/interfaces/repository/sql_gorm/model"
	"go-kafka-clean-architecture/integration/gnomocktest"
	"go-kafka-clean-architecture/integration/test"
	"go-kafka-clean-architecture/registry"
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

	mySqlDb, mySqlC, err := gnomocktest.SetupSQLHandlerMySQL()
	require.NoError(t, err)
	//defer gnomocktest.Stop(mySqlC)

	kafkaAPI, kafkaC, err := gnomocktest.SetupEventAPI()
	require.NoError(t, err)
	//defer gnomocktest.Stop(kafkaC)

	serverURL := "http://localhost:" + strconv.Itoa(port.Int())

	restAPI := rest_api.NewHttpAPI(serverURL)

	r := registry.NewRegistry()
	httpContextAppController := r.NewHttpContextRestSqlEventAppControllerMySql(restAPI, mySqlDb, kafkaAPI)
	eventContextAppController := r.NewEventContextRestSqlEventAppControllerMySql(restAPI, mySqlDb, kafkaAPI)
	logger := logger.NewDebugLogger()

	kafkaConnectionString := kafkaC.Address(kafka.BrokerPort)
	go event_context_infrastructure.StartKafkaRouter(eventContextAppController, kafkaConnectionString, logger)
	go http_context_infrastructure.StartEchoRouter(httpContextAppController, port.Int(), logger)

	test.TestCreate(t, serverURL)

	err = gnomocktest.Stop(mySqlC)
	require.NoError(t, err)

	err = gnomocktest.Stop(kafkaC)
	require.NoError(t, err)
}

func TestCreate_gnomock_postgres_sql_handler(t *testing.T) {
	port, _ := nat.NewPort("", strconv.Itoa(8080))

	postgresDb, postgresC, err := gnomocktest.SetupSQLHandlerPostgres()
	require.NoError(t, err)
	//defer gnomocktest.Stop(postgresC)

	kafkaAPI, kafkaC, err := gnomocktest.SetupEventAPI()
	require.NoError(t, err)
	//defer gnomocktest.Stop(kafkaC)

	serverURL := "http://localhost:" + strconv.Itoa(port.Int())

	restAPI := rest_api.NewHttpAPI(serverURL)

	r := registry.NewRegistry()
	httpContextAppController := r.NewHttpContextRestSqlEventAppControllerPostgres(restAPI, postgresDb, kafkaAPI)
	eventContextAppController := r.NewEventContextRestSqlEventAppControllerPostgres(restAPI, postgresDb, kafkaAPI)
	logger := logger.NewDebugLogger()

	kafkaConnectionString := kafkaC.Address(kafka.BrokerPort)
	go event_context_infrastructure.StartKafkaRouter(eventContextAppController, kafkaConnectionString, logger)
	go http_context_infrastructure.StartEchoRouter(httpContextAppController, port.Int(), logger)

	test.TestCreate(t, serverURL)

	err = gnomocktest.Stop(postgresC)
	require.NoError(t, err)

	err = gnomocktest.Stop(kafkaC)
	require.NoError(t, err)
}

func TestFindAll_gnomock_mqsql_sql_handler(t *testing.T) {
	port, _ := nat.NewPort("", strconv.Itoa(8080))

	mySqlDb, mySqlC, err := gnomocktest.SetupSQLHandlerMySQL()
	require.NoError(t, err)
	//defer gnomocktest.Stop(mySqlC)

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

	err = gnomocktest.Stop(mySqlC)
	require.NoError(t, err)
}

func TestFindAll_gnomock_postgres_sql_handler(t *testing.T) {
	port, _ := nat.NewPort("", strconv.Itoa(8080))

	postgresDb, postgresC, err := gnomocktest.SetupSQLHandlerPostgres()
	require.NoError(t, err)
	//defer gnomocktest.Stop(postgresC)

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

	err = gnomocktest.Stop(postgresC)
	require.NoError(t, err)
}

func TestFindAll_gnomock_postgres_sql_gorm(t *testing.T) {
	port, _ := nat.NewPort("", strconv.Itoa(8080))

	postgresDb, postgresC, err := gnomocktest.SetupSQLGormPostgres()
	require.NoError(t, err)
	//defer gnomocktest.Stop(postgresC)

	serverURL := "http://localhost:" + strconv.Itoa(port.Int())

	r := registry.NewRegistry()
	httpContextAppController := r.NewHttpContextRestGormEventAppController(nil, postgresDb, nil)
	logger := logger.NewDebugLogger()

	go http_context_infrastructure.StartEchoRouter(httpContextAppController, port.Int(), logger)

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
