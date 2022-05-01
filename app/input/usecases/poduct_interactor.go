package usecases

import (
	"go-kafka-clean-architecture/app/entities"
)

type ProductInteractor interface {
	Create(id *int64) (*int64, error)
	FindAll() ([]*entities.Product, error)
	Get(productID *int64) (*entities.Product, error)
}
