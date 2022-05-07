package registry

import (
	"errors"
	"go-kafka-clean-architecture/app/command/controller/event_context"
	"go-kafka-clean-architecture/app/command/controller/json_context"
	"go-kafka-clean-architecture/app/infrastructure/logger"
	event_context_infrastructure "go-kafka-clean-architecture/app/infrastructure/router/event_context"
	json_context_infrastructure "go-kafka-clean-architecture/app/infrastructure/router/json_context"
)

func (r *Registry) NewEventContextJsonContextRestApiSqlGormMySqlEventApiBrasil() {
	sqlGorm := r.NewSqlGormMySql()
	restApi := r.NewRestApi()
	eventApi := r.NewEventApi()

	eventContextSqlGormProductTranslatedController := r.NewEventContextSqlGormProductTranslatedController(sqlGorm)
	eventAppController := event_context.NewAppController(eventContextSqlGormProductTranslatedController)

	jsonContextNewRestApiGormDbEventApiChileProductController := r.NewJsonContextRestApiSqlGormEventApiBrasilProductController(restApi, sqlGorm, eventApi)
	jsonContextNewGormDbProductTranslatedController := r.NewJsonContextSqlGormProductTranslatedController(sqlGorm)
	jsonAppController := json_context.NewAppController(jsonContextNewRestApiGormDbEventApiChileProductController, jsonContextNewGormDbProductTranslatedController)

	log, err := logger.NewLogger()
	if !errors.Is(err, nil) {
		panic(err)
	}

	go event_context_infrastructure.StartKafkaRouter(eventAppController, "localhost:9092", log)

	err = json_context_infrastructure.StartEchoRouter(jsonAppController, 8080, log)
	if !errors.Is(err, nil) {
		panic(err)
	}
}

func (r *Registry) NewEventContextJsonContextRestApiSqlGormPostgresEventApiBrasil() {
	sqlGorm := r.NewSqlGormPostgres()
	restApi := r.NewRestApi()
	eventApi := r.NewEventApi()

	eventContextSqlGormProductTranslatedController := r.NewEventContextSqlGormProductTranslatedController(sqlGorm)
	eventAppController := event_context.NewAppController(eventContextSqlGormProductTranslatedController)

	jsonContextNewRestApiGormDbEventApiChileProductController := r.NewJsonContextRestApiSqlGormEventApiBrasilProductController(restApi, sqlGorm, eventApi)
	jsonContextNewGormDbProductTranslatedController := r.NewJsonContextSqlGormProductTranslatedController(sqlGorm)
	jsonAppController := json_context.NewAppController(jsonContextNewRestApiGormDbEventApiChileProductController, jsonContextNewGormDbProductTranslatedController)

	log, err := logger.NewLogger()
	if !errors.Is(err, nil) {
		panic(err)
	}

	go event_context_infrastructure.StartKafkaRouter(eventAppController, "localhost:9092", log)

	err = json_context_infrastructure.StartEchoRouter(jsonAppController, 8080, log)
	if !errors.Is(err, nil) {
		panic(err)
	}
}

func (r *Registry) NewEventContextJsonContextRestApiSqlGormMySqlEventApiChile() {
	sqlGorm := r.NewSqlGormMySql()
	restApi := r.NewRestApi()
	eventApi := r.NewEventApi()

	eventContextSqlGormProductTranslatedController := r.NewEventContextSqlGormProductTranslatedController(sqlGorm)
	eventAppController := event_context.NewAppController(eventContextSqlGormProductTranslatedController)

	jsonContextNewRestApiGormDbEventApiChileProductController := r.NewJsonContextRestApiSqlGormEventApiChileProductController(restApi, sqlGorm, eventApi)
	jsonContextNewGormDbProductTranslatedController := r.NewJsonContextSqlGormProductTranslatedController(sqlGorm)
	jsonAppController := json_context.NewAppController(jsonContextNewRestApiGormDbEventApiChileProductController, jsonContextNewGormDbProductTranslatedController)

	log, err := logger.NewLogger()
	if !errors.Is(err, nil) {
		panic(err)
	}

	go event_context_infrastructure.StartKafkaRouter(eventAppController, "localhost:9092", log)

	err = json_context_infrastructure.StartEchoRouter(jsonAppController, 8080, log)
	if !errors.Is(err, nil) {
		panic(err)
	}
}

func (r *Registry) NewEventContextJsonContextRestApiSqlGormPostgresEventApiChile() {
	sqlGorm := r.NewSqlGormPostgres()
	restApi := r.NewRestApi()
	eventApi := r.NewEventApi()

	eventContextSqlGormProductTranslatedController := r.NewEventContextSqlGormProductTranslatedController(sqlGorm)
	eventAppController := event_context.NewAppController(eventContextSqlGormProductTranslatedController)

	jsonContextNewRestApiGormDbEventApiChileProductController := r.NewJsonContextRestApiSqlGormEventApiChileProductController(restApi, sqlGorm, eventApi)
	jsonContextNewGormDbProductTranslatedController := r.NewJsonContextSqlGormProductTranslatedController(sqlGorm)
	jsonAppController := json_context.NewAppController(jsonContextNewRestApiGormDbEventApiChileProductController, jsonContextNewGormDbProductTranslatedController)

	log, err := logger.NewLogger()
	if !errors.Is(err, nil) {
		panic(err)
	}

	go event_context_infrastructure.StartKafkaRouter(eventAppController, "localhost:9092", log)

	err = json_context_infrastructure.StartEchoRouter(jsonAppController, 8080, log)
	if !errors.Is(err, nil) {
		panic(err)
	}
}
