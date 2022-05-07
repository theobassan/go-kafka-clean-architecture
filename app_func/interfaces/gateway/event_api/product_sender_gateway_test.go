package event_api

import (
	"encoding/json"
	"go-kafka-clean-architecture/app_func/entities"
	"go-kafka-clean-architecture/app_func/interfaces/api"
	"go-kafka-clean-architecture/app_func/interfaces/gateway/event_api/model"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestProductSenderGatewaySend_shoudlSend(t *testing.T) {
	productID := int64(123)
	productType := "Type"
	productName := "Name"

	modelProduct := model.Product{
		ID:   productID,
		Type: productType,
		Name: productName,
	}
	value, err := json.Marshal(modelProduct)
	require.NoError(t, err)

	msg := struct {
		Topic string
		Value []byte
	}{
		Topic: "product",
		Value: value,
	}

	eventApiMock := api.EventApiMock{}
	eventApiMock.On("Bind", "product", value).Return(msg)
	eventApiMock.On("WriteMessage", msg).Return(nil)

	product := entities.Product{
		ID:   productID,
		Type: productType,
		Name: productName,
	}
	err = SendProduct()(eventApiMock.Bind)(eventApiMock.WriteMessage)(product)
	require.NoError(t, err)

	eventApiMock.AssertExpectations(t)
}
