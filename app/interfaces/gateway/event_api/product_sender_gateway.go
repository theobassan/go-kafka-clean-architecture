package event_api

import (
	"encoding/json"
	"go-kafka-clean-architecture/app/entities"
	"go-kafka-clean-architecture/app/interfaces/api"
	"go-kafka-clean-architecture/app/interfaces/gateway/event_api/model"
	"go-kafka-clean-architecture/app/usecases/gateway"

	"github.com/go-errors/errors"
)

type productGateway struct {
	eventAPI api.EventAPI
	topic    string
}

func NewProductSenderGateway(eventAPI api.EventAPI) gateway.ProductSenderGateway {
	return &productGateway{eventAPI, "product"}
}

func (publisher *productGateway) Send(product *entities.Product) error {
	modelProduct := &model.Product{
		ID:   product.ID,
		Type: product.Type,
		Name: product.Name,
	}
	value, err := json.Marshal(modelProduct)
	if !errors.Is(err, nil) {
		return errors.Wrap(err, 1)
	}

	msg := publisher.eventAPI.Bind(publisher.topic, value)
	err = publisher.eventAPI.WriteMessage(msg)
	if !errors.Is(err, nil) {
		return errors.Wrap(err, 1)
	}

	return nil
}
