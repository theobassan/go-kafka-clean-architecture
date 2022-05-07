package model

import (
	"database/sql"
	"go-kafka-clean-architecture/app_func/entities"

	"github.com/BooleanCat/go-functional/iter"
	"github.com/didi/gendry/scanner"
	"github.com/go-errors/errors"
)

type Product struct {
	ID         int64  `ddb:"id"`
	ExternalID int64  `ddb:"external_id"`
	Type       string `ddb:"type"`
	Name       string `ddb:"name"`
}

func (Product) TableName() string { return "products" }

func MapProduct(product Product) entities.Product {
	return entities.Product{
		ID:   product.ExternalID,
		Type: product.Type,
		Name: product.Name,
	}
}

func MapProductsFromRows(rows *sql.Rows) ([]entities.Product, error) {
	modelProducts := []Product{}
	err := scanner.Scan(rows, &modelProducts)
	if !errors.Is(err, nil) {
		return nil, errors.Wrap(err, 1)
	}

	productsIteractor := iter.Lift(modelProducts)
	productsMapper := iter.Map[Product](productsIteractor, MapProduct)
	products := iter.Collect[entities.Product](productsMapper)

	return products, nil
}
