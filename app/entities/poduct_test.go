package entities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProductTranslateToBrasil_shoudlTranslate(t *testing.T) {
	productID := int64(123)
	productType := "Type"
	productName := "Name"

	product := &Product{
		ID:   &productID,
		Type: &productType,
		Name: &productName,
	}

	productReturned := product.TranslateToBrasil()
	assert.Equal(t, *productReturned.ID, productID)
	assert.Equal(t, *productReturned.Type, "Type Brasil")
	assert.Equal(t, *productReturned.Name, "Name Brasil")
}

func TestProductTranslateToChile_shoudlTranslate(t *testing.T) {
	productID := int64(123)
	productType := "Type"
	productName := "Name"

	product := &Product{
		ID:   &productID,
		Type: &productType,
		Name: &productName,
	}

	productReturned := product.TranslateToChile()
	assert.Equal(t, *productReturned.ID, productID)
	assert.Equal(t, *productReturned.Type, "Type Chile")
	assert.Equal(t, *productReturned.Name, "Name Chile")
}
