package gateway

import "go-kafka-clean-architecture/app/entities"

type ProductSenderGateway interface {
	Send(*entities.Product) error
}
