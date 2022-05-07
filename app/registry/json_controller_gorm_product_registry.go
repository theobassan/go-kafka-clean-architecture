package registry

import (
	"go-kafka-clean-architecture/app/command/controller/json_context"

	json_context_infrastructure "go-kafka-clean-architecture/app/infrastructure/command/json_context"
	"go-kafka-clean-architecture/app/interfaces/api"
	"go-kafka-clean-architecture/app/interfaces/database"
)

func (r *Registry) NewJsonContextRestApiSqlGormEventApiBrasilProductController(restApi api.RestApi, sqlGorm database.SqlGorm, eventApi api.EventApi) json_context_infrastructure.ProductController {
	interactor := r.NewRestApiSqlGormEventApiBrasilProductInteractor(restApi, sqlGorm, eventApi)

	return json_context.NewProductController(interactor)
}

func (r *Registry) NewJsonContextRestApiSqlGormEventApiChileProductController(restApi api.RestApi, sqlGorm database.SqlGorm, eventApi api.EventApi) json_context_infrastructure.ProductController {
	interactor := r.NewRestApiSqlGormEventApiChileProductInteractor(restApi, sqlGorm, eventApi)

	return json_context.NewProductController(interactor)
}

func (r *Registry) NewJsonContextSqlGormProductTranslatedController(sqlGorm database.SqlGorm) json_context_infrastructure.ProductTranslatedController {
	interactor := r.NewSqlGormProductTranslatedInteractor(sqlGorm)

	return json_context.NewProductTranslatedController(interactor)
}
