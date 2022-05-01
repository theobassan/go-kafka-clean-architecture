package event_context

import (
	"errors"

	"go-kafka-clean-architecture/app/entities"
	"go-kafka-clean-architecture/app/input/controller/event_context/model"
	"go-kafka-clean-architecture/app/input/router"
	"go-kafka-clean-architecture/app/input/usecases"
)

type productTranslatedIController struct {
	produtTranslatedInteractor usecases.ProductTranslatedInteractor
}

type ProductTranslatedController interface {
	Create(context router.EventContext) error
}

func NewProductTranslatedController(productTranslatedIIteractor usecases.ProductTranslatedInteractor) ProductTranslatedController {
	return &productTranslatedIController{productTranslatedIIteractor}
}

func (controller *productTranslatedIController) Create(context router.EventContext) error {
	var params model.Product

	err := context.Bind(&params)
	if !errors.Is(err, nil) {
		return err
	}

	product := &entities.Product{
		ID:   params.ID,
		Type: params.Type,
		Name: params.Name,
	}

	_, err = controller.produtTranslatedInteractor.Create(product)
	if !errors.Is(err, nil) {
		return err
	}

	return context.Acknowledge()
}
