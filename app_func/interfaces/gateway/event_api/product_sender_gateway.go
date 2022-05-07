package event_api

import (
	"go-kafka-clean-architecture/app_func/entities"
	"go-kafka-clean-architecture/app_func/interfaces/api"
	"go-kafka-clean-architecture/app_func/interfaces/gateway/event_api/model"
	"go-kafka-clean-architecture/app_func/usecases/gateway"

	"github.com/go-errors/errors"
)

const (
	topic = "product"
)

func SendProduct() func(bind api.EventApiBind) func(writeMessage api.EventApiWriteMessage) gateway.ProductSenderGatewaySend {
	return func(bind api.EventApiBind) func(writeMessage api.EventApiWriteMessage) gateway.ProductSenderGatewaySend {
		return func(writeMessage api.EventApiWriteMessage) gateway.ProductSenderGatewaySend {
			return func(product entities.Product) error {
				value, err := model.MapProduct(product)
				if !errors.Is(err, nil) {
					return errors.Wrap(err, 1)
				}

				msg := bind(topic, value)
				err = writeMessage(msg)
				if !errors.Is(err, nil) {
					return errors.Wrap(err, 1)
				}

				return nil
			}
		}
	}
}
