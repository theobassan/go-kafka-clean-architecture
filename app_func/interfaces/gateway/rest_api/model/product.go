package model

import (
	"encoding/json"
	"go-kafka-clean-architecture/app_func/entities"
	"io/ioutil"
	"net/http"

	"github.com/go-errors/errors"
)

type Product struct {
	ID   int64  `json:"id"`
	Type string `json:"type"`
	Name string `json:"name"`
}

func MapProductFromResponse(response *http.Response) (Product, error) {
	responseBody, err := ioutil.ReadAll(response.Body)
	if !errors.Is(err, nil) {
		return Product{}, errors.Wrap(err, 1)
	}

	var product Product
	err = json.Unmarshal(responseBody, &product)
	if !errors.Is(err, nil) {
		return Product{}, errors.Wrap(err, 1)
	}

	return product, nil
}

func MapProduct(product Product) entities.Product {
	return entities.Product{
		ID:   product.ID,
		Type: product.Type,
		Name: product.Name,
	}
}
