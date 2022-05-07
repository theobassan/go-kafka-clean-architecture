package registry

import (
	"go-kafka-clean-architecture/app/usecases/translator"
)

func (r *Registry) NewProductBrasilTranslator() translator.ProductTranslator {
	return translator.NewProductBrasilTranslator()
}

func (r *Registry) NewProductChileTranslator() translator.ProductTranslator {
	return translator.NewProductChileTranslator()
}
