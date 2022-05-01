package entities

type Product struct {
	ID   *int64
	Type *string
	Name *string
}

func (product *Product) TranslateToBrasil() *Product {

	*product.Type = *product.Type + " Brasil"
	*product.Name = *product.Name + " Brasil"

	return product
}

func (product *Product) TranslateToChile() *Product {

	*product.Type = *product.Type + " Chile"
	*product.Name = *product.Name + " Chile"

	return product
}
