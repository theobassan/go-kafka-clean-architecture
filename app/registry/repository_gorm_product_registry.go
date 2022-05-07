package registry

import (
	"go-kafka-clean-architecture/app/interfaces/database"
	"go-kafka-clean-architecture/app/interfaces/repository/sql_gorm"
	"go-kafka-clean-architecture/app/usecases/repository"
)

func (r *Registry) NewSqlGormProductRepository(sqlGorm database.SqlGorm) repository.ProductRepository {
	return sql_gorm.NewProductRepository(sqlGorm)
}

func (r *Registry) NewSqlGormProductTranslatedRepository(sqlGorm database.SqlGorm) repository.ProductTranslatedRepository {
	return sql_gorm.NewProductTranslatedRepository(sqlGorm)
}
