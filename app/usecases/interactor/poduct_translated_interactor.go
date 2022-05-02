package interactor

import (
	"go-kafka-clean-architecture/app/command/usecases"
	"go-kafka-clean-architecture/app/entities"
	"go-kafka-clean-architecture/app/usecases/repository"

	"github.com/go-errors/errors"
)

type productTranslatedInteractor struct {
	productTranslatedRepository repository.ProductTranslatedRepository
}

func NewProductTranslatedInteractor(productTranslatedRepository repository.ProductTranslatedRepository) usecases.ProductTranslatedInteractor {
	return &productTranslatedInteractor{productTranslatedRepository}
}

func (interactor *productTranslatedInteractor) Create(product *entities.Product) (*int64, error) {
	createdId, err := interactor.productTranslatedRepository.Create(product)
	if !errors.Is(err, nil) {
		return nil, errors.Wrap(err, 1)
	}

	return createdId, nil
}

func (interactor *productTranslatedInteractor) FindAll() ([]*entities.Product, error) {
	products, err := interactor.productTranslatedRepository.FindAll()
	if !errors.Is(err, nil) {
		return nil, errors.Wrap(err, 1)
	}

	return products, nil
}
