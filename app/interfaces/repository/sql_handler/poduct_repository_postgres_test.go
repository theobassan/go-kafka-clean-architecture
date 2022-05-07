package sql_handler

import (
	"go-kafka-clean-architecture/app/entities"
	"go-kafka-clean-architecture/app/interfaces/database"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestProductRepositoryPostgresCreate_shoudlCreate(t *testing.T) {
	productID := int64(123)
	productType := "Type"
	productName := "Name"

	product := &entities.Product{
		ID:   &productID,
		Type: &productType,
		Name: &productName,
	}

	db, dbMock, err := database.NewSqlHandlerMock()
	require.NoError(t, err)

	createdID := int64(1)
	dbMock.ExpectQuery(
		regexp.QuoteMeta(`
			INSERT INTO
				products(external_id, type, name)
			VALUES
				($1, $2, $3)
			RETURNING
				id
		`)).
		WithArgs(productID, productType, productName).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).
			AddRow(createdID))

	productRepositoryPostgres := NewProductRepositoryPostgres(db)

	returnedId, err := productRepositoryPostgres.Create(product)
	require.NoError(t, err)

	assert.Equal(t, *returnedId, createdID)
}

func TestProductRepositoryPostgresFindAll_shoudlFindAll(t *testing.T) {
	productID := int64(123)
	productType := "Type"
	productName := "Name"

	db, dbMock, err := database.NewSqlHandlerMock()
	require.NoError(t, err)

	dbMock.ExpectQuery(
		regexp.QuoteMeta(`
		SELECT
			external_id,
			type,
			name
		FROM
			products
		`)).
		WillReturnRows(sqlmock.NewRows([]string{"external_id", "type", "name"}).
			AddRow(productID, productType, productName))

	productRepositoryPostgres := NewProductRepositoryPostgres(db)

	returnedProducts, err := productRepositoryPostgres.FindAll()
	require.NoError(t, err)

	assert.Len(t, returnedProducts, 1)
	assert.Equal(t, *returnedProducts[0].ID, productID)
	assert.Equal(t, *returnedProducts[0].Type, productType)
	assert.Equal(t, *returnedProducts[0].Name, productName)
}
