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
	"github.com/stretchr/testify/require"
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
	require.NoError(t, err)

	mockResponse := &http.Response{
		StatusCode: 200,
		Body:       ioutil.NopCloser(bytes.NewReader(responseJson)),
	}

	restApiMock := new(api.RestApiMock)
	restApiMock.On("Get", "product?id=123").Return(mockResponse, nil)

	productFinderGateway := NewProductFinderGateway(restApiMock)

	returnedProuct, err := productFinderGateway.FindById(&productID)
	require.NoError(t, err)

	restApiMock.AssertExpectations(t)

	assert.Equal(t, *returnedProuct.ID, productID)
	assert.Equal(t, *returnedProuct.Type, productType)
	assert.Equal(t, *returnedProuct.Name, productName)
}
