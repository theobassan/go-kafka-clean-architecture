package interactor

import (
	"go-kafka-clean-architecture/app_func/command/usecases"
	"go-kafka-clean-architecture/app_func/entities"
	"go-kafka-clean-architecture/app_func/usecases/gateway"
	"go-kafka-clean-architecture/app_func/usecases/repository"
	"go-kafka-clean-architecture/app_func/usecases/translator"
	"strconv"

	"github.com/go-errors/errors"
)

func CreateProduct() func(findById gateway.ProductFinderGatewayFindById) func(create repository.ProductRepositoryCreate) func(translate translator.ProductTranslatorTranslate) func(send gateway.ProductSenderGatewaySend) usecases.ProductInteractorCreate {

	return func(findById gateway.ProductFinderGatewayFindById) func(create repository.ProductRepositoryCreate) func(translate translator.ProductTranslatorTranslate) func(send gateway.ProductSenderGatewaySend) usecases.ProductInteractorCreate {

		return func(create repository.ProductRepositoryCreate) func(translate translator.ProductTranslatorTranslate) func(send gateway.ProductSenderGatewaySend) usecases.ProductInteractorCreate {

			return func(translate translator.ProductTranslatorTranslate) func(send gateway.ProductSenderGatewaySend) usecases.ProductInteractorCreate {

				return func(send gateway.ProductSenderGatewaySend) usecases.ProductInteractorCreate {

					return func(id int64) (int64, error) {
						product, err := findById(id)
						if !errors.Is(err, nil) {
							return 0, errors.Wrap(err, 1)
						}

						createdId, err := create(product)
						if !errors.Is(err, nil) {
							return 0, errors.Wrap(err, 1)
						}

						productTranslated := translate(product)

						err = send(productTranslated)
						if !errors.Is(err, nil) {
							return 0, errors.Wrap(err, 1)
						}

						return createdId, nil
					}
				}
			}
		}
	}
}

func FindAllProducts() func(findAll repository.ProductRepositoryFindAll) usecases.ProductInteractorFindAll {
	return func(findAll repository.ProductRepositoryFindAll) usecases.ProductInteractorFindAll {
		return func() ([]entities.Product, error) {
			products, err := findAll()
			if !errors.Is(err, nil) {
				return nil, errors.Wrap(err, 1)
			}

			return products, nil
		}
	}
}

func GetProduct() usecases.ProductInteractorGet {
	return func(productID int64) (entities.Product, error) {
		productType := "Type " + strconv.FormatInt(productID, 10)
		productName := "Name " + strconv.FormatInt(productID, 10)
		product := entities.Product{
			ID:   productID,
			Type: productType,
			Name: productName,
		}

		return product, nil
	}
}
