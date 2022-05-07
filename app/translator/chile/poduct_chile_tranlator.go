package chile

import (
	"go-kafka-clean-architecture/app/entities"
	"go-kafka-clean-architecture/app/usecases/translator"
)

type productChileTranslator struct {
}

func NewProductChileTranslator() translator.ProductTranslator {
	return &productChileTranslator{}
}

func (translator *productChileTranslator) Translate(product *entities.Product) *entities.Product {

	return product.TranslateToChile()
}
