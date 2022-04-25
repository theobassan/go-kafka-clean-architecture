package registry

import (
	"go-kafka-clean-architecture/app/interfaces/broker"
	publisherImp "go-kafka-clean-architecture/app/interfaces/publisher"
	"go-kafka-clean-architecture/app/interfaces/subscriber"
	"go-kafka-clean-architecture/app/usecases/interactor"
	"go-kafka-clean-architecture/app/usecases/publisher"
)

func (r *Registry) NewBrokerProductPublisher(productWriter broker.EventWriter) publisher.ProductPublisher {
	return publisherImp.NewProductPublisher(productWriter)
}

func (r *Registry) NewBrokerProductSubscriber(productReader broker.EventReader, productIteractor interactor.ProductInteractor) subscriber.ProductSubscriber {
	return subscriber.NewProductSubscriber(productReader, productIteractor)
}
