package registry

import (
	"go-kafka-clean-architecture/app/command/controller/json_context"

	json_context_infrastructure "go-kafka-clean-architecture/app/infrastructure/command/json_context"
	"go-kafka-clean-architecture/app/interfaces/api"
	"go-kafka-clean-architecture/app/interfaces/database"
)

func (r *Registry) NewJsonContextRestApiSqlHandlerMySqlEventApiBrasilProductController(restApi api.RestApi, sqlHandler database.SqlHandler, eventApi api.EventApi) json_context_infrastructure.ProductController {
	interactor := r.NewRestApiSqlHandlerMySqlEventApiBrasilProductInteractor(restApi, sqlHandler, eventApi)

	return json_context.NewProductController(interactor)
}

func (r *Registry) NewJsonContextRestApiSqlHandlerMySqlEventApiChileProductController(restApi api.RestApi, sqlHandler database.SqlHandler, eventApi api.EventApi) json_context_infrastructure.ProductController {
	interactor := r.NewRestApiSqlHandlerMySqlEventApiChileProductInteractor(restApi, sqlHandler, eventApi)

	return json_context.NewProductController(interactor)
}

func (r *Registry) NewJsonContextSqlHandlerMySqlProductTranslatedController(sqlHandler database.SqlHandler) json_context_infrastructure.ProductTranslatedController {
	interactor := r.NewSqlHandlerMySqlProductTranslatedInteractor(sqlHandler)

	return json_context.NewProductTranslatedController(interactor)
}

func (r *Registry) NewJsonContextRestApiSqlHandlerPostgresEventApiBrasilProductController(restApi api.RestApi, sqlHandler database.SqlHandler, eventApi api.EventApi) json_context_infrastructure.ProductController {
	interactor := r.NewRestApiSqlHandlerPostgresEventApiBrasilProductInteractor(restApi, sqlHandler, eventApi)

	return json_context.NewProductController(interactor)
}

func (r *Registry) NewJsonContextRestApiSqlHandlerPostgresEventApiChileProductController(restApi api.RestApi, sqlHandler database.SqlHandler, eventApi api.EventApi) json_context_infrastructure.ProductController {
	interactor := r.NewRestApiSqlHandlerPostgresEventApiChileProductInteractor(restApi, sqlHandler, eventApi)

	return json_context.NewProductController(interactor)
}

func (r *Registry) NewJsonContextSqlHandlerPostgresProductTranslatedController(sqlHandler database.SqlHandler) json_context_infrastructure.ProductTranslatedController {
	interactor := r.NewSqlHandlerPostgresProductTranslatedInteractor(sqlHandler)

	return json_context.NewProductTranslatedController(interactor)
}
