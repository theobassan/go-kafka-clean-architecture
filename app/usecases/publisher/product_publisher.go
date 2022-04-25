package publisher

import "go-kafka-clean-architecture/app/entities"

type ProductPublisher interface {
	Publish(*entities.Product) error
}
