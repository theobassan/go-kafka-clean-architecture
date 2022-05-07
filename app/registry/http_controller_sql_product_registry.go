package registry

import (
	"go-kafka-clean-architecture/app/command/controller/http_context"

	http_context_infrastructure "go-kafka-clean-architecture/app/infrastructure/command/http_context"
	"go-kafka-clean-architecture/app/interfaces/api"
	"go-kafka-clean-architecture/app/interfaces/database"
)

func (r *Registry) NewHttpContextRestApiSqlHandlerMySqlEventApiBrasilProductController(restApi api.RestApi, sqlHandler database.SqlHandler, eventApi api.EventApi) http_context_infrastructure.ProductController {
	interactor := r.NewRestApiSqlHandlerMySqlEventApiBrasilProductInteractor(restApi, sqlHandler, eventApi)

	return http_context.NewProductController(interactor)
}

func (r *Registry) NewHttpContextRestApiSqlHandlerMySqlEventApiChileProductController(restApi api.RestApi, sqlHandler database.SqlHandler, eventApi api.EventApi) http_context_infrastructure.ProductController {
	interactor := r.NewRestApiSqlHandlerMySqlEventApiChileProductInteractor(restApi, sqlHandler, eventApi)

	return http_context.NewProductController(interactor)
}

func (r *Registry) NewHttpContextSqlHandlerMySqlProductTranslatedController(sqlHandler database.SqlHandler) http_context_infrastructure.ProductTranslatedController {
	interactor := r.NewSqlHandlerMySqlProductTranslatedInteractor(sqlHandler)

	return http_context.NewProductTranslatedController(interactor)
}

func (r *Registry) NewHttpContextRestApiSqlHandlerPostgresEventApiBrasilProductController(restApi api.RestApi, sqlHandler database.SqlHandler, eventApi api.EventApi) http_context_infrastructure.ProductController {
	interactor := r.NewRestApiSqlHandlerPostgresEventApiBrasilProductInteractor(restApi, sqlHandler, eventApi)

	return http_context.NewProductController(interactor)
}

func (r *Registry) NewHttpContextRestApiSqlHandlerPostgresEventApiChileProductController(restApi api.RestApi, sqlHandler database.SqlHandler, eventApi api.EventApi) http_context_infrastructure.ProductController {
	interactor := r.NewRestApiSqlHandlerPostgresEventApiChileProductInteractor(restApi, sqlHandler, eventApi)

	return http_context.NewProductController(interactor)
}

func (r *Registry) NewHttpContextSqlHandlerPostgresProductTranslatedController(sqlHandler database.SqlHandler) http_context_infrastructure.ProductTranslatedController {
	interactor := r.NewSqlHandlerPostgresProductTranslatedInteractor(sqlHandler)

	return http_context.NewProductTranslatedController(interactor)
}
