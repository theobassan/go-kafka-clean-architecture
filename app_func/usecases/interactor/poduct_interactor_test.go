package interactor

import (
	"go-kafka-clean-architecture/app_func/entities"
	"go-kafka-clean-architecture/app_func/usecases/gateway"
	"go-kafka-clean-architecture/app_func/usecases/repository"
	"go-kafka-clean-architecture/app_func/usecases/translator"
	"strconv"
	"testing"

	"github.com/go-errors/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestProductInteractorCreate_shoudlCreate(t *testing.T) {
	productID := int64(123)
	createdID := int64(1)
	product := entities.Product{}

	productFinderGatewayMock := new(gateway.ProductFinderGatewayMock)
	productRepositoryMock := new(repository.ProductRepositoryMock)
	productTranslatorMock := new(translator.ProductTranslatorMock)
	productSenderGatewayMock := new(gateway.ProductSenderGatewayMock)

	productFinderGatewayMock.On("FindById", productID).Return(product, nil)
	productRepositoryMock.On("Create", product).Return(createdID, nil)
	productTranslatorMock.On("Translate", product).Return(product, nil)
	productSenderGatewayMock.On("Send", product).Return(nil)

	returnedId, err := CreateProduct()(productFinderGatewayMock.FindById)(productRepositoryMock.Create)(productTranslatorMock.Translate)(productSenderGatewayMock.Send)(productID)
	require.NoError(t, err)

	productFinderGatewayMock.AssertExpectations(t)
	productRepositoryMock.AssertExpectations(t)
	productTranslatorMock.AssertExpectations(t)
	productSenderGatewayMock.AssertExpectations(t)

	assert.Equal(t, returnedId, createdID)
}

func TestProductInteractorCreate_shoudlReturnError_whenGatewayReturnError(t *testing.T) {
	mockErr := "FindByIdError"
	productID := int64(123)

	productFinderGatewayMock := new(gateway.ProductFinderGatewayMock)

	productFinderGatewayMock.On("FindById", productID).Return(entities.Product{}, errors.New(mockErr))

	returnedId, err := CreateProduct()(productFinderGatewayMock.FindById)(nil)(nil)(nil)(productID)
	assert.Equal(t, returnedId, int64(0))

	productFinderGatewayMock.AssertExpectations(t)

	assert.EqualError(t, err, mockErr)
}

func TestProductInteractorCreate_shoudlReturnError_whenRepositoryReturnError(t *testing.T) {
	mockErr := "CreateError"
	productID := int64(123)
	product := entities.Product{}

	productFinderGatewayMock := new(gateway.ProductFinderGatewayMock)
	productRepositoryMock := new(repository.ProductRepositoryMock)

	productFinderGatewayMock.On("FindById", productID).Return(product, nil)
	productRepositoryMock.On("Create", product).Return(int64(0), errors.New(mockErr))

	returnedId, err := CreateProduct()(productFinderGatewayMock.FindById)(productRepositoryMock.Create)(nil)(nil)(productID)
	assert.Equal(t, returnedId, int64(0))

	productFinderGatewayMock.AssertExpectations(t)
	productRepositoryMock.AssertExpectations(t)

	assert.EqualError(t, err, mockErr)
}

func TestProductInteractorCreate_shoudlReturnError_whenPublisherReturnError(t *testing.T) {
	mockErr := "PublishError"
	productID := int64(123)
	createdID := int64(1)
	product := entities.Product{}

	productFinderGatewayMock := new(gateway.ProductFinderGatewayMock)
	productRepositoryMock := new(repository.ProductRepositoryMock)
	productTranslatorMock := new(translator.ProductTranslatorMock)
	productSenderGatewayMock := new(gateway.ProductSenderGatewayMock)

	productFinderGatewayMock.On("FindById", productID).Return(product, nil)
	productRepositoryMock.On("Create", product).Return(createdID, nil)
	productTranslatorMock.On("Translate", product).Return(product, nil)
	productSenderGatewayMock.On("Send", product).Return(errors.New(mockErr))

	returnedId, err := CreateProduct()(productFinderGatewayMock.FindById)(productRepositoryMock.Create)(productTranslatorMock.Translate)(productSenderGatewayMock.Send)(productID)
	assert.Equal(t, returnedId, int64(0))

	productFinderGatewayMock.AssertExpectations(t)
	productRepositoryMock.AssertExpectations(t)
	productTranslatorMock.AssertExpectations(t)
	productSenderGatewayMock.AssertExpectations(t)

	assert.EqualError(t, err, mockErr)
}

func TestProductInteractorFindAll_shouldFindAll(t *testing.T) {
	productID := int64(123)
	productType := string("Type")
	productName := string("Name")

	products := []entities.Product{}
	products = append(products, entities.Product{
		ID:   productID,
		Type: productType,
		Name: productName,
	})

	productRepositoryMock := new(repository.ProductRepositoryMock)
	productRepositoryMock.On("FindAll").Return(products, nil)

	returnedProducts, err := FindAllProducts()(productRepositoryMock.FindAll)()
	require.NoError(t, err)

	productRepositoryMock.AssertExpectations(t)

	assert.Len(t, returnedProducts, 1)
	assert.Equal(t, returnedProducts[0].ID, productID)
	assert.Equal(t, returnedProducts[0].Type, productType)
	assert.Equal(t, returnedProducts[0].Name, productName)
}

func TestProductInteractorFindAll_shoudlReturnError_whenRepositoryReturnError(t *testing.T) {
	mockErr := "FindAllError"

	productRepositoryMock := new(repository.ProductRepositoryMock)
	productRepositoryMock.On("FindAll").Return(nil, errors.New(mockErr))

	returnedProducts, err := FindAllProducts()(productRepositoryMock.FindAll)()
	assert.Nil(t, returnedProducts)

	productRepositoryMock.AssertExpectations(t)

	assert.EqualError(t, err, mockErr)
}

func TestProductInteractorGet_shouldGet(t *testing.T) {
	productID := int64(123)
	productType := "Type " + strconv.FormatInt(productID, 10)
	productName := "Name " + strconv.FormatInt(productID, 10)

	returnedProduct, err := GetProduct()(productID)
	require.NoError(t, err)

	assert.Equal(t, returnedProduct.ID, productID)
	assert.Equal(t, returnedProduct.Type, productType)
	assert.Equal(t, returnedProduct.Name, productName)
}
