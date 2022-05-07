package sql_gorm

import (
	"go-kafka-clean-architecture/app/entities"
	"go-kafka-clean-architecture/app/interfaces/database"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestProductRepositoryCreate_shoudlCreateInMySql(t *testing.T) {
	productID := int64(123)
	productType := "Type"
	productName := "Name"

	product := &entities.Product{
		ID:   &productID,
		Type: &productType,
		Name: &productName,
	}

	db, dbMock, err := database.NewSqlGormMock("mysql")
	require.NoError(t, err)

	createdID := int64(1)
	dbMock.ExpectBegin()
	dbMock.ExpectExec(
		regexp.QuoteMeta("INSERT INTO `products` (`external_id`,`type`,`name`) VALUES (?,?,?)")).
		WithArgs(productID, productType, productName).
		WillReturnResult(sqlmock.NewResult(createdID, 1))
	dbMock.ExpectCommit()

	productRepository := NewProductRepository(db)

	returnedId, err := productRepository.Create(product)
	require.NoError(t, err)

	assert.Equal(t, *returnedId, createdID)
}

func TestProductRepositoryCreate_shoudlCreateInPostres(t *testing.T) {
	productID := int64(123)
	productType := "Type"
	productName := "Name"

	product := &entities.Product{
		ID:   &productID,
		Type: &productType,
		Name: &productName,
	}

	db, dbMock, err := database.NewSqlGormMock("postgres")
	require.NoError(t, err)

	createdID := int64(1)
	dbMock.ExpectBegin()
	dbMock.ExpectQuery(
		regexp.QuoteMeta(`INSERT INTO "products" ("external_id","type","name") VALUES ($1,$2,$3) RETURNING "products"."id"`)).
		WithArgs(productID, productType, productName).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).
			AddRow(createdID))
	dbMock.ExpectCommit()

	productRepository := NewProductRepository(db)

	returnedId, err := productRepository.Create(product)
	require.NoError(t, err)

	assert.Equal(t, *returnedId, createdID)
}

func TestProductRepositoryFindAll_shoudlFindAllInMySql(t *testing.T) {
	productID := int64(1)
	productExternalID := int64(123)
	productType := "Type"
	productName := "Name"

	db, dbMock, err := database.NewSqlGormMock("mysql")
	require.NoError(t, err)

	dbMock.ExpectQuery(
		regexp.QuoteMeta("SELECT * FROM `products`")).
		WillReturnRows(sqlmock.NewRows([]string{"id", "external_id", "type", "name"}).
			AddRow(productID, productExternalID, productType, productName))

	productRepository := NewProductRepository(db)

	returnedProduct, err := productRepository.FindAll()
	require.NoError(t, err)

	assert.Equal(t, *returnedProduct[0].ID, productExternalID)
	assert.Equal(t, *returnedProduct[0].Type, productType)
	assert.Equal(t, *returnedProduct[0].Name, productName)
}

func TestProductRepositoryFindAll_shoudlFindAllInPostgres(t *testing.T) {
	productID := int64(1)
	productExternalID := int64(123)
	productType := "Type"
	productName := "Name"

	db, dbMock, err := database.NewSqlGormMock("postgres")
	require.NoError(t, err)

	dbMock.ExpectQuery(
		regexp.QuoteMeta(`SELECT * FROM "products"`)).
		WillReturnRows(sqlmock.NewRows([]string{"id", "external_id", "type", "name"}).
			AddRow(productID, productExternalID, productType, productName))

	productRepository := NewProductRepository(db)

	returnedProducts, err := productRepository.FindAll()
	require.NoError(t, err)

	assert.Len(t, returnedProducts, 1)
	assert.Equal(t, *returnedProducts[0].ID, productExternalID)
	assert.Equal(t, *returnedProducts[0].Type, productType)
	assert.Equal(t, *returnedProducts[0].Name, productName)
}
