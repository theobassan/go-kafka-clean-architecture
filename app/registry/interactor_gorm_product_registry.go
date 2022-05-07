package registry

import (
	"go-kafka-clean-architecture/app/command/usecases"
	"go-kafka-clean-architecture/app/interfaces/api"
	"go-kafka-clean-architecture/app/interfaces/database"
	"go-kafka-clean-architecture/app/usecases/interactor"
)

func (r *Registry) NewRestApiSqlGormEventApiBrasilProductInteractor(restApi api.RestApi, sqlGorm database.SqlGorm, eventApi api.EventApi) usecases.ProductInteractor {
	productFinderGateway := r.NewRestApiProductFinderGateway(restApi)
	productRepository := r.NewSqlGormProductRepository(sqlGorm)
	productSenderGateway := r.NewEventApiProductSenderGateway(eventApi)
	productTranslator := r.NewProductBrasilTranslator()

	return interactor.NewProductInteractor(productFinderGateway, productRepository, productSenderGateway, productTranslator)
}

func (r *Registry) NewRestApiSqlGormEventApiChileProductInteractor(restApi api.RestApi, sqlGorm database.SqlGorm, eventApi api.EventApi) usecases.ProductInteractor {
	productFinderGateway := r.NewRestApiProductFinderGateway(restApi)
	productRepository := r.NewSqlGormProductRepository(sqlGorm)
	productSenderGateway := r.NewEventApiProductSenderGateway(eventApi)
	productTranslator := r.NewProductChileTranslator()

	return interactor.NewProductInteractor(productFinderGateway, productRepository, productSenderGateway, productTranslator)
}

func (r *Registry) NewSqlGormProductTranslatedInteractor(sqlGorm database.SqlGorm) usecases.ProductTranslatedInteractor {
	productTranslatedRepository := r.NewSqlGormProductTranslatedRepository(sqlGorm)

	return interactor.NewProductTranslatedInteractor(productTranslatedRepository)
}
