package registry

import (
	"errors"
	"go-kafka-clean-architecture/app/command/controller/event_context"
	"go-kafka-clean-architecture/app/command/controller/json_context"
	"go-kafka-clean-architecture/app/infrastructure/logger"
	event_context_infrastructure "go-kafka-clean-architecture/app/infrastructure/router/event_context"
	json_context_infrastructure "go-kafka-clean-architecture/app/infrastructure/router/json_context"
)

func (r *Registry) NewEventContextJsonContextRestApiSqlHandlerMySqlEventApiBrasil() {
	sqlHandler := r.NewSqlHandlerMySql()
	restApi := r.NewRestApi()
	eventApi := r.NewEventApi()

	eventContextSqlHandlerMySqlProductTranslatedController := r.NewEventContextSqlHandlerMySqlProductTranslatedController(sqlHandler)
	eventAppController := event_context.NewAppController(eventContextSqlHandlerMySqlProductTranslatedController)

	jsonContextNewRestApiSqlDbEventApiChileProductController := r.NewJsonContextRestApiSqlHandlerMySqlEventApiBrasilProductController(restApi, sqlHandler, eventApi)
	jsonContextNewSqlDbProductTranslatedController := r.NewJsonContextSqlHandlerMySqlProductTranslatedController(sqlHandler)
	jsonAppController := json_context.NewAppController(jsonContextNewRestApiSqlDbEventApiChileProductController, jsonContextNewSqlDbProductTranslatedController)

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

func (r *Registry) NewEventContextJsonContextRestApiSqlHandlerPostgresEventApiBrasil() {
	sqlHandler := r.NewSqlHandlerPostgres()
	restApi := r.NewRestApi()
	eventApi := r.NewEventApi()

	eventContextSqlHandlerMySqlProductTranslatedController := r.NewEventContextSqlHandlerMySqlProductTranslatedController(sqlHandler)
	eventAppController := event_context.NewAppController(eventContextSqlHandlerMySqlProductTranslatedController)

	jsonContextNewRestApiSqlDbEventApiChileProductController := r.NewJsonContextRestApiSqlHandlerMySqlEventApiBrasilProductController(restApi, sqlHandler, eventApi)
	jsonContextNewSqlDbProductTranslatedController := r.NewJsonContextSqlHandlerMySqlProductTranslatedController(sqlHandler)
	jsonAppController := json_context.NewAppController(jsonContextNewRestApiSqlDbEventApiChileProductController, jsonContextNewSqlDbProductTranslatedController)

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

func (r *Registry) NewEventContextJsonContextRestApiSqlHandlerMySqlEventApiChile() {
	sqlHandler := r.NewSqlHandlerMySql()
	restApi := r.NewRestApi()
	eventApi := r.NewEventApi()

	eventContextSqlHandlerMySqlProductTranslatedController := r.NewEventContextSqlHandlerMySqlProductTranslatedController(sqlHandler)
	eventAppController := event_context.NewAppController(eventContextSqlHandlerMySqlProductTranslatedController)

	jsonContextNewRestApiSqlDbEventApiChileProductController := r.NewJsonContextRestApiSqlHandlerMySqlEventApiChileProductController(restApi, sqlHandler, eventApi)
	jsonContextNewSqlDbProductTranslatedController := r.NewJsonContextSqlHandlerMySqlProductTranslatedController(sqlHandler)
	jsonAppController := json_context.NewAppController(jsonContextNewRestApiSqlDbEventApiChileProductController, jsonContextNewSqlDbProductTranslatedController)

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

func (r *Registry) NewEventContextJsonContextRestApiSqlHandlerPostgresEventApiChile() {
	sqlHandler := r.NewSqlHandlerPostgres()
	restApi := r.NewRestApi()
	eventApi := r.NewEventApi()

	eventContextSqlHandlerMySqlProductTranslatedController := r.NewEventContextSqlHandlerMySqlProductTranslatedController(sqlHandler)
	eventAppController := event_context.NewAppController(eventContextSqlHandlerMySqlProductTranslatedController)

	jsonContextNewRestApiSqlDbEventApiChileProductController := r.NewJsonContextRestApiSqlHandlerMySqlEventApiChileProductController(restApi, sqlHandler, eventApi)
	jsonContextNewSqlDbProductTranslatedController := r.NewJsonContextSqlHandlerMySqlProductTranslatedController(sqlHandler)
	jsonAppController := json_context.NewAppController(jsonContextNewRestApiSqlDbEventApiChileProductController, jsonContextNewSqlDbProductTranslatedController)

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
