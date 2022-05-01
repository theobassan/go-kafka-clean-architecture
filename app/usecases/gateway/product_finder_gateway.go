package gateway

import "go-kafka-clean-architecture/app/entities"

type ProductFinderGateway interface {
	FindById(id *int64) (*entities.Product, error)
}
