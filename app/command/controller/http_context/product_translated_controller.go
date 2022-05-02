package http_context

import (
	"encoding/json"
	"net/http"

	"github.com/go-errors/errors"

	"go-kafka-clean-architecture/app/command/controller/http_context/model"
	"go-kafka-clean-architecture/app/command/router"
	"go-kafka-clean-architecture/app/command/usecases"
)

type productTranslatedController struct {
	produtTranslatedInteractor usecases.ProductTranslatedInteractor
}

type ProductTranslatedController interface {
	FindAll(context router.HttpContext) error
}

func NewProductTranslatedController(productTranslatedIteractor usecases.ProductTranslatedInteractor) ProductTranslatedController {
	return &productTranslatedController{productTranslatedIteractor}
}

func (controller *productTranslatedController) FindAll(context router.HttpContext) error {

	responseWriter := context.ResponseWriter()

	products, err := controller.produtTranslatedInteractor.FindAll()
	if !errors.Is(err, nil) {
		responseWriter.Header().Set("Content-Type", "application/json")
		responseWriter.WriteHeader(http.StatusInternalServerError)

		encodeErr := json.NewEncoder(responseWriter).Encode(err)
		if !errors.Is(encodeErr, nil) {
			return errors.Wrap(encodeErr, 1)
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

	responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.WriteHeader(http.StatusOK)

	encodeErr := json.NewEncoder(responseWriter).Encode(modelProducts)
	if !errors.Is(encodeErr, nil) {
		return errors.Wrap(encodeErr, 1)
	}
	return nil
}
