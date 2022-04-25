package interactor

import (
	"go-kafka-clean-architecture/app/entities"
)

type productChileTranslatorInteractor struct {
}

func NewProductChileTranslatorInteractor() ProductTranslatorInteractor {
	return &productChileTranslatorInteractor{}
}

func (interactor *productChileTranslatorInteractor) Translate(product *entities.Product) *entities.Product {

	typeChile := *product.Type + " Chile"
	nameChile := *product.Name + " Chile"
	return &entities.Product{
		ID:   product.ID,
		Type: &typeChile,
		Name: &nameChile,
	}
}
