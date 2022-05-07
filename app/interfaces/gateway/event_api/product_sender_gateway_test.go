package event_api

import (
	"encoding/json"
	"go-kafka-clean-architecture/app/entities"
	"go-kafka-clean-architecture/app/interfaces/api"
	"go-kafka-clean-architecture/app/interfaces/gateway/event_api/model"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestProductSenderGatewaySend_shoudlSend(t *testing.T) {
	productID := int64(123)
	productType := "Type"
	productName := "Name"

	modelProduct := &model.Product{
		ID:   &productID,
		Type: &productType,
		Name: &productName,
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

	eventApiMock := new(api.EventApiMock)
	eventApiMock.On("Bind", "product", value).Return(msg)
	eventApiMock.On("WriteMessage", msg).Return(nil)

	productSenderGateway := NewProductSenderGateway(eventApiMock)

	product := &entities.Product{
		ID:   &productID,
		Type: &productType,
		Name: &productName,
	}
	err = productSenderGateway.Send(product)
	require.NoError(t, err)

	eventApiMock.AssertExpectations(t)
}
