package registry

import (
	"go-kafka-clean-architecture/app/command/usecases"
	"go-kafka-clean-architecture/app/interfaces/api"
	"go-kafka-clean-architecture/app/interfaces/database"
	"go-kafka-clean-architecture/app/usecases/interactor"
)

func (r *Registry) NewRestApiSqlHandlerMySqlEventApiBrasilProductInteractor(restApi api.RestApi, sqlHandler database.SqlHandler, eventApi api.EventApi) usecases.ProductInteractor {
	productFinderGateway := r.NewRestApiProductFinderGateway(restApi)
	productRepository := r.NewSqlHandlerMySqlProductRepository(sqlHandler)
	productSenderGateway := r.NewEventApiProductSenderGateway(eventApi)
	productTranslator := r.NewProductBrasilTranslator()

	return interactor.NewProductInteractor(productFinderGateway, productRepository, productSenderGateway, productTranslator)
}

func (r *Registry) NewRestApiSqlHandlerMySqlEventApiChileProductInteractor(restApi api.RestApi, sqlHandler database.SqlHandler, eventApi api.EventApi) usecases.ProductInteractor {
	productFinderGateway := r.NewRestApiProductFinderGateway(restApi)
	productRepository := r.NewSqlHandlerMySqlProductRepository(sqlHandler)
	productSenderGateway := r.NewEventApiProductSenderGateway(eventApi)
	productTranslator := r.NewProductChileTranslator()

	return interactor.NewProductInteractor(productFinderGateway, productRepository, productSenderGateway, productTranslator)
}

func (r *Registry) NewSqlHandlerMySqlProductTranslatedInteractor(sqlHandler database.SqlHandler) usecases.ProductTranslatedInteractor {
	productTranslatedRepository := r.NewSqlHandlerMySqlProductTranslatedRepository(sqlHandler)

	return interactor.NewProductTranslatedInteractor(productTranslatedRepository)
}

func (r *Registry) NewRestApiSqlHandlerPostgresEventApiBrasilProductInteractor(restApi api.RestApi, sqlHandler database.SqlHandler, eventApi api.EventApi) usecases.ProductInteractor {
	productFinderGateway := r.NewRestApiProductFinderGateway(restApi)
	productRepository := r.NewSqlHandlerPostgresProductRepository(sqlHandler)
	productSenderGateway := r.NewEventApiProductSenderGateway(eventApi)
	productTranslator := r.NewProductBrasilTranslator()

	return interactor.NewProductInteractor(productFinderGateway, productRepository, productSenderGateway, productTranslator)
}

func (r *Registry) NewRestApiSqlHandlerPostgresEventApiChileProductInteractor(restApi api.RestApi, sqlHandler database.SqlHandler, eventApi api.EventApi) usecases.ProductInteractor {
	productFinderGateway := r.NewRestApiProductFinderGateway(restApi)
	productRepository := r.NewSqlHandlerPostgresProductRepository(sqlHandler)
	productSenderGateway := r.NewEventApiProductSenderGateway(eventApi)
	productTranslator := r.NewProductChileTranslator()

	return interactor.NewProductInteractor(productFinderGateway, productRepository, productSenderGateway, productTranslator)
}

func (r *Registry) NewSqlHandlerPostgresProductTranslatedInteractor(sqlHandler database.SqlHandler) usecases.ProductTranslatedInteractor {
	productTranslatedRepository := r.NewSqlHandlerPostgresProductTranslatedRepository(sqlHandler)

	return interactor.NewProductTranslatedInteractor(productTranslatedRepository)
}
