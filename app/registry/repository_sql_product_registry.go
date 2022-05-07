package registry

import (
	"go-kafka-clean-architecture/app/interfaces/database"
	"go-kafka-clean-architecture/app/interfaces/repository/sql_handler"
	"go-kafka-clean-architecture/app/usecases/repository"
)

func (r *Registry) NewSqlHandlerMySqlProductRepository(sqlHandler database.SqlHandler) repository.ProductRepository {
	return sql_handler.NewProductRepositoryMySql(sqlHandler)
}

func (r *Registry) NewSqlHandlerMySqlProductTranslatedRepository(sqlHandler database.SqlHandler) repository.ProductTranslatedRepository {
	return sql_handler.NewProductTranslatedRepositoryMySql(sqlHandler)
}

func (r *Registry) NewSqlHandlerPostgresProductRepository(sqlHandler database.SqlHandler) repository.ProductRepository {
	return sql_handler.NewProductRepositoryPostgres(sqlHandler)
}

func (r *Registry) NewSqlHandlerPostgresProductTranslatedRepository(sqlHandler database.SqlHandler) repository.ProductTranslatedRepository {
	return sql_handler.NewProductTranslatedRepositoryPostgres(sqlHandler)
}
