package rest_context

import (
	"errors"
	"net/http"
	"strconv"

	"go-kafka-clean-architecture/app/entities"
	"go-kafka-clean-architecture/app/interfaces/controller/rest_context/model"
	"go-kafka-clean-architecture/app/interfaces/router"
	"go-kafka-clean-architecture/app/usecases/interactor"
)

type productController struct {
	produtInteractor interactor.ProductInteractor
}

type ProductController interface {
	Create(context router.RestContext) error
	FindAll(context router.RestContext) error
	Get(context router.RestContext) error
}

func NewProductController(productIteractor interactor.ProductInteractor) ProductController {
	return &productController{productIteractor}
}

func (controller *productController) Create(context router.RestContext) error {
	var params model.Product

	if err := context.Bind(&params); !errors.Is(err, nil) {
		return err
	}

	err := controller.produtInteractor.Create(params.ID)
	if !errors.Is(err, nil) {
		return err
	}

	return context.JSON(http.StatusCreated, nil)
}

func (controller *productController) FindAll(context router.RestContext) error {
	products, err := controller.produtInteractor.FindAll()
	if err != nil {
		return err
	}

	return context.JSON(http.StatusOK, products)
}

func (controller *productController) Get(context router.RestContext) error {
	id := context.Param("id")
	productID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	productType := "Type " + id
	productName := "Name " + id
	product := &entities.Product{
		ID:   &productID,
		Type: &productType,
		Name: &productName,
	}

	return context.JSON(http.StatusOK, product)
}
