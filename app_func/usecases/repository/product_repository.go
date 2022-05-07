package repository

import "go-kafka-clean-architecture/app_func/entities"

type ProductRepositoryCreate func(u entities.Product) (int64, error)
type ProductRepositoryFindAll func() ([]entities.Product, error)
