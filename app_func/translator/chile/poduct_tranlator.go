package chile

import (
	"go-kafka-clean-architecture/app_func/entities"
	"go-kafka-clean-architecture/app_func/usecases/translator"
)

func TranslateProduct() translator.ProductTranslatorTranslate {
	return func(product entities.Product) entities.Product {
		return entities.Product{
			ID:   product.ID,
			Type: product.Type + " Chile",
			Name: product.Name + " Chile",
		}
	}
}
