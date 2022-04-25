package interactor

import "go-kafka-clean-architecture/app/entities"

type ProductTranslatorInteractor interface {
	Translate(product *entities.Product) *entities.Product
}
