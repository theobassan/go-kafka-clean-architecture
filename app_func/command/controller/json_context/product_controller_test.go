package json_context

import (
	"encoding/json"
	"go-kafka-clean-architecture/app_func/command/controller/json_context/model"
	"go-kafka-clean-architecture/app_func/command/router"
	"go-kafka-clean-architecture/app_func/command/usecases"
	"go-kafka-clean-architecture/app_func/entities"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestProductControllerCreate_shoudlCreate(t *testing.T) {
	productID := int64(123)
	createdID := int64(1)
	productInteractorMock := new(usecases.ProductInteractorMock)
	productInteractorMock.On("Create", productID).Return(createdID, nil)

	product := model.Product{
		ID: productID,
	}

	productJSON, err := json.Marshal(product)
	require.NoError(t, err)

	request := httptest.NewRequest(http.MethodPost, "/product", strings.NewReader(string(productJSON)))
	responseWriter := httptest.NewRecorder()
	context := router.NewJsonContextMock(responseWriter, request)

	err = CreateProduct()(productInteractorMock.Create)(context.Bind, context.JSON)
	require.NoError(t, err)

	productInteractorMock.AssertExpectations(t)

	assert.Equal(t, http.StatusCreated, responseWriter.Code)
	assert.Equal(t, strconv.FormatInt(createdID, 10)+"\n", responseWriter.Body.String())
}

func TestProductControllerFindAll_shoudlFindAll(t *testing.T) {
	productID := int64(1)
	productType := "Type"
	productName := "Name"
	product := entities.Product{
		ID:   productID,
		Type: productType,
		Name: productName,
	}
	products := []entities.Product{product}
	productInteractorMock := new(usecases.ProductInteractorMock)
	productInteractorMock.On("FindAll").Return(products, nil)

	request := httptest.NewRequest(http.MethodGet, "/products", nil)
	responseWriter := httptest.NewRecorder()
	context := router.NewJsonContextMock(responseWriter, request)

	modelProduct := model.Product{
		ID:   productID,
		Type: productType,
		Name: productName,
	}
	modelProducts := []model.Product{modelProduct}
	productsJSON, err := json.Marshal(modelProducts)
	require.NoError(t, err)

	err = FindAllProducts()(productInteractorMock.FindAll)(context.JSON)
	require.NoError(t, err)

	productInteractorMock.AssertExpectations(t)

	assert.Equal(t, http.StatusOK, responseWriter.Code)
	assert.Equal(t, string(productsJSON)+"\n", responseWriter.Body.String())
}

func TestProductControllerGet_shoudlGet(t *testing.T) {
	productID := int64(123)
	productType := "Type"
	productName := "Name"
	product := entities.Product{
		ID:   productID,
		Type: productType,
		Name: productName,
	}

	productInteractorMock := new(usecases.ProductInteractorMock)
	productInteractorMock.On("Get", productID).Return(product, nil)

	request := httptest.NewRequest(http.MethodGet, "/product?id=123", nil)
	responseWriter := httptest.NewRecorder()
	context := router.NewJsonContextMock(responseWriter, request)

	modelProduct := model.Product{
		ID:   productID,
		Type: productType,
		Name: productName,
	}
	productJSON, err := json.Marshal(modelProduct)
	require.NoError(t, err)

	err = GetProduct()(productInteractorMock.Get)(context.Query, context.JSON)
	require.NoError(t, err)

	productInteractorMock.AssertExpectations(t)

	assert.Equal(t, http.StatusOK, responseWriter.Code)
	assert.Equal(t, string(productJSON)+"\n", responseWriter.Body.String())
}
