package translator

import (
	"go-kafka-clean-architecture/app/entities"
)

type productBrasilTranslator struct {
}

func NewProductBrasilTranslator() ProductTranslator {
	return &productBrasilTranslator{}
}

func (interactor *productBrasilTranslator) Translate(product *entities.Product) *entities.Product {

	return product.TranslateToBrasil()
}
