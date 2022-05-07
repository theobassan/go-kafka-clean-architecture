package usecases

import (
	"go-kafka-clean-architecture/app_func/entities"
)

type ProductTranslatedInteractorCreate func(product entities.Product) (int64, error)
type ProductTranslatedInteractorFindAll func() ([]entities.Product, error)
