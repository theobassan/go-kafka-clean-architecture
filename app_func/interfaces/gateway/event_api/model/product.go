package model

import (
	"encoding/json"
	"go-kafka-clean-architecture/app_func/entities"

	"github.com/go-errors/errors"
)

type Product struct {
	ID   int64  `json:"id"`
	Type string `json:"type"`
	Name string `json:"name"`
}

func MapProduct(product entities.Product) ([]byte, error) {
	modelProduct := Product{
		ID:   product.ID,
		Type: product.Type,
		Name: product.Name,
	}
	value, err := json.Marshal(modelProduct)
	if !errors.Is(err, nil) {
		return nil, errors.Wrap(err, 1)
	}

	return value, nil
}
