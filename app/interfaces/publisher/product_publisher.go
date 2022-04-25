package publisher

import (
	"encoding/json"
	"go-kafka-clean-architecture/app/entities"
	"go-kafka-clean-architecture/app/interfaces/broker"
	"go-kafka-clean-architecture/app/usecases/publisher"
)

type productPublisher struct {
	productWriter broker.EventWriter
}

func NewProductPublisher(productWriter broker.EventWriter) publisher.ProductPublisher {
	return &productPublisher{productWriter}
}

func (publisher *productPublisher) Publish(product *entities.Product) error {
	value, err := json.Marshal(product)
	if err != nil {
		return err
	}

	msg := publisher.productWriter.Bind(value)
	err = publisher.productWriter.WriteMessage(msg)
	if err != nil {
		panic("could not write message " + err.Error())
	}

	return nil
}
