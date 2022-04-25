package subscriber

import (
	"context"
	"encoding/json"
	"fmt"
	"go-kafka-clean-architecture/app/interfaces/broker"
	"go-kafka-clean-architecture/app/interfaces/subscriber/model"
	"go-kafka-clean-architecture/app/usecases/interactor"
	"strconv"
)

type productSubscriber struct {
	productReader    broker.EventReader
	produtInteractor interactor.ProductInteractor
}

type ProductSubscriber interface {
	Subscribe(ctx context.Context) error
}

func NewProductSubscriber(productReader broker.EventReader, productIteractor interactor.ProductInteractor) ProductSubscriber {
	return &productSubscriber{productReader, productIteractor}
}

func (subscriber *productSubscriber) Subscribe(ctx context.Context) error {

	for {
		msg, err := subscriber.productReader.ReadMessage()
		if err != nil {
			fmt.Println(err.Error())
			return err
		}

		value := subscriber.productReader.Bind(msg)

		var product model.Product
		err = json.Unmarshal(value, &product)
		if err != nil {
			return err
		}

		//err = subscriber.produtInteractor.Create(id)
		_, err = fmt.Println("ReadMessage -> ID: " + strconv.FormatInt(*product.ID, 10))
		if err != nil {
			return err
		}
	}
}
