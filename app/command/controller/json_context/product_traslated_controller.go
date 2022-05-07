package json_context

import (
	"net/http"

	"github.com/go-errors/errors"

	"go-kafka-clean-architecture/app/command/controller/json_context/model"
	"go-kafka-clean-architecture/app/command/router"
	"go-kafka-clean-architecture/app/command/usecases"
	"go-kafka-clean-architecture/app/infrastructure/command/json_context"
)

type productTranslatedController struct {
	produtTranslatedInteractor usecases.ProductTranslatedInteractor
}

func NewProductTranslatedController(productTranslatedIteractor usecases.ProductTranslatedInteractor) json_context.ProductTranslatedController {
	return &productTranslatedController{productTranslatedIteractor}
}

func (controller *productTranslatedController) FindAll(context router.JsonContext) error {
	products, err := controller.produtTranslatedInteractor.FindAll()
	if !errors.Is(err, nil) {
		jsonErr := context.JSON(http.StatusInternalServerError, err)
		if !errors.Is(err, nil) {
			return errors.Wrap(jsonErr, 1)
		}
		return errors.Wrap(err, 1)
	}

	modelProducts := []*model.Product{}
	for _, product := range products {
		modelProduct := &model.Product{
			ID:   product.ID,
			Type: product.Type,
			Name: product.Name,
		}
		modelProducts = append(modelProducts, modelProduct)
	}

	jsonErr := context.JSON(http.StatusOK, modelProducts)
	if !errors.Is(err, nil) {
		return errors.Wrap(jsonErr, 1)
	}
	return nil
}
