package registry

import (
	"go-kafka-clean-architecture/app/command/controller/event_context"

	event_context_infrastructure "go-kafka-clean-architecture/app/infrastructure/command/event_context"
	"go-kafka-clean-architecture/app/interfaces/database"
)

func (r *Registry) NewEventContextSqlGormProductTranslatedController(sqlGorm database.SqlGorm) event_context_infrastructure.ProductTranslatedController {
	interactor := r.NewSqlGormProductTranslatedInteractor(sqlGorm)

	return event_context.NewProductTranslatedController(interactor)
}
