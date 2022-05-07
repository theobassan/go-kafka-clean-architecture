package registry

import (
	"errors"
	"go-kafka-clean-architecture/app_func/infrastructure/logger"
	"go-kafka-clean-architecture/app_func/infrastructure/router/event_context"
	"go-kafka-clean-architecture/app_func/infrastructure/router/http_context"
)

func EventContextHttpContextRestApiSqlGormMySqlEventApiBrasil() {
	dbCreate, dbFind := SqlGormMySql()
	restApiGet := RestApiGet()
	eventApiBind, eventApiWriteMessage := EventApiWriteMessage()

	eventContextSqlGormProductTranslatedControllerCreate := EventContextSqlGormProductTranslatedControllerCreate(dbCreate)
	httpContextRestApiSqlGormEventApiBrasilProductControllerCreate := HttpContextRestApiSqlGormEventApiBrasilProductControllerCreate(restApiGet, dbCreate, eventApiBind, eventApiWriteMessage)
	httpContextSqlGormProductControllerFindAll := HttpContextSqlGormProductControllerFindAll(dbFind)
	httpContextProductControllerGet := HttpContextProductControllerGet()
	httpContextSqlGormProductTranslatedControllerFindAll := HttpContextSqlGormProductTranslatedControllerFindAll(dbFind)

	log, err := logger.NewLogger()
	if !errors.Is(err, nil) {
		panic(err)
	}

	go event_context.StartKafkaRouter("localhost:9092")(logger.Error(log))(eventContextSqlGormProductTranslatedControllerCreate)

	err = http_context.StartEchoRouter(8080)(logger.Error(log))(httpContextRestApiSqlGormEventApiBrasilProductControllerCreate)(httpContextSqlGormProductControllerFindAll)(httpContextProductControllerGet)(httpContextSqlGormProductTranslatedControllerFindAll)
	if !errors.Is(err, nil) {
		panic(err)
	}
}

func EventContextHttpContextRestApiSqlGormPostgresEventApiBrasil() {
	dbCreate, dbFind := SqlGormPostgres()
	restApiGet := RestApiGet()
	eventApiBind, eventApiWriteMessage := EventApiWriteMessage()

	eventContextSqlGormProductTranslatedControllerCreate := EventContextSqlGormProductTranslatedControllerCreate(dbCreate)
	httpContextRestApiSqlGormEventApiBrasilProductControllerCreate := HttpContextRestApiSqlGormEventApiBrasilProductControllerCreate(restApiGet, dbCreate, eventApiBind, eventApiWriteMessage)
	httpContextSqlGormProductControllerFindAll := HttpContextSqlGormProductControllerFindAll(dbFind)
	httpContextProductControllerGet := HttpContextProductControllerGet()
	httpContextSqlGormProductTranslatedControllerFindAll := HttpContextSqlGormProductTranslatedControllerFindAll(dbFind)

	log, err := logger.NewLogger()
	if !errors.Is(err, nil) {
		panic(err)
	}

	go event_context.StartKafkaRouter("localhost:9092")(logger.Error(log))(eventContextSqlGormProductTranslatedControllerCreate)

	err = http_context.StartEchoRouter(8080)(logger.Error(log))(httpContextRestApiSqlGormEventApiBrasilProductControllerCreate)(httpContextSqlGormProductControllerFindAll)(httpContextProductControllerGet)(httpContextSqlGormProductTranslatedControllerFindAll)
	if !errors.Is(err, nil) {
		panic(err)
	}
}

func EventContextHttpContextRestApiSqlGormMySqlEventApiChile() {
	dbCreate, dbFind := SqlGormMySql()
	restApiGet := RestApiGet()
	eventApiBind, eventApiWriteMessage := EventApiWriteMessage()

	eventContextSqlGormProductTranslatedControllerCreate := EventContextSqlGormProductTranslatedControllerCreate(dbCreate)
	httpContextRestApiSqlGormEventApiChileProductControllerCreate := HttpContextRestApiSqlGormEventApiChileProductControllerCreate(restApiGet, dbCreate, eventApiBind, eventApiWriteMessage)
	httpContextSqlGormProductControllerFindAll := HttpContextSqlGormProductControllerFindAll(dbFind)
	httpContextProductControllerGet := HttpContextProductControllerGet()
	httpContextSqlGormProductTranslatedControllerFindAll := HttpContextSqlGormProductTranslatedControllerFindAll(dbFind)

	log, err := logger.NewLogger()
	if !errors.Is(err, nil) {
		panic(err)
	}

	go event_context.StartKafkaRouter("localhost:9092")(logger.Error(log))(eventContextSqlGormProductTranslatedControllerCreate)

	err = http_context.StartEchoRouter(8080)(logger.Error(log))(httpContextRestApiSqlGormEventApiChileProductControllerCreate)(httpContextSqlGormProductControllerFindAll)(httpContextProductControllerGet)(httpContextSqlGormProductTranslatedControllerFindAll)
	if !errors.Is(err, nil) {
		panic(err)
	}
}

func EventContextHttpContextRestApiSqlGormPostgresEventApiChile() {
	dbCreate, dbFind := SqlGormPostgres()
	restApiGet := RestApiGet()
	eventApiBind, eventApiWriteMessage := EventApiWriteMessage()

	eventContextSqlGormProductTranslatedControllerCreate := EventContextSqlGormProductTranslatedControllerCreate(dbCreate)
	httpContextRestApiSqlGormEventApiChileProductControllerCreate := HttpContextRestApiSqlGormEventApiChileProductControllerCreate(restApiGet, dbCreate, eventApiBind, eventApiWriteMessage)
	httpContextSqlGormProductControllerFindAll := HttpContextSqlGormProductControllerFindAll(dbFind)
	httpContextProductControllerGet := HttpContextProductControllerGet()
	httpContextSqlGormProductTranslatedControllerFindAll := HttpContextSqlGormProductTranslatedControllerFindAll(dbFind)

	log, err := logger.NewLogger()
	if !errors.Is(err, nil) {
		panic(err)
	}

	go event_context.StartKafkaRouter("localhost:9092")(logger.Error(log))(eventContextSqlGormProductTranslatedControllerCreate)

	err = http_context.StartEchoRouter(8080)(logger.Error(log))(httpContextRestApiSqlGormEventApiChileProductControllerCreate)(httpContextSqlGormProductControllerFindAll)(httpContextProductControllerGet)(httpContextSqlGormProductTranslatedControllerFindAll)
	if !errors.Is(err, nil) {
		panic(err)
	}
}
