package repository

import "go-kafka-clean-architecture/app_func/entities"

type ProductTranslatedRepositoryCreate func(u entities.Product) (int64, error)
type ProductTranslatedRepositoryFindAll func() ([]entities.Product, error)
