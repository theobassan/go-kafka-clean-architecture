package sql_gorm

import (
	"go-kafka-clean-architecture/app/entities"
	"go-kafka-clean-architecture/app/interfaces/database"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestProductTranslatedRepositoryCreate_shoudlCreateInMySQL(t *testing.T) {
	productID := int64(123)
	productType := "Type"
	productName := "Name"

	product := &entities.Product{
		ID:   &productID,
		Type: &productType,
		Name: &productName,
	}

	db, dbMock, err := database.NewSqlGormMock("mysql")
	assert.NoError(t, err)

	createdID := int64(1)
	dbMock.ExpectBegin()
	dbMock.ExpectExec(
		regexp.QuoteMeta("INSERT INTO `products_translated` (`external_id`,`type`,`name`) VALUES (?,?,?)")).
		WithArgs(productID, productType, productName).
		WillReturnResult(sqlmock.NewResult(createdID, 1))
	dbMock.ExpectCommit()

	productTranslatedRepository := NewProductTranslatedRepository(db)

	returnedId, err := productTranslatedRepository.Create(product)
	assert.NoError(t, err)
	assert.Equal(t, *returnedId, createdID)
}

func TestProductTranslatedRepositoryCreate_shoudlCreateInPostres(t *testing.T) {
	productID := int64(123)
	productType := "Type"
	productName := "Name"

	product := &entities.Product{
		ID:   &productID,
		Type: &productType,
		Name: &productName,
	}

	db, dbMock, err := database.NewSqlGormMock("postgres")
	assert.NoError(t, err)

	createdID := int64(1)
	dbMock.ExpectBegin()
	dbMock.ExpectQuery(
		regexp.QuoteMeta(`INSERT INTO "products_translated" ("external_id","type","name") VALUES ($1,$2,$3) RETURNING "products_translated"."id"`)).
		WithArgs(productID, productType, productName).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).
			AddRow(createdID))
	dbMock.ExpectCommit()

	productTranslatedRepository := NewProductTranslatedRepository(db)

	returnedId, err := productTranslatedRepository.Create(product)
	assert.NoError(t, err)
	assert.Equal(t, *returnedId, createdID)
}

func TestProductTranslatedRepositoryFindAll_shoudlFindAllInMySQL(t *testing.T) {
	productID := int64(1)
	productExternalID := int64(123)
	productType := "Type"
	productName := "Name"

	db, dbMock, err := database.NewSqlGormMock("mysql")
	assert.NoError(t, err)

	dbMock.ExpectQuery(
		regexp.QuoteMeta("SELECT * FROM `products_translated`")).
		WillReturnRows(sqlmock.NewRows([]string{"id", "external_id", "type", "name"}).
			AddRow(productID, productExternalID, productType, productName))

	productTranslatedRepository := NewProductTranslatedRepository(db)

	returnedProduct, err := productTranslatedRepository.FindAll()
	assert.NoError(t, err)
	assert.Equal(t, *returnedProduct[0].ID, productExternalID)
	assert.Equal(t, *returnedProduct[0].Type, productType)
	assert.Equal(t, *returnedProduct[0].Name, productName)
}

func TestProductTranslatedRepositoryFindAll_shoudlFindAllInPostgres(t *testing.T) {
	productID := int64(1)
	productExternalID := int64(123)
	productType := "Type"
	productName := "Name"

	db, dbMock, err := database.NewSqlGormMock("postgres")
	assert.NoError(t, err)

	dbMock.ExpectQuery(
		regexp.QuoteMeta(`SELECT * FROM "products_translated"`)).
		WillReturnRows(sqlmock.NewRows([]string{"id", "external_id", "type", "name"}).
			AddRow(productID, productExternalID, productType, productName))

	productTranslatedRepository := NewProductTranslatedRepository(db)

	returnedProducts, err := productTranslatedRepository.FindAll()
	assert.NoError(t, err)

	assert.Len(t, returnedProducts, 1)
	assert.Equal(t, *returnedProducts[0].ID, productExternalID)
	assert.Equal(t, *returnedProducts[0].Type, productType)
	assert.Equal(t, *returnedProducts[0].Name, productName)
}
