package registry

import (
	"go-kafka-clean-architecture/app_func/translator/brasil"
	"go-kafka-clean-architecture/app_func/translator/chile"
	"go-kafka-clean-architecture/app_func/usecases/translator"
)

func TranslateProductToChile() translator.ProductTranslatorTranslate {
	return chile.TranslateProduct()
}

func TranslateProductToBrasil() translator.ProductTranslatorTranslate {
	return brasil.TranslateProduct()
}
