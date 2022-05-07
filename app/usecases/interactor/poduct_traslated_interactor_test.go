package interactor

import (
	"go-kafka-clean-architecture/app/entities"
	"go-kafka-clean-architecture/app/usecases/repository"

	"github.com/go-errors/errors"

	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestProductTranslatedInteractorCreate_shoudlCreate(t *testing.T) {
	createdID := int64(1)
	product := &entities.Product{}

	productTranslatedRepositoryMock := new(repository.ProductTranslatedRepositoryMock)

	productTranslatedRepositoryMock.On("Create", product).Return(&createdID, nil)

	productTranslatedInteractor := NewProductTranslatedInteractor(
		productTranslatedRepositoryMock,
	)

	returnedId, err := productTranslatedInteractor.Create(product)
	require.NoError(t, err)

	productTranslatedRepositoryMock.AssertExpectations(t)

	assert.Equal(t, *returnedId, createdID)
}

func TestProductTranslatedInteractorCreate_shoudlReturnError_whenRepositoryReturnError(t *testing.T) {
	product := &entities.Product{}
	err := "CreateError"

	productTranslatedRepositoryMock := new(repository.ProductTranslatedRepositoryMock)

	productTranslatedRepositoryMock.On("Create", product).Return(nil, errors.New(err))

	productTranslatedInteractor := NewProductTranslatedInteractor(
		productTranslatedRepositoryMock,
	)

	returnedId, returnedErr := productTranslatedInteractor.Create(product)
	assert.Nil(t, returnedId)

	productTranslatedRepositoryMock.AssertExpectations(t)

	assert.EqualError(t, returnedErr, err)
}

func TestProductTranslatedInteractorFindAll_shouldFindAll(t *testing.T) {
	productID := int64(123)
	productType := string("Type")
	productName := string("Name")

	products := []*entities.Product{}
	products = append(products, &entities.Product{
		ID:   &productID,
		Type: &productType,
		Name: &productName,
	})

	productTranslatedRepositoryMock := new(repository.ProductTranslatedRepositoryMock)
	productTranslatedRepositoryMock.On("FindAll").Return(products, nil)

	productTranslatedInteractor := NewProductTranslatedInteractor(
		productTranslatedRepositoryMock,
	)

	returnedProducts, err := productTranslatedInteractor.FindAll()
	require.NoError(t, err)

	productTranslatedRepositoryMock.AssertExpectations(t)

	assert.Len(t, returnedProducts, 1)
	assert.Equal(t, *returnedProducts[0].ID, productID)
	assert.Equal(t, *returnedProducts[0].Type, productType)
	assert.Equal(t, *returnedProducts[0].Name, productName)
}

func TestProductTranslatedInteractorFindAll_shoudlReturnError_whenRepositoryReturnError(t *testing.T) {
	err := "FindAllError"

	productTranslatedRepositoryMock := new(repository.ProductTranslatedRepositoryMock)

	productTranslatedRepositoryMock.On("FindAll").Return(nil, errors.New(err))

	productTranslatedInteractor := NewProductTranslatedInteractor(
		productTranslatedRepositoryMock,
	)

	returnedProducts, returnedErr := productTranslatedInteractor.FindAll()
	assert.Nil(t, returnedProducts)

	productTranslatedRepositoryMock.AssertExpectations(t)

	assert.EqualError(t, returnedErr, err)
}
