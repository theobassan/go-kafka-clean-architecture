package registry

import (
	"go-kafka-clean-architecture/app/interfaces/database"
	"go-kafka-clean-architecture/app/interfaces/repository/sql_gorm"
	"go-kafka-clean-architecture/app/interfaces/repository/sql_handler"
	"go-kafka-clean-architecture/app/usecases/repository"
)

func (r *Registry) NewSqlProductRepository(sqlHandler database.SQLHandler) repository.ProductRepository {
	return sql_handler.NewProductRepository(sqlHandler)
}

func (r *Registry) NewGormProductRepository(sqlGorm database.SQLGorm) repository.ProductRepository {
	return sql_gorm.NewProductRepository(sqlGorm)
}

func (r *Registry) NewSqlProductTranslatedRepository(sqlHandler database.SQLHandler) repository.ProductRepository {
	return sql_handler.NewProductTranslatedRepository(sqlHandler)
}

func (r *Registry) NewGormProductTranslatedRepository(sqlGorm database.SQLGorm) repository.ProductRepository {
	return sql_gorm.NewProductTranslatedRepository(sqlGorm)
}
