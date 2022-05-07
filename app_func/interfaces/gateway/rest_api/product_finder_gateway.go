package rest_api

import (
	"strconv"

	"github.com/go-errors/errors"

	"go-kafka-clean-architecture/app_func/entities"
	"go-kafka-clean-architecture/app_func/interfaces/api"
	"go-kafka-clean-architecture/app_func/interfaces/gateway/rest_api/model"
	"go-kafka-clean-architecture/app_func/usecases/gateway"
)

func FindProductById() func(get api.RestApiGet) gateway.ProductFinderGatewayFindById {
	return func(get api.RestApiGet) gateway.ProductFinderGatewayFindById {
		return func(id int64) (entities.Product, error) {
			response, err := get("product?id=" + strconv.FormatInt(id, 10))
			if !errors.Is(err, nil) {
				return entities.Product{}, errors.Wrap(err, 1)
			}

			modelProduct, err := model.MapProductFromResponse(response)
			if !errors.Is(err, nil) {
				return entities.Product{}, errors.Wrap(err, 1)
			}

			product := model.MapProduct(model.Product(modelProduct))

			return product, nil
		}
	}
}
