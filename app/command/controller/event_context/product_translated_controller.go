package event_context

import (
	"github.com/go-errors/errors"

	"go-kafka-clean-architecture/app/command/controller/event_context/model"
	"go-kafka-clean-architecture/app/command/router"
	"go-kafka-clean-architecture/app/command/usecases"
	"go-kafka-clean-architecture/app/entities"
	"go-kafka-clean-architecture/app/infrastructure/command/event_context"
)

type productTranslatedIController struct {
	produtTranslatedInteractor usecases.ProductTranslatedInteractor
}

func NewProductTranslatedController(productTranslatedIIteractor usecases.ProductTranslatedInteractor) event_context.ProductTranslatedController {
	return &productTranslatedIController{productTranslatedIIteractor}
}

func (controller *productTranslatedIController) Create(context router.EventContext) error {
	var params model.Product

	err := context.Bind(&params)
	if !errors.Is(err, nil) {
		return errors.Wrap(err, 1)
	}

	product := &entities.Product{
		ID:   params.ID,
		Type: params.Type,
		Name: params.Name,
	}

	_, err = controller.produtTranslatedInteractor.Create(product)
	if !errors.Is(err, nil) {
		return errors.Wrap(err, 1)
	}

	err = context.Acknowledge()
	if !errors.Is(err, nil) {
		return errors.Wrap(err, 1)
	}

	return nil
}
