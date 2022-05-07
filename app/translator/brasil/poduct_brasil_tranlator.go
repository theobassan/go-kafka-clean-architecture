package brasil

import (
	"go-kafka-clean-architecture/app/entities"
	"go-kafka-clean-architecture/app/usecases/translator"
)

type productBrasilTranslator struct {
}

func NewProductBrasilTranslator() translator.ProductTranslator {
	return &productBrasilTranslator{}
}

func (translator *productBrasilTranslator) Translate(product *entities.Product) *entities.Product {

	return product.TranslateToBrasil()
}
