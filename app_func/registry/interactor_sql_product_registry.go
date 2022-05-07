package registry

import (
	"go-kafka-clean-architecture/app_func/command/usecases"
	"go-kafka-clean-architecture/app_func/interfaces/api"
	"go-kafka-clean-architecture/app_func/interfaces/database"
	"go-kafka-clean-architecture/app_func/usecases/interactor"
)

func RestApiSqlHandlerMySqlEventApiBrasilProductInteractorCreate(get api.RestApiGet, exec database.SqlHandlerExec, bind api.EventApiBind, writeMessage api.EventApiWriteMessage) usecases.ProductInteractorCreate {
	findProductById := RestApiProductFinderGatewayFindById(get)
	createProduct := SqlHandlerMySqlProductRepositoryCreate(exec)
	sendProduct := EventApiProductSenderGatewaySend(bind, writeMessage)
	translator := TranslateProductToBrasil()

	return interactor.CreateProduct()(findProductById)(createProduct)(translator)(sendProduct)
}

func RestApiSqlHandlerMySqlEventApiChileProductInteractorCreate(get api.RestApiGet, exec database.SqlHandlerExec, bind api.EventApiBind, writeMessage api.EventApiWriteMessage) usecases.ProductInteractorCreate {
	findProductById := RestApiProductFinderGatewayFindById(get)
	createProduct := SqlHandlerMySqlProductRepositoryCreate(exec)
	sendProduct := EventApiProductSenderGatewaySend(bind, writeMessage)
	translator := TranslateProductToChile()

	return interactor.CreateProduct()(findProductById)(createProduct)(translator)(sendProduct)
}

func SqlHandlerMySqlProductInteractorFindAll(query database.SqlHandlerQuery) usecases.ProductInteractorFindAll {
	findAllProducts := SqlHandlerMySqlProductRepositoryFindAll(query)

	return interactor.FindAllProducts()(findAllProducts)
}

func SqlHandlerMySqlProductTranslatedInteractorCreate(exec database.SqlHandlerExec) usecases.ProductTranslatedInteractorCreate {
	createProductTranslated := SqlHandlerMySqlProductTranslatedRepositoryCreate(exec)
	return interactor.CreateProductTranslated()(createProductTranslated)
}

func SqlHandlerMySqlProductTranslatedInteractorFindAll(query database.SqlHandlerQuery) usecases.ProductTranslatedInteractorFindAll {
	findAllProductsTranslated := SqlHandlerMySqlProductTranslatedRepositoryFindAll(query)

	return interactor.FindAllProductsTranslated()(findAllProductsTranslated)
}

func RestApiSqlHandlerPostgresEventApiBrasilProductInteractorCreate(get api.RestApiGet, queryRow database.SqlHandlerQueryRow, bind api.EventApiBind, writeMessage api.EventApiWriteMessage) usecases.ProductInteractorCreate {
	findProductById := RestApiProductFinderGatewayFindById(get)
	createProduct := SqlHandlerPostgresProductRepositoryCreate(queryRow)
	sendProduct := EventApiProductSenderGatewaySend(bind, writeMessage)
	translator := TranslateProductToBrasil()

	return interactor.CreateProduct()(findProductById)(createProduct)(translator)(sendProduct)
}

func RestApiSqlHandlerPostgresEventApiChileProductInteractorCreate(get api.RestApiGet, queryRow database.SqlHandlerQueryRow, bind api.EventApiBind, writeMessage api.EventApiWriteMessage) usecases.ProductInteractorCreate {
	findProductById := RestApiProductFinderGatewayFindById(get)
	createProduct := SqlHandlerPostgresProductRepositoryCreate(queryRow)
	sendProduct := EventApiProductSenderGatewaySend(bind, writeMessage)
	translator := TranslateProductToChile()

	return interactor.CreateProduct()(findProductById)(createProduct)(translator)(sendProduct)
}

func SqlHandlerPostgresProductInteractorFindAll(query database.SqlHandlerQuery) usecases.ProductInteractorFindAll {
	findAllProducts := SqlHandlerPostgresProductRepositoryFindAll(query)

	return interactor.FindAllProducts()(findAllProducts)
}

func SqlHandlerPostgresProductTranslatedInteractorCreate(queryRow database.SqlHandlerQueryRow) usecases.ProductTranslatedInteractorCreate {
	createProductTranslated := SqlHandlerPostgresProductTranslatedRepositoryCreate(queryRow)
	return interactor.CreateProductTranslated()(createProductTranslated)
}

func SqlHandlerPostgresProductTranslatedInteractorFindAll(query database.SqlHandlerQuery) usecases.ProductTranslatedInteractorFindAll {
	findAllProductsTranslated := SqlHandlerPostgresProductTranslatedRepositoryFindAll(query)

	return interactor.FindAllProductsTranslated()(findAllProductsTranslated)
}
