package model

import (
	"go-kafka-clean-architecture/app_func/entities"
)

type ProductTranslated struct {
	ID         int64  `gorm:"primary_key:id"`
	ExternalID int64  `gorm:"column:external_id"`
	Type       string `gorm:"column:type"`
	Name       string `gorm:"column:name"`
}

func (ProductTranslated) TableName() string { return "products_translated" }

func MapProductTranslated(productTranslated ProductTranslated) entities.Product {
	return entities.Product{
		ID:   productTranslated.ExternalID,
		Type: productTranslated.Type,
		Name: productTranslated.Name,
	}
}
