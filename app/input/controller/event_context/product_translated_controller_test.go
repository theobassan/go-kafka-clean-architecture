package event_context

import (
	"go-kafka-clean-architecture/app/entities"
	"go-kafka-clean-architecture/app/input/controller/event_context/model"
	"go-kafka-clean-architecture/app/input/router"
	"go-kafka-clean-architecture/app/input/usecases"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProductControllerCreate_shoudlCreate(t *testing.T) {
	//var productID *int64
	createdID := int64(1)

	eventContextMock := new(router.EventContextMock)
	productTranslatedInteractorMock := new(usecases.ProductTranslatedInteractorMock)

	product := &model.Product{}

	//TODO: mock bind to set product value
	eventContextMock.On("Bind", product).Return(nil)

	productCreate := &entities.Product{}
	productTranslatedInteractorMock.On("Create", productCreate).Return(&createdID, nil)

	eventContextMock.On("Acknowledge").Return(nil)

	productController := NewProductTranslatedController(productTranslatedInteractorMock)

	err := productController.Create(eventContextMock)
	assert.NoError(t, err)

	eventContextMock.AssertExpectations(t)
	productTranslatedInteractorMock.AssertExpectations(t)
}
