package registry

import (
	"errors"
	"go-kafka-clean-architecture/app/command/controller/event_context"
	"go-kafka-clean-architecture/app/command/controller/http_context"
	"go-kafka-clean-architecture/app/infrastructure/logger"
	event_context_infrastructure "go-kafka-clean-architecture/app/infrastructure/router/event_context"
	http_context_infrastructure "go-kafka-clean-architecture/app/infrastructure/router/http_context"
)

func (r *Registry) NewEventContextHttpContextRestApiSqlGormMySqlEventApiBrasil() {
	sqlGorm := r.NewSqlGormMySql()
	restApi := r.NewRestApi()
	eventApi := r.NewEventApi()

	eventContextGormDbProductTranslatedInteractor := r.NewEventContextSqlGormProductTranslatedController(sqlGorm)
	eventAppController := event_context.NewAppController(eventContextGormDbProductTranslatedInteractor)

	httpContextNewRestApiGormDbEventApiChileProductController := r.NewHttpContextRestApiSqlGormEventApiBrasilProductController(restApi, sqlGorm, eventApi)
	httpContextNewGormDbProductTranslatedController := r.NewHttpContextSqlGormProductTranslatedController(sqlGorm)
	httpAppController := http_context.NewAppController(httpContextNewRestApiGormDbEventApiChileProductController, httpContextNewGormDbProductTranslatedController)

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

func (r *Registry) NewEventContextHttpContextRestApiSqlGormPostgresEventApiBrasil() {
	sqlGorm := r.NewSqlGormPostgres()
	restApi := r.NewRestApi()
	eventApi := r.NewEventApi()

	eventContextGormDbProductTranslatedInteractor := r.NewEventContextSqlGormProductTranslatedController(sqlGorm)
	eventAppController := event_context.NewAppController(eventContextGormDbProductTranslatedInteractor)

	httpContextNewRestApiGormDbEventApiChileProductController := r.NewHttpContextRestApiSqlGormEventApiBrasilProductController(restApi, sqlGorm, eventApi)
	httpContextNewGormDbProductTranslatedController := r.NewHttpContextSqlGormProductTranslatedController(sqlGorm)
	httpAppController := http_context.NewAppController(httpContextNewRestApiGormDbEventApiChileProductController, httpContextNewGormDbProductTranslatedController)

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

func (r *Registry) NewEventContextHttpContextRestApiSqlGormMySqlEventApiChile() {
	sqlGorm := r.NewSqlGormMySql()
	restApi := r.NewRestApi()
	eventApi := r.NewEventApi()

	eventContextGormDbProductTranslatedInteractor := r.NewEventContextSqlGormProductTranslatedController(sqlGorm)
	eventAppController := event_context.NewAppController(eventContextGormDbProductTranslatedInteractor)

	httpContextNewRestApiGormDbEventApiChileProductController := r.NewHttpContextRestApiSqlGormEventApiChileProductController(restApi, sqlGorm, eventApi)
	httpContextNewGormDbProductTranslatedController := r.NewHttpContextSqlGormProductTranslatedController(sqlGorm)
	httpAppController := http_context.NewAppController(httpContextNewRestApiGormDbEventApiChileProductController, httpContextNewGormDbProductTranslatedController)

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

func (r *Registry) NewEventContextHttpContextRestApiSqlGormPostgresEventApiChile() {
	sqlGorm := r.NewSqlGormPostgres()
	restApi := r.NewRestApi()
	eventApi := r.NewEventApi()

	eventContextGormDbProductTranslatedInteractor := r.NewEventContextSqlGormProductTranslatedController(sqlGorm)
	eventAppController := event_context.NewAppController(eventContextGormDbProductTranslatedInteractor)

	httpContextNewRestApiGormDbEventApiChileProductController := r.NewHttpContextRestApiSqlGormEventApiChileProductController(restApi, sqlGorm, eventApi)
	httpContextNewGormDbProductTranslatedController := r.NewHttpContextSqlGormProductTranslatedController(sqlGorm)
	httpAppController := http_context.NewAppController(httpContextNewRestApiGormDbEventApiChileProductController, httpContextNewGormDbProductTranslatedController)

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
