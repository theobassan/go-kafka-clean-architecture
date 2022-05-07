package http_context

import (
	"encoding/json"
	"go-kafka-clean-architecture/app_func/command/controller/http_context/model"
	"go-kafka-clean-architecture/app_func/command/router"
	"go-kafka-clean-architecture/app_func/command/usecases"
	"go-kafka-clean-architecture/app_func/entities"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestProductTranslatedControllerFindAll_shoudlFindAll(t *testing.T) {
	productID := int64(1)
	productType := "Type"
	productName := "Name"
	product := entities.Product{
		ID:   productID,
		Type: productType,
		Name: productName,
	}
	products := []entities.Product{product}
	productTranslatedInteractorMock := new(usecases.ProductTranslatedInteractorMock)
	productTranslatedInteractorMock.On("FindAll").Return(products, nil)

	request := httptest.NewRequest(http.MethodGet, "/productstranslated", nil)
	responseWriter := httptest.NewRecorder()
	context := router.NewHttpContextMock(responseWriter, request)

	modelProduct := model.Product{
		ID:   productID,
		Type: productType,
		Name: productName,
	}
	modelProducts := []model.Product{modelProduct}
	productsJSON, err := json.Marshal(modelProducts)
	require.NoError(t, err)

	err = FindAllProductsTranslated()(productTranslatedInteractorMock.FindAll)(context.ResponseWriter)
	require.NoError(t, err)

	productTranslatedInteractorMock.AssertExpectations(t)

	assert.Equal(t, http.StatusOK, responseWriter.Code)
	assert.Equal(t, string(productsJSON)+"\n", responseWriter.Body.String())
}
