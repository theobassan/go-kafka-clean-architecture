package registry

import (
	"go-kafka-clean-architecture/app_func/command/usecases"
	"go-kafka-clean-architecture/app_func/interfaces/api"
	"go-kafka-clean-architecture/app_func/interfaces/database"
	"go-kafka-clean-architecture/app_func/usecases/interactor"
)

func RestApiSqlGormEventApiBrasilProductInteractorCreate(get api.RestApiGet, create database.SqlGormCreate, bind api.EventApiBind, writeMessage api.EventApiWriteMessage) usecases.ProductInteractorCreate {
	findProductById := RestApiProductFinderGatewayFindById(get)
	createProduct := SqlGormProductRepositoryCreate(create)
	sendProduct := EventApiProductSenderGatewaySend(bind, writeMessage)
	translator := TranslateProductToBrasil()

	return interactor.CreateProduct()(findProductById)(createProduct)(translator)(sendProduct)
}

func RestApiSqlGormEventApiChileProductInteractorCreate(get api.RestApiGet, create database.SqlGormCreate, bind api.EventApiBind, writeMessage api.EventApiWriteMessage) usecases.ProductInteractorCreate {
	findProductById := RestApiProductFinderGatewayFindById(get)
	createProduct := SqlGormProductRepositoryCreate(create)
	sendProduct := EventApiProductSenderGatewaySend(bind, writeMessage)
	translator := TranslateProductToChile()

	return interactor.CreateProduct()(findProductById)(createProduct)(translator)(sendProduct)
}

func SqlGormProductInteractorFindAll(find database.SqlGormFind) usecases.ProductInteractorFindAll {
	findAllProducts := SqlGormProductRepositoryFindAll(find)

	return interactor.FindAllProducts()(findAllProducts)
}

func ProductInteractorGet() usecases.ProductInteractorGet {
	return interactor.GetProduct()
}

func SqlGormProductTranslatedInteractorCreate(create database.SqlGormCreate) usecases.ProductTranslatedInteractorCreate {
	createProductTranslated := SqlGormProductTranslatedRepositoryCreate(create)
	return interactor.CreateProductTranslated()(createProductTranslated)
}

func SqlGormProductTranslatedInteractorFindAll(find database.SqlGormFind) usecases.ProductTranslatedInteractorFindAll {
	findAllProductsTranslated := SqlGormProductTranslatedRepositoryFindAll(find)

	return interactor.FindAllProductsTranslated()(findAllProductsTranslated)
}
