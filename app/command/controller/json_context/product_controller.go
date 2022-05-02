package json_context

import (
	"net/http"
	"strconv"

	"github.com/go-errors/errors"

	"go-kafka-clean-architecture/app/command/controller/json_context/model"
	"go-kafka-clean-architecture/app/command/router"
	"go-kafka-clean-architecture/app/command/usecases"
)

type productController struct {
	produtInteractor usecases.ProductInteractor
}

type ProductController interface {
	Create(context router.JsonContext)
	FindAll(context router.JsonContext)
	Get(context router.JsonContext)
}

func NewProductController(productIteractor usecases.ProductInteractor) ProductController {
	return &productController{productIteractor}
}

func (controller *productController) Create(context router.JsonContext) {
	var params model.Product

	err := context.Bind(&params)
	if !errors.Is(err, nil) {
		context.JSON(http.StatusInternalServerError, err)
	}

	id, err := controller.produtInteractor.Create(params.ID)
	if !errors.Is(err, nil) {
		context.JSON(http.StatusCreated, err)
	}

	context.JSON(http.StatusCreated, id)
}

func (controller *productController) FindAll(context router.JsonContext) {
	products, err := controller.produtInteractor.FindAll()
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

func (controller *productController) Get(context router.JsonContext) {
	id := context.Query("id")
	productID, err := strconv.ParseInt(id, 10, 64)
	if !errors.Is(err, nil) {
		context.JSON(http.StatusInternalServerError, err)
	}
	product, err := controller.produtInteractor.Get(&productID)
	if !errors.Is(err, nil) {
		context.JSON(http.StatusInternalServerError, err)
	}

	modelProduct := &model.Product{
		ID:   product.ID,
		Type: product.Type,
		Name: product.Name,
	}

	context.JSON(http.StatusOK, modelProduct)
}
