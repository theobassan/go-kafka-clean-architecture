package registry

import (
	"go-kafka-clean-architecture/app_func/command/controller/json_context"
	json_context_infrastructure "go-kafka-clean-architecture/app_func/infrastructure/command/json_context"
	"go-kafka-clean-architecture/app_func/interfaces/api"
	"go-kafka-clean-architecture/app_func/interfaces/database"
)

func JsonContextRestApiSqlHandlerMySqlEventApiBrasilProductControllerCreate(get api.RestApiGet, exec database.SqlHandlerExec, bind api.EventApiBind, writeMessage api.EventApiWriteMessage) json_context_infrastructure.ProductControllerCreate {
	interactor := RestApiSqlHandlerMySqlEventApiBrasilProductInteractorCreate(get, exec, bind, writeMessage)

	return json_context.CreateProduct()(interactor)
}

func JsonContextRestApiSqlHandlerMySqlEventApiChileProductControllerCreate(get api.RestApiGet, exec database.SqlHandlerExec, bind api.EventApiBind, writeMessage api.EventApiWriteMessage) json_context_infrastructure.ProductControllerCreate {
	interactor := RestApiSqlHandlerMySqlEventApiChileProductInteractorCreate(get, exec, bind, writeMessage)

	return json_context.CreateProduct()(interactor)
}

func JsonContextSqlHandlerMySqlProductControllerFindAll(query database.SqlHandlerQuery) json_context_infrastructure.ProductControllerFindAll {
	interactor := SqlHandlerMySqlProductInteractorFindAll(query)

	return json_context.FindAllProducts()(interactor)
}

func JsonContextSqlHandlerMySqlProductTranslatedControllerFindAll(query database.SqlHandlerQuery) json_context_infrastructure.ProductTranslatedControllerFindAll {
	interactor := SqlHandlerMySqlProductTranslatedInteractorFindAll(query)

	return json_context.FindAllProductsTranslated()(interactor)
}

func JsonContextRestApiSqlHandlerPostgresEventApiBrasilProductControllerCreate(get api.RestApiGet, queryRow database.SqlHandlerQueryRow, bind api.EventApiBind, writeMessage api.EventApiWriteMessage) json_context_infrastructure.ProductControllerCreate {
	interactor := RestApiSqlHandlerPostgresEventApiBrasilProductInteractorCreate(get, queryRow, bind, writeMessage)

	return json_context.CreateProduct()(interactor)
}

func JsonContextRestApiSqlHandlerPostgresEventApiChileProductControllerCreate(get api.RestApiGet, queryRow database.SqlHandlerQueryRow, bind api.EventApiBind, writeMessage api.EventApiWriteMessage) json_context_infrastructure.ProductControllerCreate {
	interactor := RestApiSqlHandlerPostgresEventApiChileProductInteractorCreate(get, queryRow, bind, writeMessage)

	return json_context.CreateProduct()(interactor)
}

func JsonContextSqlHandlerPostgresProductControllerFindAll(query database.SqlHandlerQuery) json_context_infrastructure.ProductControllerFindAll {
	interactor := SqlHandlerPostgresProductInteractorFindAll(query)

	return json_context.FindAllProducts()(interactor)
}

func JsonContextSqlHandlerPostgresProductTranslatedControllerFindAll(query database.SqlHandlerQuery) json_context_infrastructure.ProductTranslatedControllerFindAll {
	interactor := SqlHandlerPostgresProductTranslatedInteractorFindAll(query)

	return json_context.FindAllProductsTranslated()(interactor)
}
