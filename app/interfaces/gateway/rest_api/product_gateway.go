package rest_api

import (
	"encoding/json"
	"go-kafka-clean-architecture/app/entities"
	"go-kafka-clean-architecture/app/interfaces/api"
	"go-kafka-clean-architecture/app/interfaces/gateway/rest_api/model"
	"go-kafka-clean-architecture/app/usecases/gateway"
	"io/ioutil"
	"strconv"
)

type productGateway struct {
	restAPI api.RestAPI
}

func NewProductGateway(restAPI api.RestAPI) gateway.ProductGateway {
	return &productGateway{restAPI}
}

func (gateway *productGateway) FindById(id *int64) (*entities.Product, error) {

	response, err := gateway.restAPI.Get("http://localhost:8080/product/" + strconv.FormatInt(*id, 10))
	if err != nil {
		return nil, err
	}

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var product model.Product
	err = json.Unmarshal(responseBody, &product)
	if err != nil {
		return nil, err
	}

	return &entities.Product{
		ID:   &product.ID,
		Type: &product.Type,
		Name: &product.Name,
	}, nil
}
