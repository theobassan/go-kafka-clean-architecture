package test

import (
	"encoding/json"
	"go-kafka-clean-architecture/app/interfaces/gateway/rest_api/model"
	"io"
	"net/http"
	"strings"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreate(t *testing.T, serverURL string, productID int64) {
	product := &model.Product{
		ID: &productID,
	}

	productJSON, err := json.Marshal(product)
	require.NoError(t, err)

	responseCreate, err := http.Post(serverURL+"/product", "application/json", strings.NewReader(string(productJSON)))
	require.NoError(t, err)
	assert.Equal(t, http.StatusCreated, responseCreate.StatusCode)

	responseFindAll, err := http.Get(serverURL + "/products")
	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, responseFindAll.StatusCode)

	time.Sleep(5 * time.Second)

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
	assert.Equal(t, productID, *products[0].ID)

	assert.Len(t, productsTranslated, 1)
	assert.Equal(t, productID, *products[0].ID)
}
