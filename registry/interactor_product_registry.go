package registry

import (
	"go-kafka-clean-architecture/app/command/usecases"
	"go-kafka-clean-architecture/app/interfaces/api"
	"go-kafka-clean-architecture/app/interfaces/database"
	"go-kafka-clean-architecture/app/usecases/interactor"
	"go-kafka-clean-architecture/app/usecases/translator"
)

func (r *Registry) NewRestSqlEventProductInteractor(restAPI api.RestAPI, sqlHandler database.SQLHandler, eventAPI api.EventAPI, translator translator.ProductTranslator) usecases.ProductInteractor {
	return interactor.NewProductInteractor(r.NewProductFinderGateway(restAPI), r.NewSqlProductRepository(sqlHandler), r.NewProductSenderGateway(eventAPI), translator)
}

func (r *Registry) NewRestGormEventProductInteractor(restAPI api.RestAPI, sqlGorm database.SQLGorm, eventAPI api.EventAPI, translator translator.ProductTranslator) usecases.ProductInteractor {
	return interactor.NewProductInteractor(r.NewProductFinderGateway(restAPI), r.NewGormProductRepository(sqlGorm), r.NewProductSenderGateway(eventAPI), translator)
}

func (r *Registry) NewRestSqlEventProductTraslatedInteractor(sqlHandler database.SQLHandler) usecases.ProductTranslatedInteractor {
	return interactor.NewProductTranslatedInteractor(r.NewSqlProductTranslatedRepository(sqlHandler))
}

func (r *Registry) NewRestGormEventProductTraslatedInteractor(sqlGorm database.SQLGorm) usecases.ProductTranslatedInteractor {
	return interactor.NewProductTranslatedInteractor(r.NewGormProductTranslatedRepository(sqlGorm))
}
