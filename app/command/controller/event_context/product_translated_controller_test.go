package event_context

import (
	"go-kafka-clean-architecture/app/command/controller/event_context/model"
	"go-kafka-clean-architecture/app/command/router"
	"go-kafka-clean-architecture/app/command/usecases"
	"go-kafka-clean-architecture/app/entities"
	"testing"

	"github.com/stretchr/testify/require"
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
	require.NoError(t, err)

	eventContextMock.AssertExpectations(t)
	productTranslatedInteractorMock.AssertExpectations(t)
}
