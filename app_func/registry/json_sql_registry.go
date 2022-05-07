package registry

import (
	"errors"
	"go-kafka-clean-architecture/app_func/infrastructure/logger"
	"go-kafka-clean-architecture/app_func/infrastructure/router/event_context"
	"go-kafka-clean-architecture/app_func/infrastructure/router/json_context"
)

func EventContextJsonContextRestApiSqlHandlerMySqlEventApiBrasil() {
	dbExec, dbQuery, _ := SqlHandlerMySql()
	restApiGet := RestApiGet()
	eventApiBind, eventApiWriteMessage := EventApiWriteMessage()

	eventContextSqlHandlerMySqlProductTranslatedControllerCreate := EventContextSqlHandlerMySqlProductTranslatedControllerCreate(dbExec)
	jsonContextRestApiSqlHandlerMySqlEventApiBrasilProductControllerCreate := JsonContextRestApiSqlHandlerMySqlEventApiBrasilProductControllerCreate(restApiGet, dbExec, eventApiBind, eventApiWriteMessage)
	jsonContextSqlHandlerMySqlProductControllerFindAll := JsonContextSqlHandlerMySqlProductControllerFindAll(dbQuery)
	jsonContextProductControllerGet := JsonContextProductControllerGet()
	jsonContextSqlHandlerMySqlProductTranslatedControllerFindAll := JsonContextSqlHandlerMySqlProductTranslatedControllerFindAll(dbQuery)

	log, err := logger.NewLogger()
	if !errors.Is(err, nil) {
		panic(err)
	}

	go event_context.StartKafkaRouter("localhost:9092")(logger.Error(log))(eventContextSqlHandlerMySqlProductTranslatedControllerCreate)

	err = json_context.StartEchoRouter(8080)(logger.Error(log))(jsonContextRestApiSqlHandlerMySqlEventApiBrasilProductControllerCreate)(jsonContextSqlHandlerMySqlProductControllerFindAll)(jsonContextProductControllerGet)(jsonContextSqlHandlerMySqlProductTranslatedControllerFindAll)
	if !errors.Is(err, nil) {
		panic(err)
	}
}

func EventContextJsonContextRestApiSqlHandlerPostgresEventApiBrasil() {
	_, dbQuery, queryRow := SqlHandlerPostgres()
	restApiGet := RestApiGet()
	eventApiBind, eventApiWriteMessage := EventApiWriteMessage()

	eventContextSqlHandlerPostgresProductTranslatedControllerCreate := EventContextSqlHandlerPostgresProductTranslatedControllerCreate(queryRow)
	jsonContextRestApiSqlHandlerPostgresBrasilProductControllerCreate := JsonContextRestApiSqlHandlerPostgresEventApiBrasilProductControllerCreate(restApiGet, queryRow, eventApiBind, eventApiWriteMessage)
	jsonContextSqlHandlerPostgresProductControllerFindAll := JsonContextSqlHandlerPostgresProductControllerFindAll(dbQuery)
	jsonContextProductControllerGet := JsonContextProductControllerGet()
	jsonContextSqlHandlerPostgresProductTranslatedControllerFindAll := JsonContextSqlHandlerPostgresProductTranslatedControllerFindAll(dbQuery)

	log, err := logger.NewLogger()
	if !errors.Is(err, nil) {
		panic(err)
	}

	go event_context.StartKafkaRouter("localhost:9092")(logger.Error(log))(eventContextSqlHandlerPostgresProductTranslatedControllerCreate)

	err = json_context.StartEchoRouter(8080)(logger.Error(log))(jsonContextRestApiSqlHandlerPostgresBrasilProductControllerCreate)(jsonContextSqlHandlerPostgresProductControllerFindAll)(jsonContextProductControllerGet)(jsonContextSqlHandlerPostgresProductTranslatedControllerFindAll)
	if !errors.Is(err, nil) {
		panic(err)
	}
}

func EventContextJsonContextRestApiSqlHandlerMySqlEventApiChile() {
	dbExec, dbQuery, _ := SqlHandlerMySql()
	restApiGet := RestApiGet()
	eventApiBind, eventApiWriteMessage := EventApiWriteMessage()

	eventContextSqlHandlerMySqlProductTranslatedControllerCreate := EventContextSqlHandlerMySqlProductTranslatedControllerCreate(dbExec)
	jsonContextRestApiSqlHandlerMySqlEventApiChileProductControllerCreate := JsonContextRestApiSqlHandlerMySqlEventApiChileProductControllerCreate(restApiGet, dbExec, eventApiBind, eventApiWriteMessage)
	jsonContextSqlHandlerMySqlProductControllerFindAll := JsonContextSqlHandlerMySqlProductControllerFindAll(dbQuery)
	jsonContextProductControllerGet := JsonContextProductControllerGet()
	jsonContextSqlHandlerMySqlProductTranslatedControllerFindAll := JsonContextSqlHandlerMySqlProductTranslatedControllerFindAll(dbQuery)

	log, err := logger.NewLogger()
	if !errors.Is(err, nil) {
		panic(err)
	}

	go event_context.StartKafkaRouter("localhost:9092")(logger.Error(log))(eventContextSqlHandlerMySqlProductTranslatedControllerCreate)

	err = json_context.StartEchoRouter(8080)(logger.Error(log))(jsonContextRestApiSqlHandlerMySqlEventApiChileProductControllerCreate)(jsonContextSqlHandlerMySqlProductControllerFindAll)(jsonContextProductControllerGet)(jsonContextSqlHandlerMySqlProductTranslatedControllerFindAll)
	if !errors.Is(err, nil) {
		panic(err)
	}
}

func EventContextJsonContextRestApiSqlHandlerPostgresEventApiChile() {
	_, dbQuery, queryRow := SqlHandlerPostgres()
	restApiGet := RestApiGet()
	eventApiBind, eventApiWriteMessage := EventApiWriteMessage()

	eventContextSqlHandlerPostgresProductTranslatedControllerCreate := EventContextSqlHandlerPostgresProductTranslatedControllerCreate(queryRow)
	jsonContextRestApiSqlHandlerPostgresEventApiChileProductControllerCreate := JsonContextRestApiSqlHandlerPostgresEventApiChileProductControllerCreate(restApiGet, queryRow, eventApiBind, eventApiWriteMessage)
	jsonContextSqlHandlerPostgresProductControllerFindAll := JsonContextSqlHandlerPostgresProductControllerFindAll(dbQuery)
	jsonContextProductControllerGet := JsonContextProductControllerGet()
	jsonContextSqlHandlerPostgresProductTranslatedControllerFindAll := JsonContextSqlHandlerPostgresProductTranslatedControllerFindAll(dbQuery)

	log, err := logger.NewLogger()
	if !errors.Is(err, nil) {
		panic(err)
	}

	go event_context.StartKafkaRouter("localhost:9092")(logger.Error(log))(eventContextSqlHandlerPostgresProductTranslatedControllerCreate)

	err = json_context.StartEchoRouter(8080)(logger.Error(log))(jsonContextRestApiSqlHandlerPostgresEventApiChileProductControllerCreate)(jsonContextSqlHandlerPostgresProductControllerFindAll)(jsonContextProductControllerGet)(jsonContextSqlHandlerPostgresProductTranslatedControllerFindAll)
	if !errors.Is(err, nil) {
		panic(err)
	}
}
