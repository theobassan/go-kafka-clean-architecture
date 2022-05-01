package translator

import "go-kafka-clean-architecture/app/entities"

type ProductTranslator interface {
	Translate(product *entities.Product) *entities.Product
}
