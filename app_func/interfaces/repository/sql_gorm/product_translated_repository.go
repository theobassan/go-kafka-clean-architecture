package sql_gorm

import (
	"github.com/BooleanCat/go-functional/iter"
	"github.com/go-errors/errors"

	"go-kafka-clean-architecture/app_func/entities"
	"go-kafka-clean-architecture/app_func/interfaces/database"
	"go-kafka-clean-architecture/app_func/interfaces/repository/sql_gorm/model"
	"go-kafka-clean-architecture/app_func/usecases/repository"
)

func FindAllProductsTranslated() func(find database.SqlGormFind) repository.ProductTranslatedRepositoryFindAll {
	return func(find database.SqlGormFind) repository.ProductTranslatedRepositoryFindAll {
		return func() ([]entities.Product, error) {
			modelProducts := []model.ProductTranslated{}
			err := find(&modelProducts).Error
			if !errors.Is(err, nil) {
				return nil, errors.Wrap(err, 1)
			}

			productsIteractor := iter.Lift(modelProducts)
			productsMapper := iter.Map[model.ProductTranslated](productsIteractor, model.MapProductTranslated)
			products := iter.Collect[entities.Product](productsMapper)

			return products, nil
		}
	}
}

func CreateProductTranslated() func(create database.SqlGormCreate) repository.ProductTranslatedRepositoryCreate {
	return func(create database.SqlGormCreate) repository.ProductTranslatedRepositoryCreate {
		return func(product entities.Product) (int64, error) {
			modelProduct := model.ProductTranslated{
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
