package rest_api

import (
	"bytes"
	"encoding/json"
	"go-kafka-clean-architecture/app/entities"
	"go-kafka-clean-architecture/app/interfaces/api"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProductFinderGatewayFindById_shoudlFindById(t *testing.T) {
	productID := int64(123)
	productType := "Type"
	productName := "Name"

	product := &entities.Product{
		ID:   &productID,
		Type: &productType,
		Name: &productName,
	}
	responseJson, err := json.Marshal(product)
	assert.NoError(t, err)

	mockResponse := &http.Response{
		StatusCode: 200,
		Body:       ioutil.NopCloser(bytes.NewReader(responseJson)),
	}

	restAPIMock := new(api.RestAPIMock)
	restAPIMock.On("Get", "product?id=123").Return(mockResponse, nil)

	productFinderGateway := NewProductFinderGateway(restAPIMock)

	returnedProuct, err := productFinderGateway.FindById(&productID)
	assert.NoError(t, err)

	restAPIMock.AssertExpectations(t)

	assert.Equal(t, *returnedProuct.ID, productID)
	assert.Equal(t, *returnedProuct.Type, productType)
	assert.Equal(t, *returnedProuct.Name, productName)
}
