package repository

import "go-kafka-clean-architecture/app/entities"

type ProductTranslatedRepository interface {
	Create(u *entities.Product) (*int64, error)
	FindAll() ([]*entities.Product, error)
}
