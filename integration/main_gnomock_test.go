package integration

import (
	"go-kafka-clean-architecture/app/infrastructure/api/rest_api"
	event_context_infrastructure "go-kafka-clean-architecture/app/infrastructure/router/event_context"
	http_context_infrastructure "go-kafka-clean-architecture/app/infrastructure/router/http_context"
	"go-kafka-clean-architecture/integration/gnomock"
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

func TestCreate_gnomock(t *testing.T) {
	port, _ := nat.NewPort("", strconv.Itoa(8080))

	mySqlDb, mySqlC, err := gnomock.SetupSQLHandlerMySQL()
	require.NoError(t, err)
	//defer gnomock.Stop(mySqlC)

	kafkaAPI, kafkaC, err := gnomock.SetupEventAPI()
	require.NoError(t, err)
	//defer gnomock.Stop(kafkaC)

	serverURL := "http://localhost:" + strconv.Itoa(port.Int())

	restAPI := rest_api.NewHttpAPI(serverURL)

	r := registry.NewRegistry()
	httpContextAppController := r.NewHttpContextRestSqlEventAppController(restAPI, mySqlDb, kafkaAPI)
	eventContextAppController := r.NewEventContextRestSqlEventAppController(restAPI, mySqlDb, kafkaAPI)

	kafkaConnectionString := kafkaC.Address(kafka.BrokerPort)
	go event_context_infrastructure.StartKafkaRouter(eventContextAppController, kafkaConnectionString)
	go http_context_infrastructure.StartEchoRouter(httpContextAppController, port.Int())

	test.TestCreate(t, serverURL)

	err = gnomock.Stop(mySqlC)
	require.NoError(t, err)

	err = gnomock.Stop(kafkaC)
	require.NoError(t, err)
}

func TestFindAll_gnomock(t *testing.T) {
	port, _ := nat.NewPort("", strconv.Itoa(8080))

	mySqlDb, mySqlC, err := gnomock.SetupSQLHandlerMySQL()
	require.NoError(t, err)
	//defer gnomock.Stop(mySqlC)

	kafkaAPI, kafkaC, err := gnomock.SetupEventAPI()
	require.NoError(t, err)
	//defer gnomock.Stop(kafkaC)

	serverURL := "http://localhost:" + strconv.Itoa(port.Int())

	restAPI := rest_api.NewHttpAPI(serverURL)

	r := registry.NewRegistry()
	httpContextAppController := r.NewHttpContextRestSqlEventAppController(restAPI, mySqlDb, kafkaAPI)
	eventContextAppController := r.NewEventContextRestSqlEventAppController(restAPI, mySqlDb, kafkaAPI)

	kafkaConnectionString := kafkaC.Address(kafka.BrokerPort)
	go event_context_infrastructure.StartKafkaRouter(eventContextAppController, kafkaConnectionString)
	go http_context_infrastructure.StartEchoRouter(httpContextAppController, port.Int())

	test.TestFindAll(t, serverURL)

	err = gnomock.Stop(mySqlC)
	require.NoError(t, err)

	err = gnomock.Stop(kafkaC)
	require.NoError(t, err)
}
