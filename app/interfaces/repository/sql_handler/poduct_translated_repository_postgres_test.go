package sql_handler

import (
	"go-kafka-clean-architecture/app/entities"
	"go-kafka-clean-architecture/app/interfaces/database"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestProductTranslatedRepositoryPostgresCreate_shoudlCreate(t *testing.T) {
	productID := int64(123)
	productType := "Type"
	productName := "Name"

	product := &entities.Product{
		ID:   &productID,
		Type: &productType,
		Name: &productName,
	}

	db, dbMock, err := database.NewSQLHandlerMock()
	assert.NoError(t, err)

	createdID := int64(1)
	dbMock.ExpectQuery(
		regexp.QuoteMeta(`
			INSERT INTO
				products_translated(external_id, type, name)
			VALUES
				($1, $2, $3)
			RETURNING
				id
		`)).
		WithArgs(productID, productType, productName).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).
			AddRow(createdID))

	productTranslatedRepositoryPostgres := NewProductTranslatedRepositoryPostgres(db)

	returnedId, err := productTranslatedRepositoryPostgres.Create(product)
	assert.NoError(t, err)
	assert.Equal(t, *returnedId, createdID)
}

func TestProductTranslatedRepositoryPostgresFindAll_shoudlFindAll(t *testing.T) {
	productID := int64(123)
	productType := "Type"
	productName := "Name"

	db, dbMock, err := database.NewSQLHandlerMock()
	assert.NoError(t, err)

	dbMock.ExpectQuery(
		regexp.QuoteMeta(`
		SELECT
			external_id,
			type,
			name
		FROM
			products_translated
		`)).
		WillReturnRows(sqlmock.NewRows([]string{"external_id", "type", "name"}).
			AddRow(productID, productType, productName))

	productTranslatedRepositoryPostgres := NewProductTranslatedRepositoryPostgres(db)

	returnedProduct, err := productTranslatedRepositoryPostgres.FindAll()
	assert.NoError(t, err)
	assert.Equal(t, *returnedProduct[0].ID, productID)
	assert.Equal(t, *returnedProduct[0].Type, productType)
	assert.Equal(t, *returnedProduct[0].Name, productName)
}
