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
	Create(context router.JsonContext) error
	FindAll(context router.JsonContext) error
	Get(context router.JsonContext) error
}

func NewProductController(productIteractor usecases.ProductInteractor) ProductController {
	return &productController{productIteractor}
}

func (controller *productController) Create(context router.JsonContext) error {
	var params model.Product

	err := context.Bind(&params)
	if !errors.Is(err, nil) {
		jsonErr := context.JSON(http.StatusInternalServerError, err)
		if !errors.Is(err, nil) {
			return errors.Wrap(jsonErr, 1)
		}
		return errors.Wrap(err, 1)
	}

	id, err := controller.produtInteractor.Create(params.ID)
	if !errors.Is(err, nil) {
		jsonErr := context.JSON(http.StatusInternalServerError, err)
		if !errors.Is(err, nil) {
			return errors.Wrap(jsonErr, 1)
		}
		return errors.Wrap(err, 1)
	}

	jsonErr := context.JSON(http.StatusCreated, id)
	if !errors.Is(err, nil) {
		return errors.Wrap(jsonErr, 1)
	}
	return nil
}

func (controller *productController) FindAll(context router.JsonContext) error {
	products, err := controller.produtInteractor.FindAll()
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

func (controller *productController) Get(context router.JsonContext) error {
	id := context.Query("id")
	productID, err := strconv.ParseInt(id, 10, 64)
	if !errors.Is(err, nil) {
		jsonErr := context.JSON(http.StatusInternalServerError, err)
		if !errors.Is(err, nil) {
			return errors.Wrap(jsonErr, 1)
		}
		return errors.Wrap(err, 1)
	}
	product, err := controller.produtInteractor.Get(&productID)
	if !errors.Is(err, nil) {
		jsonErr := context.JSON(http.StatusInternalServerError, err)
		if !errors.Is(err, nil) {
			return errors.Wrap(jsonErr, 1)
		}
		return errors.Wrap(err, 1)
	}

	modelProduct := &model.Product{
		ID:   product.ID,
		Type: product.Type,
		Name: product.Name,
	}

	jsonErr := context.JSON(http.StatusOK, modelProduct)
	if !errors.Is(err, nil) {
		return errors.Wrap(jsonErr, 1)
	}
	return nil
}
