package registry

import (
	"go-kafka-clean-architecture/app_func/interfaces/database"
	"go-kafka-clean-architecture/app_func/interfaces/repository/sql_handler"
	"go-kafka-clean-architecture/app_func/usecases/repository"
)

func SqlHandlerMySqlProductRepositoryCreate(exec database.SqlHandlerExec) repository.ProductRepositoryCreate {
	return sql_handler.CreateProductMySql()(exec)
}

func SqlHandlerMySqlProductRepositoryFindAll(query database.SqlHandlerQuery) repository.ProductRepositoryFindAll {
	return sql_handler.FindAllProductsMySql()(query)
}

func SqlHandlerMySqlProductTranslatedRepositoryCreate(exec database.SqlHandlerExec) repository.ProductTranslatedRepositoryCreate {
	return sql_handler.CreateProductTranslatedMySql()(exec)
}

func SqlHandlerMySqlProductTranslatedRepositoryFindAll(query database.SqlHandlerQuery) repository.ProductTranslatedRepositoryFindAll {
	return sql_handler.FindAllProductsTranslatedMySql()(query)
}

func SqlHandlerPostgresProductRepositoryCreate(queryRow database.SqlHandlerQueryRow) repository.ProductRepositoryCreate {
	return sql_handler.CreateProductPostgres()(queryRow)
}

func SqlHandlerPostgresProductRepositoryFindAll(query database.SqlHandlerQuery) repository.ProductRepositoryFindAll {
	return sql_handler.FindAllProductsPostgres()(query)
}

func SqlHandlerPostgresProductTranslatedRepositoryCreate(queryRow database.SqlHandlerQueryRow) repository.ProductTranslatedRepositoryCreate {
	return sql_handler.CreateProductTranslatedPostgres()(queryRow)
}

func SqlHandlerPostgresProductTranslatedRepositoryFindAll(query database.SqlHandlerQuery) repository.ProductTranslatedRepositoryFindAll {
	return sql_handler.FindAllProductsTranslatedPostgres()(query)
}
