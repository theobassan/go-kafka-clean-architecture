package interactor

import (
	"go-kafka-clean-architecture/app/entities"
)

type productBrasilTranslatorInteractor struct {
}

func NewProductBrasilTranslatorInteractor() ProductTranslatorInteractor {
	return &productBrasilTranslatorInteractor{}
}

func (interactor *productBrasilTranslatorInteractor) Translate(product *entities.Product) *entities.Product {

	typeChile := *product.Type + " Brasil"
	nameChile := *product.Name + " Brasil"
	return &entities.Product{
		ID:   product.ID,
		Type: &typeChile,
		Name: &nameChile,
	}
}
