package model

import "go-kafka-clean-architecture/app_func/entities"

type Product struct {
	ID   int64  `json:"id"`
	Type string `json:"type"`
	Name string `json:"name"`
}

func MapProduct(product Product) entities.Product {
	return entities.Product{
		ID:   product.ID,
		Type: product.Type,
		Name: product.Name,
	}
}
