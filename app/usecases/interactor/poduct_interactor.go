package interactor

import (
	"go-kafka-clean-architecture/app/command/usecases"
	"go-kafka-clean-architecture/app/entities"
	"go-kafka-clean-architecture/app/usecases/gateway"
	"go-kafka-clean-architecture/app/usecases/repository"
	"go-kafka-clean-architecture/app/usecases/translator"
	"strconv"

	"github.com/go-errors/errors"
)

type productInteractor struct {
	productFinderGateway gateway.ProductFinderGateway
	productRepository    repository.ProductRepository
	productSenderGateway gateway.ProductSenderGateway
	productTranslator    translator.ProductTranslator
}

func NewProductInteractor(productFinderGateway gateway.ProductFinderGateway, productRepository repository.ProductRepository, productSenderGateway gateway.ProductSenderGateway, productTranslator translator.ProductTranslator) usecases.ProductInteractor {
	return &productInteractor{productFinderGateway, productRepository, productSenderGateway, productTranslator}
}

func (interactor *productInteractor) Create(id *int64) (*int64, error) {
	product, err := interactor.productFinderGateway.FindById(id)
	if !errors.Is(err, nil) {
		return nil, errors.Wrap(err, 1)
	}

	createdId, err := interactor.productRepository.Create(product)
	if !errors.Is(err, nil) {
		return nil, errors.Wrap(err, 1)
	}

	productTranslated := interactor.productTranslator.Translate(product)
	err = interactor.productSenderGateway.Send(productTranslated)
	if !errors.Is(err, nil) {
		return nil, errors.Wrap(err, 1)
	}

	return createdId, nil
}

func (interactor *productInteractor) FindAll() ([]*entities.Product, error) {
	products, err := interactor.productRepository.FindAll()
	if !errors.Is(err, nil) {
		return nil, errors.Wrap(err, 1)
	}

	return products, nil
}

func (interactor *productInteractor) Get(productID *int64) (*entities.Product, error) {
	productType := "Type " + strconv.FormatInt(*productID, 10)
	productName := "Name " + strconv.FormatInt(*productID, 10)
	product := &entities.Product{
		ID:   productID,
		Type: &productType,
		Name: &productName,
	}

	return product, nil
}
