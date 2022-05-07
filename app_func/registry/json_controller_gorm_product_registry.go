package registry

import (
	"go-kafka-clean-architecture/app_func/command/controller/json_context"
	json_context_infrastructure "go-kafka-clean-architecture/app_func/infrastructure/command/json_context"
	"go-kafka-clean-architecture/app_func/interfaces/api"
	"go-kafka-clean-architecture/app_func/interfaces/database"
)

func JsonContextRestApiSqlGormEventApiBrasilProductControllerCreate(get api.RestApiGet, create database.SqlGormCreate, bind api.EventApiBind, writeMessage api.EventApiWriteMessage) json_context_infrastructure.ProductControllerCreate {
	interactor := RestApiSqlGormEventApiBrasilProductInteractorCreate(get, create, bind, writeMessage)

	return json_context.CreateProduct()(interactor)
}

func JsonContextRestApiSqlGormEventApiChileProductControllerCreate(get api.RestApiGet, create database.SqlGormCreate, bind api.EventApiBind, writeMessage api.EventApiWriteMessage) json_context_infrastructure.ProductControllerCreate {
	interactor := RestApiSqlGormEventApiChileProductInteractorCreate(get, create, bind, writeMessage)

	return json_context.CreateProduct()(interactor)
}

func JsonContextSqlGormProductControllerFindAll(find database.SqlGormFind) json_context_infrastructure.ProductControllerFindAll {
	interactor := SqlGormProductInteractorFindAll(find)

	return json_context.FindAllProducts()(interactor)
}

func JsonContextProductControllerGet() json_context_infrastructure.ProductControllerGet {
	interactor := ProductInteractorGet()

	return json_context.GetProduct()(interactor)
}

func JsonContextSqlGormProductTranslatedControllerFindAll(find database.SqlGormFind) json_context_infrastructure.ProductTranslatedControllerFindAll {
	interactor := SqlGormProductTranslatedInteractorFindAll(find)

	return json_context.FindAllProductsTranslated()(interactor)
}
