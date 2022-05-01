package json_context

import (
	"errors"
	"net/http"

	"go-kafka-clean-architecture/app/input/controller/json_context/model"
	"go-kafka-clean-architecture/app/input/router"
	"go-kafka-clean-architecture/app/input/usecases"
)

type productTranslatedController struct {
	produtTranslatedInteractor usecases.ProductTranslatedInteractor
}

type ProductTranslatedController interface {
	FindAll(context router.JsonContext)
}

func NewProductTranslatedController(productTranslatedIteractor usecases.ProductTranslatedInteractor) ProductTranslatedController {
	return &productTranslatedController{productTranslatedIteractor}
}

func (controller *productTranslatedController) FindAll(context router.JsonContext) {
	products, err := controller.produtTranslatedInteractor.FindAll()
	if !errors.Is(err, nil) {
		context.JSON(http.StatusInternalServerError, err)
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

	context.JSON(http.StatusOK, modelProducts)
}
