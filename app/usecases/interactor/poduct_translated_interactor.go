package interactor

import (
	"errors"
	"go-kafka-clean-architecture/app/entities"
	"go-kafka-clean-architecture/app/input/usecases"
	"go-kafka-clean-architecture/app/usecases/repository"
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
		return nil, err
	}

	return createdId, nil
}

func (interactor *productTranslatedInteractor) FindAll() ([]*entities.Product, error) {
	products, err := interactor.productTranslatedRepository.FindAll()
	if !errors.Is(err, nil) {
		return nil, err
	}

	return products, nil
}
