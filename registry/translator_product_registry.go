package registry

import (
	"go-kafka-clean-architecture/app/usecases/interactor"
)

func (r *Registry) NewProductBrasilTranslatorInteractor() interactor.ProductTranslatorInteractor {
	return interactor.NewProductBrasilTranslatorInteractor()
}

func (r *Registry) NewProductChileTranslatorInteractor() interactor.ProductTranslatorInteractor {
	return interactor.NewProductChileTranslatorInteractor()
}
