package usecases

import (
	"go-kafka-clean-architecture/app_func/entities"
)

type ProductInteractorCreate func(id int64) (int64, error)
type ProductInteractorFindAll func() ([]entities.Product, error)
type ProductInteractorGet func(productID int64) (entities.Product, error)
