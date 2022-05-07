package registry

import (
	"go-kafka-clean-architecture/app/command/controller/event_context"

	event_context_infrastructure "go-kafka-clean-architecture/app/infrastructure/command/event_context"
	"go-kafka-clean-architecture/app/interfaces/database"
)

func (r *Registry) NewEventContextSqlHandlerMySqlProductTranslatedController(sqlHandler database.SqlHandler) event_context_infrastructure.ProductTranslatedController {
	interactor := r.NewSqlHandlerMySqlProductTranslatedInteractor(sqlHandler)

	return event_context.NewProductTranslatedController(interactor)
}

func (r *Registry) NewEventContextSqlHandlerPostgresProductTranslatedController(sqlHandler database.SqlHandler) event_context_infrastructure.ProductTranslatedController {
	interactor := r.NewSqlHandlerPostgresProductTranslatedInteractor(sqlHandler)

	return event_context.NewProductTranslatedController(interactor)
}
