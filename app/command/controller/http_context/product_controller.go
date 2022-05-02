package http_context

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-errors/errors"

	"go-kafka-clean-architecture/app/command/controller/http_context/model"
	"go-kafka-clean-architecture/app/command/router"
	"go-kafka-clean-architecture/app/command/usecases"
)

type productController struct {
	produtInteractor usecases.ProductInteractor
}

type ProductController interface {
	Create(context router.HttpContext)
	FindAll(context router.HttpContext)
	Get(context router.HttpContext)
}

func NewProductController(productIteractor usecases.ProductInteractor) ProductController {
	return &productController{productIteractor}
}

func (controller *productController) Create(context router.HttpContext) {

	responseWriter := context.ResponseWriter()
	request := context.Request()

	product := model.Product{}
	err := json.NewDecoder(request.Body).Decode(&product)
	if !errors.Is(err, nil) {

		responseWriter.Header().Set("Content-Type", "application/json")
		responseWriter.WriteHeader(http.StatusInternalServerError)

		json.NewEncoder(responseWriter).Encode(err)
	}

	id, err := controller.produtInteractor.Create(product.ID)
	if !errors.Is(err, nil) {

		responseWriter.Header().Set("Content-Type", "application/json")
		responseWriter.WriteHeader(http.StatusInternalServerError)

		json.NewEncoder(responseWriter).Encode(err)
	}

	responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.WriteHeader(http.StatusCreated)
	json.NewEncoder(responseWriter).Encode(id)
}

func (controller *productController) FindAll(context router.HttpContext) {

	responseWriter := context.ResponseWriter()

	products, err := controller.produtInteractor.FindAll()
	if !errors.Is(err, nil) {

		responseWriter.Header().Set("Content-Type", "application/json")
		responseWriter.WriteHeader(http.StatusInternalServerError)

		json.NewEncoder(responseWriter).Encode(err)
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
	json.NewEncoder(responseWriter).Encode(modelProducts)
}

func (controller *productController) Get(context router.HttpContext) {

	responseWriter := context.ResponseWriter()
	request := context.Request()

	id := request.URL.Query().Get("id")
	productID, err := strconv.ParseInt(id, 10, 64)
	if !errors.Is(err, nil) {

		responseWriter.Header().Set("Content-Type", "application/json")
		responseWriter.WriteHeader(http.StatusInternalServerError)

		json.NewEncoder(responseWriter).Encode(err)
	}
	product, err := controller.produtInteractor.Get(&productID)
	if !errors.Is(err, nil) {

		responseWriter.Header().Set("Content-Type", "application/json")
		responseWriter.WriteHeader(http.StatusInternalServerError)

		json.NewEncoder(responseWriter).Encode(err)
	}

	modelProduct := &model.Product{
		ID:   product.ID,
		Type: product.Type,
		Name: product.Name,
	}

	responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.WriteHeader(http.StatusOK)
	json.NewEncoder(responseWriter).Encode(modelProduct)
}
