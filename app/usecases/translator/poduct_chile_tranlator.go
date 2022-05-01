package translator

import (
	"go-kafka-clean-architecture/app/entities"
)

type productChileTranslator struct {
}

func NewProductChileTranslator() ProductTranslator {
	return &productChileTranslator{}
}

func (interactor *productChileTranslator) Translate(product *entities.Product) *entities.Product {

	return product.TranslateToChile()
}
