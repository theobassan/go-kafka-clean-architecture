package interactor

import (
	"go-kafka-clean-architecture/app_func/command/usecases"
	"go-kafka-clean-architecture/app_func/entities"
	"go-kafka-clean-architecture/app_func/usecases/repository"

	"github.com/go-errors/errors"
)

func CreateProductTranslated() func(create repository.ProductTranslatedRepositoryCreate) usecases.ProductTranslatedInteractorCreate {
	return func(create repository.ProductTranslatedRepositoryCreate) usecases.ProductTranslatedInteractorCreate {
		return func(product entities.Product) (int64, error) {
			createdId, err := create(product)
			if !errors.Is(err, nil) {
				return 0, errors.Wrap(err, 1)
			}

			return createdId, nil
		}
	}
}

func FindAllProductsTranslated() func(findAll repository.ProductTranslatedRepositoryFindAll) usecases.ProductTranslatedInteractorFindAll {
	return func(findAll repository.ProductTranslatedRepositoryFindAll) usecases.ProductTranslatedInteractorFindAll {
		return func() ([]entities.Product, error) {
			productsTranslated, err := findAll()
			if !errors.Is(err, nil) {
				return nil, errors.Wrap(err, 1)
			}

			return productsTranslated, nil
		}
	}
}
