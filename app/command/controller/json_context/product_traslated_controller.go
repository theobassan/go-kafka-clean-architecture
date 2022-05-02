package json_context

import (
	"net/http"

	"github.com/go-errors/errors"

	"go-kafka-clean-architecture/app/command/controller/json_context/model"
	"go-kafka-clean-architecture/app/command/router"
	"go-kafka-clean-architecture/app/command/usecases"
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
