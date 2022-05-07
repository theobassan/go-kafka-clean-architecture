package registry

import (
	"go-kafka-clean-architecture/app_func/command/controller/http_context"

	http_context_infrastructure "go-kafka-clean-architecture/app_func/infrastructure/command/http_context"
	"go-kafka-clean-architecture/app_func/interfaces/api"
	"go-kafka-clean-architecture/app_func/interfaces/database"
)

func HttpContextRestApiSqlGormEventApiBrasilProductControllerCreate(get api.RestApiGet, create database.SqlGormCreate, bind api.EventApiBind, writeMessage api.EventApiWriteMessage) http_context_infrastructure.ProductControllerCreate {
	interactor := RestApiSqlGormEventApiBrasilProductInteractorCreate(get, create, bind, writeMessage)

	return http_context.CreateProduct()(interactor)
}

func HttpContextRestApiSqlGormEventApiChileProductControllerCreate(get api.RestApiGet, create database.SqlGormCreate, bind api.EventApiBind, writeMessage api.EventApiWriteMessage) http_context_infrastructure.ProductControllerCreate {
	interactor := RestApiSqlGormEventApiChileProductInteractorCreate(get, create, bind, writeMessage)

	return http_context.CreateProduct()(interactor)
}

func HttpContextSqlGormProductControllerFindAll(find database.SqlGormFind) http_context_infrastructure.ProductControllerFindAll {
	interactor := SqlGormProductInteractorFindAll(find)

	return http_context.FindAllProducts()(interactor)
}

func HttpContextProductControllerGet() http_context_infrastructure.ProductControllerGet {
	interactor := ProductInteractorGet()

	return http_context.GetProduct()(interactor)
}

func HttpContextSqlGormProductTranslatedControllerFindAll(find database.SqlGormFind) http_context_infrastructure.ProductTranslatedControllerFindAll {
	interactor := SqlGormProductTranslatedInteractorFindAll(find)

	return http_context.FindAllProductsTranslated()(interactor)
}
