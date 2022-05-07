package registry

import (
	"go-kafka-clean-architecture/app_func/command/controller/http_context"

	http_context_infrastructure "go-kafka-clean-architecture/app_func/infrastructure/command/http_context"
	"go-kafka-clean-architecture/app_func/interfaces/api"
	"go-kafka-clean-architecture/app_func/interfaces/database"
)

func HttpContextRestApiSqlHandlerMySqlEventApiBrasilProductControllerCreate(get api.RestApiGet, exec database.SqlHandlerExec, bind api.EventApiBind, writeMessage api.EventApiWriteMessage) http_context_infrastructure.ProductControllerCreate {
	interactor := RestApiSqlHandlerMySqlEventApiBrasilProductInteractorCreate(get, exec, bind, writeMessage)

	return http_context.CreateProduct()(interactor)
}

func HttpContextRestApiSqlHandlerMySqlEventApiChileProductControllerCreate(get api.RestApiGet, exec database.SqlHandlerExec, bind api.EventApiBind, writeMessage api.EventApiWriteMessage) http_context_infrastructure.ProductControllerCreate {
	interactor := RestApiSqlHandlerMySqlEventApiChileProductInteractorCreate(get, exec, bind, writeMessage)

	return http_context.CreateProduct()(interactor)
}

func HttpContextSqlHandlerMySqlProductControllerFindAll(query database.SqlHandlerQuery) http_context_infrastructure.ProductControllerFindAll {
	interactor := SqlHandlerMySqlProductInteractorFindAll(query)

	return http_context.FindAllProducts()(interactor)
}

func HttpContextSqlHandlerMySqlProductTranslatedControllerFindAll(query database.SqlHandlerQuery) http_context_infrastructure.ProductTranslatedControllerFindAll {
	interactor := SqlHandlerMySqlProductTranslatedInteractorFindAll(query)

	return http_context.FindAllProductsTranslated()(interactor)
}

func HttpContextRestApiSqlHandlerPostgresEventApiBrasilProductControllerCreate(get api.RestApiGet, queryRow database.SqlHandlerQueryRow, bind api.EventApiBind, writeMessage api.EventApiWriteMessage) http_context_infrastructure.ProductControllerCreate {
	interactor := RestApiSqlHandlerPostgresEventApiBrasilProductInteractorCreate(get, queryRow, bind, writeMessage)

	return http_context.CreateProduct()(interactor)
}

func HttpContextRestApiSqlHandlerPostgresEventApiChileProductControllerCreate(get api.RestApiGet, queryRow database.SqlHandlerQueryRow, bind api.EventApiBind, writeMessage api.EventApiWriteMessage) http_context_infrastructure.ProductControllerCreate {
	interactor := RestApiSqlHandlerPostgresEventApiChileProductInteractorCreate(get, queryRow, bind, writeMessage)

	return http_context.CreateProduct()(interactor)
}

func HttpContextSqlHandlerPostgresProductControllerFindAll(query database.SqlHandlerQuery) http_context_infrastructure.ProductControllerFindAll {
	interactor := SqlHandlerPostgresProductInteractorFindAll(query)

	return http_context.FindAllProducts()(interactor)
}

func HttpContextSqlHandlerPostgresProductTranslatedControllerFindAll(query database.SqlHandlerQuery) http_context_infrastructure.ProductTranslatedControllerFindAll {
	interactor := SqlHandlerPostgresProductTranslatedInteractorFindAll(query)

	return http_context.FindAllProductsTranslated()(interactor)
}
