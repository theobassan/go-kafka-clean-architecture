package gateway

import "go-kafka-clean-architecture/app/entities"

type ProductGateway interface {
	FindById(id *int64) (*entities.Product, error)
}
