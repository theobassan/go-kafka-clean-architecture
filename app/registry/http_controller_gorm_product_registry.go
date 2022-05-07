package registry

import (
	"go-kafka-clean-architecture/app/command/controller/http_context"

	http_context_infrastructure "go-kafka-clean-architecture/app/infrastructure/command/http_context"
	"go-kafka-clean-architecture/app/interfaces/api"
	"go-kafka-clean-architecture/app/interfaces/database"
)

func (r *Registry) NewHttpContextRestApiSqlGormEventApiBrasilProductController(restApi api.RestApi, sqlGorm database.SqlGorm, eventApi api.EventApi) http_context_infrastructure.ProductController {
	interactor := r.NewRestApiSqlGormEventApiBrasilProductInteractor(restApi, sqlGorm, eventApi)

	return http_context.NewProductController(interactor)
}

func (r *Registry) NewHttpContextRestApiSqlGormEventApiChileProductController(restApi api.RestApi, sqlGorm database.SqlGorm, eventApi api.EventApi) http_context_infrastructure.ProductController {
	interactor := r.NewRestApiSqlGormEventApiChileProductInteractor(restApi, sqlGorm, eventApi)

	return http_context.NewProductController(interactor)
}

func (r *Registry) NewHttpContextSqlGormProductTranslatedController(sqlGorm database.SqlGorm) http_context_infrastructure.ProductTranslatedController {
	interactor := r.NewSqlGormProductTranslatedInteractor(sqlGorm)

	return http_context.NewProductTranslatedController(interactor)
}
