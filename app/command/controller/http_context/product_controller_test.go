package http_context

import (
	"encoding/json"
	"go-kafka-clean-architecture/app/command/controller/http_context/model"
	"go-kafka-clean-architecture/app/command/router"
	"go-kafka-clean-architecture/app/command/usecases"
	"go-kafka-clean-architecture/app/entities"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProductControllerCreate_shoudlCreate(t *testing.T) {
	productID := int64(123)
	createdID := int64(1)
	productInteractorMock := new(usecases.ProductInteractorMock)
	productInteractorMock.On("Create", &productID).Return(&createdID, nil)
	productController := NewProductController(productInteractorMock)

	product := &model.Product{
		ID: &productID,
	}

	productJSON, err := json.Marshal(product)
	assert.NoError(t, err)

	request := httptest.NewRequest(http.MethodPost, "/product", strings.NewReader(string(productJSON)))
	responseWriter := httptest.NewRecorder()
	context := router.NewHttpContextMock(responseWriter, request)

	productController.Create(context)

	productInteractorMock.AssertExpectations(t)

	assert.Equal(t, http.StatusCreated, responseWriter.Code)
	assert.Equal(t, strconv.FormatInt(createdID, 10)+"\n", responseWriter.Body.String())
}

func TestProductControllerFindAll_shoudlFindAll(t *testing.T) {
	productID := int64(1)
	productType := "Type"
	productName := "Name"
	product := &entities.Product{
		ID:   &productID,
		Type: &productType,
		Name: &productName,
	}
	products := []*entities.Product{product}
	productInteractorMock := new(usecases.ProductInteractorMock)
	productInteractorMock.On("FindAll").Return(products, nil)
	productController := NewProductController(productInteractorMock)

	request := httptest.NewRequest(http.MethodGet, "/products", nil)
	responseWriter := httptest.NewRecorder()
	context := router.NewHttpContextMock(responseWriter, request)

	modelProduct := &model.Product{
		ID:   &productID,
		Type: &productType,
		Name: &productName,
	}
	modelProducts := []*model.Product{modelProduct}
	productsJSON, err := json.Marshal(modelProducts)
	assert.NoError(t, err)

	productController.FindAll(context)

	productInteractorMock.AssertExpectations(t)

	assert.Equal(t, http.StatusOK, responseWriter.Code)
	assert.Equal(t, string(productsJSON)+"\n", responseWriter.Body.String())
}

func TestProductControllerGet_shoudlGet(t *testing.T) {
	productID := int64(123)
	productType := "Type"
	productName := "Name"
	product := &entities.Product{
		ID:   &productID,
		Type: &productType,
		Name: &productName,
	}

	productInteractorMock := new(usecases.ProductInteractorMock)
	productInteractorMock.On("Get", &productID).Return(product, nil)
	productController := NewProductController(productInteractorMock)

	request := httptest.NewRequest(http.MethodGet, "/product?id=123", nil)
	responseWriter := httptest.NewRecorder()
	context := router.NewHttpContextMock(responseWriter, request)

	modelProduct := &model.Product{
		ID:   &productID,
		Type: &productType,
		Name: &productName,
	}
	productJSON, err := json.Marshal(modelProduct)
	assert.NoError(t, err)

	productController.Get(context)

	productInteractorMock.AssertExpectations(t)

	assert.Equal(t, http.StatusOK, responseWriter.Code)
	assert.Equal(t, string(productJSON)+"\n", responseWriter.Body.String())
}
