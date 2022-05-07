package registry

import (
	"errors"
	"go-kafka-clean-architecture/app_func/infrastructure/logger"
	"go-kafka-clean-architecture/app_func/infrastructure/router/event_context"
	"go-kafka-clean-architecture/app_func/infrastructure/router/http_context"
)

func EventContextHttpContextRestApiSqlHandlerMySqlEventApiBrasil() {
	dbExec, dbQuery, _ := SqlHandlerMySql()
	restApiGet := RestApiGet()
	eventApiBind, eventApiWriteMessage := EventApiWriteMessage()

	eventContextSqlHandlerMySqlProductTranslatedControllerCreate := EventContextSqlHandlerMySqlProductTranslatedControllerCreate(dbExec)
	httpContextRestApiSqlHandlerMySqlEventApiBrasilProductControllerCreate := HttpContextRestApiSqlHandlerMySqlEventApiBrasilProductControllerCreate(restApiGet, dbExec, eventApiBind, eventApiWriteMessage)
	httpContextSqlHandlerMySqlProductControllerFindAll := HttpContextSqlHandlerMySqlProductControllerFindAll(dbQuery)
	httpContextProductControllerGet := HttpContextProductControllerGet()
	httpContextSqlHandlerMySqlProductTranslatedControllerFindAll := HttpContextSqlHandlerMySqlProductTranslatedControllerFindAll(dbQuery)

	log, err := logger.NewLogger()
	if !errors.Is(err, nil) {
		panic(err)
	}

	go event_context.StartKafkaRouter("localhost:9092")(logger.Error(log))(eventContextSqlHandlerMySqlProductTranslatedControllerCreate)

	err = http_context.StartEchoRouter(8080)(logger.Error(log))(httpContextRestApiSqlHandlerMySqlEventApiBrasilProductControllerCreate)(httpContextSqlHandlerMySqlProductControllerFindAll)(httpContextProductControllerGet)(httpContextSqlHandlerMySqlProductTranslatedControllerFindAll)
	if !errors.Is(err, nil) {
		panic(err)
	}
}

func EventContextHttpContextRestApiSqlHandlerPostgresEventApiBrasil() {
	_, dbQuery, queryRow := SqlHandlerPostgres()
	restApiGet := RestApiGet()
	eventApiBind, eventApiWriteMessage := EventApiWriteMessage()

	eventContextSqlHandlerPostgresProductTranslatedControllerCreate := EventContextSqlHandlerPostgresProductTranslatedControllerCreate(queryRow)
	httpContextRestApiSqlHandlerPostgresBrasilProductControllerCreate := HttpContextRestApiSqlHandlerPostgresEventApiBrasilProductControllerCreate(restApiGet, queryRow, eventApiBind, eventApiWriteMessage)
	httpContextSqlHandlerPostgresProductControllerFindAll := HttpContextSqlHandlerPostgresProductControllerFindAll(dbQuery)
	httpContextProductControllerGet := HttpContextProductControllerGet()
	httpContextSqlHandlerPostgresProductTranslatedControllerFindAll := HttpContextSqlHandlerPostgresProductTranslatedControllerFindAll(dbQuery)

	log, err := logger.NewLogger()
	if !errors.Is(err, nil) {
		panic(err)
	}

	go event_context.StartKafkaRouter("localhost:9092")(logger.Error(log))(eventContextSqlHandlerPostgresProductTranslatedControllerCreate)

	err = http_context.StartEchoRouter(8080)(logger.Error(log))(httpContextRestApiSqlHandlerPostgresBrasilProductControllerCreate)(httpContextSqlHandlerPostgresProductControllerFindAll)(httpContextProductControllerGet)(httpContextSqlHandlerPostgresProductTranslatedControllerFindAll)
	if !errors.Is(err, nil) {
		panic(err)
	}
}

func EventContextHttpContextRestApiSqlHandlerMySqlEventApiChile() {
	dbExec, dbQuery, _ := SqlHandlerMySql()
	restApiGet := RestApiGet()
	eventApiBind, eventApiWriteMessage := EventApiWriteMessage()

	eventContextSqlHandlerMySqlProductTranslatedControllerCreate := EventContextSqlHandlerMySqlProductTranslatedControllerCreate(dbExec)
	httpContextRestApiSqlHandlerMySqlEventApiChileProductControllerCreate := HttpContextRestApiSqlHandlerMySqlEventApiChileProductControllerCreate(restApiGet, dbExec, eventApiBind, eventApiWriteMessage)
	httpContextSqlHandlerMySqlProductControllerFindAll := HttpContextSqlHandlerMySqlProductControllerFindAll(dbQuery)
	httpContextProductControllerGet := HttpContextProductControllerGet()
	httpContextSqlHandlerMySqlProductTranslatedControllerFindAll := HttpContextSqlHandlerMySqlProductTranslatedControllerFindAll(dbQuery)

	log, err := logger.NewLogger()
	if !errors.Is(err, nil) {
		panic(err)
	}

	go event_context.StartKafkaRouter("localhost:9092")(logger.Error(log))(eventContextSqlHandlerMySqlProductTranslatedControllerCreate)

	err = http_context.StartEchoRouter(8080)(logger.Error(log))(httpContextRestApiSqlHandlerMySqlEventApiChileProductControllerCreate)(httpContextSqlHandlerMySqlProductControllerFindAll)(httpContextProductControllerGet)(httpContextSqlHandlerMySqlProductTranslatedControllerFindAll)
	if !errors.Is(err, nil) {
		panic(err)
	}
}

func EventContextHttpContextRestApiSqlHandlerPostgresEventApiChile() {
	_, dbQuery, queryRow := SqlHandlerPostgres()
	restApiGet := RestApiGet()
	eventApiBind, eventApiWriteMessage := EventApiWriteMessage()

	eventContextSqlHandlerPostgresProductTranslatedControllerCreate := EventContextSqlHandlerPostgresProductTranslatedControllerCreate(queryRow)
	httpContextRestApiSqlHandlerPostgresEventApiChileProductControllerCreate := HttpContextRestApiSqlHandlerPostgresEventApiChileProductControllerCreate(restApiGet, queryRow, eventApiBind, eventApiWriteMessage)
	httpContextSqlHandlerPostgresProductControllerFindAll := HttpContextSqlHandlerPostgresProductControllerFindAll(dbQuery)
	httpContextProductControllerGet := HttpContextProductControllerGet()
	httpContextSqlHandlerPostgresProductTranslatedControllerFindAll := HttpContextSqlHandlerPostgresProductTranslatedControllerFindAll(dbQuery)

	log, err := logger.NewLogger()
	if !errors.Is(err, nil) {
		panic(err)
	}

	go event_context.StartKafkaRouter("localhost:9092")(logger.Error(log))(eventContextSqlHandlerPostgresProductTranslatedControllerCreate)

	err = http_context.StartEchoRouter(8080)(logger.Error(log))(httpContextRestApiSqlHandlerPostgresEventApiChileProductControllerCreate)(httpContextSqlHandlerPostgresProductControllerFindAll)(httpContextProductControllerGet)(httpContextSqlHandlerPostgresProductTranslatedControllerFindAll)
	if !errors.Is(err, nil) {
		panic(err)
	}
}
