package http_context

import (
	"encoding/json"
	"go-kafka-clean-architecture/app/entities"
	"go-kafka-clean-architecture/app/input/controller/http_context/model"
	"go-kafka-clean-architecture/app/input/router"
	"go-kafka-clean-architecture/app/input/usecases"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProductTranslatedControllerFindAll_shoudlFindAll(t *testing.T) {
	productID := int64(1)
	productType := "Type"
	productName := "Name"
	product := &entities.Product{
		ID:   &productID,
		Type: &productType,
		Name: &productName,
	}
	products := []*entities.Product{product}
	productTranslatedInteractorMock := new(usecases.ProductTranslatedInteractorMock)
	productTranslatedInteractorMock.On("FindAll").Return(products, nil)
	productTranslatedController := NewProductTranslatedController(productTranslatedInteractorMock)

	request := httptest.NewRequest(http.MethodGet, "/productstranslated", nil)
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

	productTranslatedController.FindAll(context)

	productTranslatedInteractorMock.AssertExpectations(t)

	assert.Equal(t, http.StatusOK, responseWriter.Code)
	assert.Equal(t, string(productsJSON)+"\n", responseWriter.Body.String())
}
