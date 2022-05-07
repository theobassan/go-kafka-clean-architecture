package sql_gorm

import (
	"github.com/go-errors/errors"

	"go-kafka-clean-architecture/app_func/entities"
	"go-kafka-clean-architecture/app_func/interfaces/database"
	"go-kafka-clean-architecture/app_func/interfaces/repository/sql_gorm/model"
	"go-kafka-clean-architecture/app_func/usecases/repository"

	"github.com/BooleanCat/go-functional/iter"
)

func FindAllProducts() func(find database.SqlGormFind) repository.ProductRepositoryFindAll {
	return func(find database.SqlGormFind) repository.ProductRepositoryFindAll {
		return func() ([]entities.Product, error) {
			modelProducts := []model.Product{}
			err := find(&modelProducts).Error
			if !errors.Is(err, nil) {
				return nil, errors.Wrap(err, 1)
			}

			productsIteractor := iter.Lift(modelProducts)
			productsMapper := iter.Map[model.Product](productsIteractor, model.MapProduct)
			products := iter.Collect[entities.Product](productsMapper)

			return products, nil
		}
	}
}

func CreateProduct() func(create database.SqlGormCreate) repository.ProductRepositoryCreate {
	return func(create database.SqlGormCreate) repository.ProductRepositoryCreate {
		return func(product entities.Product) (int64, error) {
			modelProduct := model.Product{
				ExternalID: product.ID,
				Type:       product.Type,
				Name:       product.Name,
			}

			err := create(&modelProduct).Error
			if !errors.Is(err, nil) {
				return 0, errors.Wrap(err, 1)
			}

			return modelProduct.ID, nil
		}
	}
}
