package integration

import (
	"fmt"
	"go-kafka-clean-architecture/app/command/controller/event_context"
	"go-kafka-clean-architecture/app/command/controller/http_context"
	"go-kafka-clean-architecture/app/infrastructure/api/rest_api"
	"go-kafka-clean-architecture/app/infrastructure/logger"
	event_context_infrastructure "go-kafka-clean-architecture/app/infrastructure/router/event_context"
	http_context_infrastructure "go-kafka-clean-architecture/app/infrastructure/router/http_context"
	"go-kafka-clean-architecture/app/interfaces/api"
	"go-kafka-clean-architecture/app/interfaces/database"
	"go-kafka-clean-architecture/app/registry"
	"go-kafka-clean-architecture/integration/gnomocktest"
	"go-kafka-clean-architecture/integration/test"
	"strconv"
	"testing"

	"github.com/docker/go-connections/nat"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"github.com/orlangure/gnomock"
	"github.com/orlangure/gnomock/preset/kafka"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type GnomockSqlHandlerMySql struct {
	suite.Suite

	mySqlDb database.SqlHandler
	mySqlC  *gnomock.Container

	kafkaApi api.EventApi
	kafkaC   *gnomock.Container

	serverURL string
	productID int64

	err error
}

func (suite *GnomockSqlHandlerMySql) SetupSuite() {
	suite.mySqlDb, suite.mySqlC, suite.err = gnomocktest.SetupSqlHandlerMySql()
	require.NoError(suite.T(), suite.err)

	suite.kafkaApi, suite.kafkaC, suite.err = gnomocktest.SetupEventApi()
	require.NoError(suite.T(), suite.err)

	port, _ := nat.NewPort("", strconv.Itoa(8080))
	suite.serverURL = "http://localhost:" + strconv.Itoa(port.Int())

	restApi := rest_api.NewHttpApi(suite.serverURL)

	r := registry.NewRegistry()
	httpContextRestApiSqlHandlerMySqlEventApiBrasilProductController := r.NewHttpContextRestApiSqlHandlerMySqlEventApiBrasilProductController(restApi, suite.mySqlDb, suite.kafkaApi)
	newHttpContextSqlHandlerMySqlProductTranslatedController := r.NewHttpContextSqlHandlerMySqlProductTranslatedController(suite.mySqlDb)
	httpAppController := http_context.NewAppController(httpContextRestApiSqlHandlerMySqlEventApiBrasilProductController, newHttpContextSqlHandlerMySqlProductTranslatedController)

	eventContextSqlHandlerMySqlProductTranslatedController := r.NewEventContextSqlHandlerMySqlProductTranslatedController(suite.mySqlDb)
	eventAppController := event_context.NewAppController(eventContextSqlHandlerMySqlProductTranslatedController)
	logger := logger.NewDebugLogger()

	kafkaConnectionString := suite.kafkaC.Address(kafka.BrokerPort)

	go event_context_infrastructure.StartKafkaRouter(eventAppController, kafkaConnectionString, logger)
	go http_context_infrastructure.StartEchoRouter(httpAppController, port.Int(), logger)

	suite.productID = int64(123)
}

func (suite *GnomockSqlHandlerMySql) TearDownSuite() {

	suite.err = gnomocktest.Stop(suite.mySqlC)
	require.NoError(suite.T(), suite.err)

	suite.err = gnomocktest.Stop(suite.kafkaC)
	require.NoError(suite.T(), suite.err)
}

func (suite *GnomockSqlHandlerMySql) TestCreate() {
	test.TestCreate(suite.T(), suite.serverURL, suite.productID)
}

func (suite *GnomockSqlHandlerMySql) TestFindAll() {

	productType := fmt.Sprintf("Type %d", suite.productID)
	productName := fmt.Sprintf("Name %d", suite.productID)

	productTranslatedType := fmt.Sprintf("Type %d Brasil", suite.productID)
	productTranslatedName := fmt.Sprintf("Name %d Brasil", suite.productID)

	test.TestFindAll(suite.T(), suite.serverURL, &suite.productID, &productType, &productName, &suite.productID, &productTranslatedType, &productTranslatedName)
}

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(GnomockSqlHandlerMySql))
}
