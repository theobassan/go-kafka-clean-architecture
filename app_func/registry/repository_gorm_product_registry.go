package registry

import (
	"go-kafka-clean-architecture/app_func/interfaces/database"
	"go-kafka-clean-architecture/app_func/interfaces/repository/sql_gorm"
	"go-kafka-clean-architecture/app_func/usecases/repository"
)

func SqlGormProductRepositoryCreate(create database.SqlGormCreate) repository.ProductRepositoryCreate {
	return sql_gorm.CreateProduct()(create)
}

func SqlGormProductRepositoryFindAll(find database.SqlGormFind) repository.ProductRepositoryFindAll {
	return sql_gorm.FindAllProducts()(find)
}

func SqlGormProductTranslatedRepositoryCreate(create database.SqlGormCreate) repository.ProductTranslatedRepositoryCreate {
	return sql_gorm.CreateProductTranslated()(create)
}

func SqlGormProductTranslatedRepositoryFindAll(find database.SqlGormFind) repository.ProductTranslatedRepositoryFindAll {
	return sql_gorm.FindAllProductsTranslated()(find)
}
