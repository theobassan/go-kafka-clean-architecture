package registry

import (
	"go-kafka-clean-architecture/app/interfaces/database"
	"go-kafka-clean-architecture/app/interfaces/repository/sql_gorm"
	"go-kafka-clean-architecture/app/interfaces/repository/sql_handler"
	"go-kafka-clean-architecture/app/usecases/repository"
)

func (r *Registry) NewSqlProductRepositoryMySql(sqlHandler database.SQLHandler) repository.ProductRepository {
	return sql_handler.NewProductRepositoryMySql(sqlHandler)
}

func (r *Registry) NewSqlProductTranslatedRepositoryMySql(sqlHandler database.SQLHandler) repository.ProductRepository {
	return sql_handler.NewProductTranslatedRepositoryMySql(sqlHandler)
}

func (r *Registry) NewSqlProductRepositoryPostgres(sqlHandler database.SQLHandler) repository.ProductRepository {
	return sql_handler.NewProductRepositoryPostgres(sqlHandler)
}

func (r *Registry) NewSqlProductTranslatedRepositoryPostgres(sqlHandler database.SQLHandler) repository.ProductRepository {
	return sql_handler.NewProductTranslatedRepositoryPostgres(sqlHandler)
}

func (r *Registry) NewGormProductRepository(sqlGorm database.SQLGorm) repository.ProductRepository {
	return sql_gorm.NewProductRepository(sqlGorm)
}

func (r *Registry) NewGormProductTranslatedRepository(sqlGorm database.SQLGorm) repository.ProductRepository {
	return sql_gorm.NewProductTranslatedRepository(sqlGorm)
}
