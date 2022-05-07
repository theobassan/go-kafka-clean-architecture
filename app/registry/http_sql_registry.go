package registry

import (
	"errors"
	"go-kafka-clean-architecture/app/command/controller/event_context"
	"go-kafka-clean-architecture/app/command/controller/http_context"
	"go-kafka-clean-architecture/app/infrastructure/logger"
	event_context_infrastructure "go-kafka-clean-architecture/app/infrastructure/router/event_context"
	http_context_infrastructure "go-kafka-clean-architecture/app/infrastructure/router/http_context"
)

func (r *Registry) NewEventContextHttpContextRestApiSqlHandlerMySqlEventApiBrasil() {
	sqlHandler := r.NewSqlHandlerMySql()
	restApi := r.NewRestApi()
	eventApi := r.NewEventApi()

	eventContextSqlDbProductTranslatedInteractor := r.NewEventContextSqlHandlerMySqlProductTranslatedController(sqlHandler)
	eventAppController := event_context.NewAppController(eventContextSqlDbProductTranslatedInteractor)

	httpContextNewRestApiSqlDbEventApiChileProductController := r.NewHttpContextRestApiSqlHandlerMySqlEventApiBrasilProductController(restApi, sqlHandler, eventApi)
	httpContextNewSqlDbProductTranslatedController := r.NewHttpContextSqlHandlerMySqlProductTranslatedController(sqlHandler)
	httpAppController := http_context.NewAppController(httpContextNewRestApiSqlDbEventApiChileProductController, httpContextNewSqlDbProductTranslatedController)

	log, err := logger.NewLogger()
	if !errors.Is(err, nil) {
		panic(err)
	}

	go event_context_infrastructure.StartKafkaRouter(eventAppController, "localhost:9092", log)

	err = http_context_infrastructure.StartEchoRouter(httpAppController, 8080, log)
	if !errors.Is(err, nil) {
		panic(err)
	}
}

func (r *Registry) NewEventContextHttpContextRestApiSqlHandlerPostgresEventApiBrasil() {
	sqlHandler := r.NewSqlHandlerPostgres()
	restApi := r.NewRestApi()
	eventApi := r.NewEventApi()

	eventContextSqlDbProductTranslatedInteractor := r.NewEventContextSqlHandlerMySqlProductTranslatedController(sqlHandler)
	eventAppController := event_context.NewAppController(eventContextSqlDbProductTranslatedInteractor)

	httpContextNewRestApiSqlDbEventApiChileProductController := r.NewHttpContextRestApiSqlHandlerMySqlEventApiBrasilProductController(restApi, sqlHandler, eventApi)
	httpContextNewSqlDbProductTranslatedController := r.NewHttpContextSqlHandlerMySqlProductTranslatedController(sqlHandler)
	httpAppController := http_context.NewAppController(httpContextNewRestApiSqlDbEventApiChileProductController, httpContextNewSqlDbProductTranslatedController)

	log, err := logger.NewLogger()
	if !errors.Is(err, nil) {
		panic(err)
	}

	go event_context_infrastructure.StartKafkaRouter(eventAppController, "localhost:9092", log)

	err = http_context_infrastructure.StartEchoRouter(httpAppController, 8080, log)
	if !errors.Is(err, nil) {
		panic(err)
	}
}

func (r *Registry) NewEventContextHttpContextRestApiSqlHandlerMySqlEventApiChile() {
	sqlHandler := r.NewSqlHandlerMySql()
	restApi := r.NewRestApi()
	eventApi := r.NewEventApi()

	eventContextSqlDbProductTranslatedInteractor := r.NewEventContextSqlHandlerMySqlProductTranslatedController(sqlHandler)
	eventAppController := event_context.NewAppController(eventContextSqlDbProductTranslatedInteractor)

	httpContextNewRestApiSqlDbEventApiChileProductController := r.NewHttpContextRestApiSqlHandlerMySqlEventApiChileProductController(restApi, sqlHandler, eventApi)
	httpContextNewSqlDbProductTranslatedController := r.NewHttpContextSqlHandlerMySqlProductTranslatedController(sqlHandler)
	httpAppController := http_context.NewAppController(httpContextNewRestApiSqlDbEventApiChileProductController, httpContextNewSqlDbProductTranslatedController)

	log, err := logger.NewLogger()
	if !errors.Is(err, nil) {
		panic(err)
	}

	go event_context_infrastructure.StartKafkaRouter(eventAppController, "localhost:9092", log)

	err = http_context_infrastructure.StartEchoRouter(httpAppController, 8080, log)
	if !errors.Is(err, nil) {
		panic(err)
	}
}

func (r *Registry) NewEventContextHttpContextRestApiSqlHandlerPostgresEventApiChile() {
	sqlHandler := r.NewSqlHandlerPostgres()
	restApi := r.NewRestApi()
	eventApi := r.NewEventApi()

	eventContextSqlDbProductTranslatedInteractor := r.NewEventContextSqlHandlerMySqlProductTranslatedController(sqlHandler)
	eventAppController := event_context.NewAppController(eventContextSqlDbProductTranslatedInteractor)

	httpContextNewRestApiSqlDbEventApiChileProductController := r.NewHttpContextRestApiSqlHandlerMySqlEventApiChileProductController(restApi, sqlHandler, eventApi)
	httpContextNewSqlDbProductTranslatedController := r.NewHttpContextSqlHandlerMySqlProductTranslatedController(sqlHandler)
	httpAppController := http_context.NewAppController(httpContextNewRestApiSqlDbEventApiChileProductController, httpContextNewSqlDbProductTranslatedController)

	log, err := logger.NewLogger()
	if !errors.Is(err, nil) {
		panic(err)
	}

	go event_context_infrastructure.StartKafkaRouter(eventAppController, "localhost:9092", log)

	err = http_context_infrastructure.StartEchoRouter(httpAppController, 8080, log)
	if !errors.Is(err, nil) {
		panic(err)
	}
}
