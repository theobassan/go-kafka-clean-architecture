package registry

import (
	"go-kafka-clean-architecture/app_func/command/controller/event_context"

	event_context_infrastructure "go-kafka-clean-architecture/app_func/infrastructure/command/event_context"
	"go-kafka-clean-architecture/app_func/interfaces/database"
)

func EventContextSqlGormProductTranslatedControllerCreate(create database.SqlGormCreate) event_context_infrastructure.ProductTranslatedControllerCreate {
	interactor := SqlGormProductTranslatedInteractorCreate(create)

	return event_context.CreateProductTranslated()(interactor)
}
