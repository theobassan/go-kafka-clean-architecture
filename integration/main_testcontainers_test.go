package integration

import (
	"context"
	"go-kafka-clean-architecture/app/infrastructure/api/rest_api"
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

func TestCreate_testcontainers(t *testing.T) {
	port, _ := nat.NewPort("", strconv.Itoa(8080))
	ctx := context.Background()

	mySqlDb, mySqlC, err := testcontainers.SetupSQLHandlerMySql(ctx)
	require.NoError(t, err)
	//defer mySqlC.Terminate()

	kafkaAPI, kafkaC, err := testcontainers.SetupEventAPICompose(ctx)
	require.NoError(t, err)
	//defer kafkaC.Down()

	serverURL := "http://localhost:" + strconv.Itoa(port.Int())

	restAPI := rest_api.NewHttpAPI(serverURL)

	r := registry.NewRegistry()
	httpContextAppController := r.NewHttpContextRestSqlEventAppController(restAPI, mySqlDb, kafkaAPI)
	eventContextAppController := r.NewEventContextRestSqlEventAppController(restAPI, mySqlDb, kafkaAPI)

	go event_context_infrastructure.StartKafkaRouter(eventContextAppController, "localhost:9092")
	go http_context_infrastructure.StartEchoRouter(httpContextAppController, port.Int())

	//test.TestCreate(t, serverURL)

	err = mySqlC.Terminate(ctx)
	require.NoError(t, err)

	err = kafkaC.Down().Error
	require.NoError(t, err)
}

func TestFindAll_testcontainers(t *testing.T) {
	port, _ := nat.NewPort("", strconv.Itoa(8080))
	ctx := context.Background()

	mySqlDb, mySqlC, err := testcontainers.SetupSQLHandlerMySql(ctx)
	require.NoError(t, err)
	//defer mySqlC.Terminate()

	kafkaAPI, kafkaC, err := testcontainers.SetupEventAPICompose(ctx)
	require.NoError(t, err)
	//defer kafkaC.Down()

	serverURL := "http://localhost:" + strconv.Itoa(port.Int())

	restAPI := rest_api.NewHttpAPI(serverURL)

	r := registry.NewRegistry()
	httpContextAppController := r.NewHttpContextRestSqlEventAppController(restAPI, mySqlDb, kafkaAPI)
	eventContextAppController := r.NewEventContextRestSqlEventAppController(restAPI, mySqlDb, kafkaAPI)

	go event_context_infrastructure.StartKafkaRouter(eventContextAppController, "localhost:9092")
	go http_context_infrastructure.StartEchoRouter(httpContextAppController, port.Int())

	test.TestFindAll(t, serverURL)

	err = mySqlC.Terminate(ctx)
	require.NoError(t, err)

	err = kafkaC.Down().Error
	require.NoError(t, err)
}
