package event_api

import (
	"encoding/json"
	"go-kafka-clean-architecture/app/entities"
	"go-kafka-clean-architecture/app/interfaces/api"
	"go-kafka-clean-architecture/app/interfaces/gateway/event_api/model"
	"testing"

	"github.com/stretchr/testify/assert"
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
	assert.NoError(t, err)

	msg := struct {
		Topic string
		Value []byte
	}{
		Topic: "product",
		Value: value,
	}

	eventAPIMock := new(api.EventAPIMock)
	eventAPIMock.On("Bind", "product", value).Return(msg)
	eventAPIMock.On("WriteMessage", msg).Return(nil)

	productSenderGateway := NewProductSenderGateway(eventAPIMock)

	product := &entities.Product{
		ID:   &productID,
		Type: &productType,
		Name: &productName,
	}
	err = productSenderGateway.Send(product)
	assert.NoError(t, err)

	eventAPIMock.AssertExpectations(t)
}
