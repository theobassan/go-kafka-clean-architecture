package model

import (
	"database/sql"
	"go-kafka-clean-architecture/app_func/entities"

	"github.com/BooleanCat/go-functional/iter"
	"github.com/didi/gendry/scanner"
	"github.com/go-errors/errors"
)

type ProductTranslated struct {
	ID         int64  `ddb:"id"`
	ExternalID int64  `ddb:"external_id"`
	Type       string `ddb:"type"`
	Name       string `ddb:"name"`
}

func MapProductTranslated(productTranslated ProductTranslated) entities.Product {
	return entities.Product{
		ID:   productTranslated.ExternalID,
		Type: productTranslated.Type,
		Name: productTranslated.Name,
	}
}

func MapProductsTranslatedFromRows(rows *sql.Rows) ([]entities.Product, error) {
	modelProducts := []ProductTranslated{}
	err := scanner.Scan(rows, &modelProducts)
	if !errors.Is(err, nil) {
		return nil, errors.Wrap(err, 1)
	}

	productsIteractor := iter.Lift(modelProducts)
	productsMapper := iter.Map[ProductTranslated](productsIteractor, MapProductTranslated)
	products := iter.Collect[entities.Product](productsMapper)

	return products, nil
}
