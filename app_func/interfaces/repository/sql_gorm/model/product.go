package model

import (
	"go-kafka-clean-architecture/app_func/entities"
)

type Product struct {
	ID         int64  `gorm:"primary_key:id"`
	ExternalID int64  `gorm:"column:external_id"`
	Type       string `gorm:"column:type"`
	Name       string `gorm:"column:name"`
}

func (Product) TableName() string { return "products" }

func MapProduct(product Product) entities.Product {
	return entities.Product{
		ID:   product.ExternalID,
		Type: product.Type,
		Name: product.Name,
	}
}
