package registry

import (
	"errors"
	"go-kafka-clean-architecture/app_func/infrastructure/logger"
	"go-kafka-clean-architecture/app_func/infrastructure/router/event_context"
	"go-kafka-clean-architecture/app_func/infrastructure/router/json_context"
)

func EventContextJsonContextRestApiSqlGormMySqlEventApiBrasil() {
	dbCreate, dbFind := SqlGormMySql()
	restApiGet := RestApiGet()
	eventApiBind, eventApiWriteMessage := EventApiWriteMessage()

	eventContextSqlGormProductTranslatedControllerCreate := EventContextSqlGormProductTranslatedControllerCreate(dbCreate)
	jsonContextRestApiSqlGormEventApiBrasilProductControllerCreate := JsonContextRestApiSqlGormEventApiBrasilProductControllerCreate(restApiGet, dbCreate, eventApiBind, eventApiWriteMessage)
	jsonContextSqlGormProductControllerFindAll := JsonContextSqlGormProductControllerFindAll(dbFind)
	jsonContextProductControllerGet := JsonContextProductControllerGet()
	jsonContextSqlGormProductTranslatedControllerFindAll := JsonContextSqlGormProductTranslatedControllerFindAll(dbFind)

	log, err := logger.NewLogger()
	if !errors.Is(err, nil) {
		panic(err)
	}

	go event_context.StartKafkaRouter("localhost:9092")(logger.Error(log))(eventContextSqlGormProductTranslatedControllerCreate)

	err = json_context.StartEchoRouter(8080)(logger.Error(log))(jsonContextRestApiSqlGormEventApiBrasilProductControllerCreate)(jsonContextSqlGormProductControllerFindAll)(jsonContextProductControllerGet)(jsonContextSqlGormProductTranslatedControllerFindAll)
	if !errors.Is(err, nil) {
		panic(err)
	}
}

func EventContextJsonContextRestApiSqlGormPostgresEventApiBrasil() {
	dbCreate, dbFind := SqlGormPostgres()
	restApiGet := RestApiGet()
	eventApiBind, eventApiWriteMessage := EventApiWriteMessage()

	eventContextSqlGormProductTranslatedControllerCreate := EventContextSqlGormProductTranslatedControllerCreate(dbCreate)
	jsonContextRestApiSqlGormEventApiBrasilProductControllerCreate := JsonContextRestApiSqlGormEventApiBrasilProductControllerCreate(restApiGet, dbCreate, eventApiBind, eventApiWriteMessage)
	jsonContextSqlGormProductControllerFindAll := JsonContextSqlGormProductControllerFindAll(dbFind)
	jsonContextProductControllerGet := JsonContextProductControllerGet()
	jsonContextSqlGormProductTranslatedControllerFindAll := JsonContextSqlGormProductTranslatedControllerFindAll(dbFind)

	log, err := logger.NewLogger()
	if !errors.Is(err, nil) {
		panic(err)
	}

	go event_context.StartKafkaRouter("localhost:9092")(logger.Error(log))(eventContextSqlGormProductTranslatedControllerCreate)

	err = json_context.StartEchoRouter(8080)(logger.Error(log))(jsonContextRestApiSqlGormEventApiBrasilProductControllerCreate)(jsonContextSqlGormProductControllerFindAll)(jsonContextProductControllerGet)(jsonContextSqlGormProductTranslatedControllerFindAll)
	if !errors.Is(err, nil) {
		panic(err)
	}
}

func EventContextJsonContextRestApiSqlGormMySqlEventApiChile() {
	dbCreate, dbFind := SqlGormMySql()
	restApiGet := RestApiGet()
	eventApiBind, eventApiWriteMessage := EventApiWriteMessage()

	eventContextSqlGormProductTranslatedControllerCreate := EventContextSqlGormProductTranslatedControllerCreate(dbCreate)
	jsonContextRestApiSqlGormEventApiChileProductControllerCreate := JsonContextRestApiSqlGormEventApiChileProductControllerCreate(restApiGet, dbCreate, eventApiBind, eventApiWriteMessage)
	jsonContextSqlGormProductControllerFindAll := JsonContextSqlGormProductControllerFindAll(dbFind)
	jsonContextProductControllerGet := JsonContextProductControllerGet()
	jsonContextSqlGormProductTranslatedControllerFindAll := JsonContextSqlGormProductTranslatedControllerFindAll(dbFind)

	log, err := logger.NewLogger()
	if !errors.Is(err, nil) {
		panic(err)
	}

	go event_context.StartKafkaRouter("localhost:9092")(logger.Error(log))(eventContextSqlGormProductTranslatedControllerCreate)

	err = json_context.StartEchoRouter(8080)(logger.Error(log))(jsonContextRestApiSqlGormEventApiChileProductControllerCreate)(jsonContextSqlGormProductControllerFindAll)(jsonContextProductControllerGet)(jsonContextSqlGormProductTranslatedControllerFindAll)
	if !errors.Is(err, nil) {
		panic(err)
	}
}

func EventContextJsonContextRestApiSqlGormPostgresEventApiChile() {
	dbCreate, dbFind := SqlGormPostgres()
	restApiGet := RestApiGet()
	eventApiBind, eventApiWriteMessage := EventApiWriteMessage()

	eventContextSqlGormProductTranslatedControllerCreate := EventContextSqlGormProductTranslatedControllerCreate(dbCreate)
	jsonContextRestApiSqlGormEventApiChileProductControllerCreate := JsonContextRestApiSqlGormEventApiChileProductControllerCreate(restApiGet, dbCreate, eventApiBind, eventApiWriteMessage)
	jsonContextSqlGormProductControllerFindAll := JsonContextSqlGormProductControllerFindAll(dbFind)
	jsonContextProductControllerGet := JsonContextProductControllerGet()
	jsonContextSqlGormProductTranslatedControllerFindAll := JsonContextSqlGormProductTranslatedControllerFindAll(dbFind)

	log, err := logger.NewLogger()
	if !errors.Is(err, nil) {
		panic(err)
	}

	go event_context.StartKafkaRouter("localhost:9092")(logger.Error(log))(eventContextSqlGormProductTranslatedControllerCreate)

	err = json_context.StartEchoRouter(8080)(logger.Error(log))(jsonContextRestApiSqlGormEventApiChileProductControllerCreate)(jsonContextSqlGormProductControllerFindAll)(jsonContextProductControllerGet)(jsonContextSqlGormProductTranslatedControllerFindAll)
	if !errors.Is(err, nil) {
		panic(err)
	}
}
