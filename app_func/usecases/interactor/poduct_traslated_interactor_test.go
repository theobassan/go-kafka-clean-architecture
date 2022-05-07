package interactor

import (
	"go-kafka-clean-architecture/app_func/entities"
	"go-kafka-clean-architecture/app_func/usecases/repository"

	"github.com/go-errors/errors"

	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestProductTranslatedInteractorCreate_shoudlCreate(t *testing.T) {
	createdID := int64(1)
	product := entities.Product{}

	productTranslatedRepositoryMock := new(repository.ProductTranslatedRepositoryMock)
	productTranslatedRepositoryMock.On("Create", product).Return(createdID, nil)

	returnedId, err := CreateProductTranslated()(productTranslatedRepositoryMock.Create)(product)
	require.NoError(t, err)

	productTranslatedRepositoryMock.AssertExpectations(t)

	assert.Equal(t, returnedId, createdID)
}

func TestProductTranslatedInteractorCreate_shoudlReturnError_whenRepositoryReturnError(t *testing.T) {
	product := entities.Product{}
	mockErr := "CreateError"

	productTranslatedRepositoryMock := new(repository.ProductTranslatedRepositoryMock)
	productTranslatedRepositoryMock.On("Create", product).Return(int64(0), errors.New(mockErr))

	returnedId, err := CreateProductTranslated()(productTranslatedRepositoryMock.Create)(product)
	assert.Equal(t, returnedId, int64(0))

	productTranslatedRepositoryMock.AssertExpectations(t)

	assert.EqualError(t, err, mockErr)
}

func TestProductTranslatedInteractorFindAll_shouldFindAll(t *testing.T) {
	productID := int64(123)
	productType := string("Type")
	productName := string("Name")

	products := []entities.Product{}
	products = append(products, entities.Product{
		ID:   productID,
		Type: productType,
		Name: productName,
	})

	productTranslatedRepositoryMock := new(repository.ProductTranslatedRepositoryMock)
	productTranslatedRepositoryMock.On("FindAll").Return(products, nil)

	returnedProducts, err := FindAllProductsTranslated()(productTranslatedRepositoryMock.FindAll)()
	require.NoError(t, err)

	productTranslatedRepositoryMock.AssertExpectations(t)

	assert.Len(t, returnedProducts, 1)
	assert.Equal(t, returnedProducts[0].ID, productID)
	assert.Equal(t, returnedProducts[0].Type, productType)
	assert.Equal(t, returnedProducts[0].Name, productName)
}

func TestProductTranslatedInteractorFindAll_shoudlReturnError_whenRepositoryReturnError(t *testing.T) {
	mockErr := "FindAllError"

	productTranslatedRepositoryMock := new(repository.ProductTranslatedRepositoryMock)
	productTranslatedRepositoryMock.On("FindAll").Return(nil, errors.New(mockErr))

	returnedProducts, err := FindAllProductsTranslated()(productTranslatedRepositoryMock.FindAll)()
	assert.Nil(t, returnedProducts)

	productTranslatedRepositoryMock.AssertExpectations(t)

	assert.EqualError(t, err, mockErr)
}
