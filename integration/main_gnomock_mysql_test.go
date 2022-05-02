package integration

import (
	"go-kafka-clean-architecture/app/infrastructure/api/rest_api"
	"go-kafka-clean-architecture/app/infrastructure/logger"
	event_context_infrastructure "go-kafka-clean-architecture/app/infrastructure/router/event_context"
	http_context_infrastructure "go-kafka-clean-architecture/app/infrastructure/router/http_context"
	"go-kafka-clean-architecture/app/interfaces/api"
	"go-kafka-clean-architecture/app/interfaces/database"
	"go-kafka-clean-architecture/integration/gnomocktest"
	"go-kafka-clean-architecture/integration/test"
	"go-kafka-clean-architecture/registry"
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

type GnomockSQLHandlerMySQL struct {
	suite.Suite

	mySqlDb database.SQLHandler
	mySqlC  *gnomock.Container

	kafkaAPI api.EventAPI
	kafkaC   *gnomock.Container

	serverURL string

	err error
}

func (suite *GnomockSQLHandlerMySQL) SetupSuite() {
	suite.mySqlDb, suite.mySqlC, suite.err = gnomocktest.SetupSQLHandlerMySQL()
	require.NoError(suite.T(), suite.err)

	suite.kafkaAPI, suite.kafkaC, suite.err = gnomocktest.SetupEventAPI()
	require.NoError(suite.T(), suite.err)

	port, _ := nat.NewPort("", strconv.Itoa(8080))
	suite.serverURL = "http://localhost:" + strconv.Itoa(port.Int())

	restAPI := rest_api.NewHttpAPI(suite.serverURL)

	r := registry.NewRegistry()
	httpContextAppController := r.NewHttpContextRestSqlEventAppControllerMySql(restAPI, suite.mySqlDb, suite.kafkaAPI)
	eventContextAppController := r.NewEventContextRestSqlEventAppControllerMySql(restAPI, suite.mySqlDb, suite.kafkaAPI)
	logger := logger.NewDebugLogger()

	kafkaConnectionString := suite.kafkaC.Address(kafka.BrokerPort)
	go event_context_infrastructure.StartKafkaRouter(eventContextAppController, kafkaConnectionString, logger)
	go http_context_infrastructure.StartEchoRouter(httpContextAppController, port.Int(), logger)
}

func (suite *GnomockSQLHandlerMySQL) TearDownSuite() {

	suite.err = gnomocktest.Stop(suite.mySqlC)
	require.NoError(suite.T(), suite.err)

	suite.err = gnomocktest.Stop(suite.kafkaC)
	require.NoError(suite.T(), suite.err)
}

func (suite *GnomockSQLHandlerMySQL) TestCreate() {
	test.TestCreate(suite.T(), suite.serverURL)
}

func (suite *GnomockSQLHandlerMySQL) TestFindAll() {
	productID := int64(123)
	productType := "Type 123"
	productName := "Name 123"

	productTranslatedID := int64(123)
	productTranslatedType := "Type 123 Brasil"
	productTranslatedName := "Name 123 Brasil"

	test.TestFindAll(suite.T(), suite.serverURL, &productID, &productType, &productName, &productTranslatedID, &productTranslatedType, &productTranslatedName)
}

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(GnomockSQLHandlerMySQL))
}
