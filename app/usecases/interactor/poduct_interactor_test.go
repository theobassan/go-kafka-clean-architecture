package interactor

import (
	"go-kafka-clean-architecture/app/entities"
	"go-kafka-clean-architecture/app/usecases/gateway"
	"go-kafka-clean-architecture/app/usecases/repository"
	"go-kafka-clean-architecture/app/usecases/translator"
	"strconv"
	"testing"

	"github.com/go-errors/errors"
	"github.com/stretchr/testify/assert"
)

func TestProductInteractorCreate_shoudlCreate(t *testing.T) {
	id := int64(123)
	createdID := int64(1)
	product := &entities.Product{}

	productFinderGatewayMock := new(gateway.ProductFinderGatewayMock)
	productRepositoryMock := new(repository.ProductRepositoryMock)
	productTranslatorMock := new(translator.ProductTranslatorMock)
	productSenderGatewayMock := new(gateway.ProductSenderGatewayMock)

	productFinderGatewayMock.On("FindById", &id).Return(product, nil)
	productRepositoryMock.On("Create", product).Return(&createdID, nil)
	productTranslatorMock.On("Translate", product).Return(product, nil)
	productSenderGatewayMock.On("Send", product).Return(nil)

	productInteractor := NewProductInteractor(
		productFinderGatewayMock,
		productRepositoryMock,
		productSenderGatewayMock,
		productTranslatorMock,
	)

	returnedId, err := productInteractor.Create(&id)
	assert.NoError(t, err)

	productFinderGatewayMock.AssertExpectations(t)
	productRepositoryMock.AssertExpectations(t)
	productTranslatorMock.AssertExpectations(t)
	productSenderGatewayMock.AssertExpectations(t)

	assert.Equal(t, *returnedId, createdID)
}

func TestProductInteractorCreate_shoudlReturnError_whenGatewayReturnError(t *testing.T) {
	id := int64(123)
	err := "FindByIdError"

	productFinderGatewayMock := new(gateway.ProductFinderGatewayMock)
	productRepositoryMock := new(repository.ProductRepositoryMock)
	productSenderGatewayMock := new(gateway.ProductSenderGatewayMock)
	productTranslatorMock := new(translator.ProductTranslatorMock)

	productFinderGatewayMock.On("FindById", &id).Return(nil, errors.New(err))

	productInteractor := NewProductInteractor(
		productFinderGatewayMock,
		productRepositoryMock,
		productSenderGatewayMock,
		productTranslatorMock,
	)

	returnedId, returnedErr := productInteractor.Create(&id)
	assert.Nil(t, returnedId)

	productFinderGatewayMock.AssertExpectations(t)

	assert.EqualError(t, returnedErr, err)
}

func TestProductInteractorCreate_shoudlReturnError_whenRepositoryReturnError(t *testing.T) {
	id := int64(123)
	product := &entities.Product{}
	err := "CreateError"

	productFinderGatewayMock := new(gateway.ProductFinderGatewayMock)
	productRepositoryMock := new(repository.ProductRepositoryMock)
	productSenderGatewayMock := new(gateway.ProductSenderGatewayMock)
	productTranslatorMock := new(translator.ProductTranslatorMock)

	productFinderGatewayMock.On("FindById", &id).Return(product, nil)
	productRepositoryMock.On("Create", product).Return(nil, errors.New(err))

	productInteractor := NewProductInteractor(
		productFinderGatewayMock,
		productRepositoryMock,
		productSenderGatewayMock,
		productTranslatorMock,
	)

	returnedId, returnedErr := productInteractor.Create(&id)
	assert.Nil(t, returnedId)

	productFinderGatewayMock.AssertExpectations(t)
	productRepositoryMock.AssertExpectations(t)

	assert.EqualError(t, returnedErr, err)
}

