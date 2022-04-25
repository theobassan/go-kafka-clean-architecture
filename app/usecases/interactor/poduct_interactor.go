package interactor

import (
	"go-kafka-clean-architecture/app/entities"
	"go-kafka-clean-architecture/app/usecases/gateway"
	"go-kafka-clean-architecture/app/usecases/publisher"
	"go-kafka-clean-architecture/app/usecases/repository"
)

type productInteractor struct {
	productGateway              gateway.ProductGateway
	productRepository           repository.ProductRepository
	productPublisher            publisher.ProductPublisher
	productTranslatorInteractor ProductTranslatorInteractor
}

type ProductInteractor interface {
	Create(id *int64) error
	FindAll() ([]*entities.Product, error)
}

func NewProductInteractor(productGateway gateway.ProductGateway, productRepository repository.ProductRepository, productPublisher publisher.ProductPublisher, productTranslatorInteractor ProductTranslatorInteractor) ProductInteractor {
	return &productInteractor{productGateway, productRepository, productPublisher, productTranslatorInteractor}
}

func (interactor *productInteractor) Create(id *int64) error {
	product, err := interactor.productGateway.FindById(id)
	if err != nil {
		return err
	}

	productTranslated := interactor.productTranslatorInteractor.Translate(product)
	_, err = interactor.productRepository.Create(productTranslated)
	if err != nil {
		return err
	}

	err = interactor.productPublisher.Publish(productTranslated)
	if err != nil {
		return err
	}

	return nil
}

func (interactor *productInteractor) FindAll() ([]*entities.Product, error) {
	products, err := interactor.productRepository.FindAll()
	if err != nil {
		return nil, err
	}

	return products, nil
}
