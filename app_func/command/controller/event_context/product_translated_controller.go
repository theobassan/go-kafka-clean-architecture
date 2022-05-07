package event_context

import (
	"github.com/go-errors/errors"

	"go-kafka-clean-architecture/app_func/command/controller/event_context/model"
	"go-kafka-clean-architecture/app_func/command/router"
	"go-kafka-clean-architecture/app_func/command/usecases"
	"go-kafka-clean-architecture/app_func/infrastructure/command/event_context"
)

func CreateProductTranslated() func(create usecases.ProductTranslatedInteractorCreate) event_context.ProductTranslatedControllerCreate {
	return func(create usecases.ProductTranslatedInteractorCreate) event_context.ProductTranslatedControllerCreate {
		return func(bind router.EventContextBind, acknowledge router.EventContextAcknowledge) error {
			var params model.Product

			err := bind(params)
			if !errors.Is(err, nil) {
				return errors.Wrap(err, 1)
			}

			product := model.MapProduct(params)

			_, err = create(product)
			if !errors.Is(err, nil) {
				return errors.Wrap(err, 1)
			}

			err = acknowledge()
			if !errors.Is(err, nil) {
				return errors.Wrap(err, 1)
			}

			return nil
		}
	}
}