func TestProductInteractorCreate_shoudlReturnError_whenPublisherReturnError(t *testing.T) {
	id := int64(123)
	createdID := int64(1)
	product := &entities.Product{}
	err := "PublishError"

	productFinderGatewayMock := new(gateway.ProductFinderGatewayMock)
	productRepositoryMock := new(repository.ProductRepositoryMock)
	productSenderGatewayMock := new(gateway.ProductSenderGatewayMock)
	productTranslatorMock := new(translator.ProductTranslatorMock)

	productFinderGatewayMock.On("FindById", &id).Return(product, nil)
	productRepositoryMock.On("Create", product).Return(&createdID, nil)
	productTranslatorMock.On("Translate", product).Return(product, nil)
	productSenderGatewayMock.On("Send", product).Return(errors.New(err))

	productInteractor := NewProductInteractor(
		productFinderGatewayMock,
		productRepositoryMock,
		productSenderGatewayMock,
		productTranslatorMock,
	)

	returnedId, returnedErr := productInteractor.Create(&id)
	assert.Nil(t, returnedId)

	productFinderGatewayMock.AssertExpectations(t)
	productRepositoryMock.AssertExpectations(t)
	productTranslatorMock.AssertExpectations(t)
	productSenderGatewayMock.AssertExpectations(t)

	assert.EqualError(t, returnedErr, err)
}

func TestProductInteractorFindAll_shouldFindAll(t *testing.T) {
	productID := int64(123)
	productType := string("Type")
	productName := string("Name")

	products := []*entities.Product{}
	products = append(products, &entities.Product{
		ID:   &productID,
		Type: &productType,
		Name: &productName,
	})

	productFinderGatewayMock := new(gateway.ProductFinderGatewayMock)
	productRepositoryMock := new(repository.ProductRepositoryMock)
	productSenderGatewayMock := new(gateway.ProductSenderGatewayMock)
	productTranslatorMock := new(translator.ProductTranslatorMock)

	productRepositoryMock.On("FindAll").Return(products, nil)

	productInteractor := NewProductInteractor(
		productFinderGatewayMock,
		productRepositoryMock,
		productSenderGatewayMock,
		productTranslatorMock,
	)

	returnedProducts, err := productInteractor.FindAll()
	assert.NoError(t, err)

	productRepositoryMock.AssertExpectations(t)

	assert.Len(t, returnedProducts, 1)

	returnedProduct := returnedProducts[0]
	assert.Equal(t, *returnedProduct.ID, productID)
	assert.Equal(t, *returnedProduct.Type, productType)
	assert.Equal(t, *returnedProduct.Name, productName)
}

func TestProductInteractorFindAll_shoudlReturnError_whenRepositoryReturnError(t *testing.T) {
	err := "FindAllError"

	productFinderGatewayMock := new(gateway.ProductFinderGatewayMock)
	productRepositoryMock := new(repository.ProductRepositoryMock)
	productSenderGatewayMock := new(gateway.ProductSenderGatewayMock)
	productTranslatorMock := new(translator.ProductTranslatorMock)

	productRepositoryMock.On("FindAll").Return(nil, errors.New(err))

	productInteractor := NewProductInteractor(
		productFinderGatewayMock,
		productRepositoryMock,
		productSenderGatewayMock,
		productTranslatorMock,
	)

	returnedProducts, returnedErr := productInteractor.FindAll()
	assert.Nil(t, returnedProducts)
	assert.EqualError(t, returnedErr, err)

	productFinderGatewayMock.AssertExpectations(t)
}

func TestProductInteractorGet_shouldGet(t *testing.T) {
	productID := int64(123)
	productType := "Type " + strconv.FormatInt(productID, 10)
	productName := "Name " + strconv.FormatInt(productID, 10)

	productFinderGatewayMock := new(gateway.ProductFinderGatewayMock)
	productRepositoryMock := new(repository.ProductRepositoryMock)
	productSenderGatewayMock := new(gateway.ProductSenderGatewayMock)
	productTranslatorMock := new(translator.ProductTranslatorMock)

	productInteractor := NewProductInteractor(
		productFinderGatewayMock,
		productRepositoryMock,
		productSenderGatewayMock,
		productTranslatorMock,
	)

	returnedProduct, err := productInteractor.Get(&productID)
	assert.NoError(t, err)

	assert.Equal(t, *returnedProduct.ID, productID)
	assert.Equal(t, *returnedProduct.Type, productType)
	assert.Equal(t, *returnedProduct.Name, productName)
}
