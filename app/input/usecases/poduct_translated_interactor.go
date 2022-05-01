package usecases

import (
	"go-kafka-clean-architecture/app/entities"
)

type ProductTranslatedInteractor interface {
	Create(product *entities.Product) (*int64, error)
	FindAll() ([]*entities.Product, error)
}
