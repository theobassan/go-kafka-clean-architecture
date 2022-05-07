package registry

import (
	"go-kafka-clean-architecture/app_func/command/controller/event_context"

	event_context_infrastructure "go-kafka-clean-architecture/app_func/infrastructure/command/event_context"
	"go-kafka-clean-architecture/app_func/interfaces/database"
)

func EventContextSqlHandlerMySqlProductTranslatedControllerCreate(exec database.SqlHandlerExec) event_context_infrastructure.ProductTranslatedControllerCreate {
	interactor := SqlHandlerMySqlProductTranslatedInteractorCreate(exec)

	return event_context.CreateProductTranslated()(interactor)
}

func EventContextSqlHandlerPostgresProductTranslatedControllerCreate(queryRow database.SqlHandlerQueryRow) event_context_infrastructure.ProductTranslatedControllerCreate {
	interactor := SqlHandlerPostgresProductTranslatedInteractorCreate(queryRow)

	return event_context.CreateProductTranslated()(interactor)
}
