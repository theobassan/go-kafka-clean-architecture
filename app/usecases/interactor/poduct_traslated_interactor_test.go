package interactor

import (
	"go-kafka-clean-architecture/app/entities"
	"go-kafka-clean-architecture/app/usecases/repository"

	"github.com/go-errors/errors"

	"testing"

	"github.com/stretchr/testify/assert"
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
	assert.NoError(t, err)

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
	assert.NoError(t, err)

	productTranslatedRepositoryMock.AssertExpectations(t)

	assert.Len(t, returnedProducts, 1)

	returnedProduct := returnedProducts[0]
	assert.Equal(t, *returnedProduct.ID, productID)
	assert.Equal(t, *returnedProduct.Type, productType)
	assert.Equal(t, *returnedProduct.Name, productName)
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
	assert.EqualError(t, returnedErr, err)

	productTranslatedRepositoryMock.AssertExpectations(t)
}
