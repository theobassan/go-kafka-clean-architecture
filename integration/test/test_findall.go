package test

import (
	"encoding/json"
	"go-kafka-clean-architecture/app/interfaces/gateway/rest_api/model"
	"io"
	"net/http"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFindAll(t *testing.T, serverURL string, productID *int64, productType *string, productName *string, productTranslatedID *int64, productTranslatedType *string, productTranslatedName *string) {
	responseFindAll, err := http.Get(serverURL + "/products")
	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, responseFindAll.StatusCode)

	responseFindAllTranslated, err := http.Get(serverURL + "/productstranslated")
	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, responseFindAllTranslated.StatusCode)

	responseFindAllBodyBytes, err := io.ReadAll(responseFindAll.Body)
	require.NoError(t, err)
	products := []*model.Product{}
	err = json.Unmarshal(responseFindAllBodyBytes, &products)
	require.NoError(t, err)

	responseFindAllTranslatedBodyBytes, err := io.ReadAll(responseFindAllTranslated.Body)
	require.NoError(t, err)
	productsTranslated := []*model.Product{}
	err = json.Unmarshal(responseFindAllTranslatedBodyBytes, &productsTranslated)
	require.NoError(t, err)

	assert.Len(t, products, 1)
	assert.Equal(t, *products[0].ID, *productID)
	assert.Equal(t, *products[0].Type, *productType)
	assert.Equal(t, *products[0].Name, *productName)

	assert.Len(t, productsTranslated, 1)
	assert.Equal(t, *productsTranslated[0].ID, *productTranslatedID)
	assert.Equal(t, *productsTranslated[0].Type, *productTranslatedType)
	assert.Equal(t, *productsTranslated[0].Name, *productTranslatedName)
}
