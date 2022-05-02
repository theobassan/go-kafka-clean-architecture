package registry

import (
	"go-kafka-clean-architecture/app/command/usecases"
	"go-kafka-clean-architecture/app/interfaces/api"
	"go-kafka-clean-architecture/app/interfaces/database"
	"go-kafka-clean-architecture/app/usecases/interactor"
	"go-kafka-clean-architecture/app/usecases/translator"
)

func (r *Registry) NewRestSqlEventProductInteractorMySql(restAPI api.RestAPI, sqlHandler database.SQLHandler, eventAPI api.EventAPI, translator translator.ProductTranslator) usecases.ProductInteractor {
	return interactor.NewProductInteractor(r.NewProductFinderGateway(restAPI), r.NewSqlProductRepositoryMySql(sqlHandler), r.NewProductSenderGateway(eventAPI), translator)
}

func (r *Registry) NewRestSqlEventProductTraslatedInteractorMySql(sqlHandler database.SQLHandler) usecases.ProductTranslatedInteractor {
	return interactor.NewProductTranslatedInteractor(r.NewSqlProductTranslatedRepositoryMySql(sqlHandler))
}

func (r *Registry) NewRestSqlEventProductInteractorPostgres(restAPI api.RestAPI, sqlHandler database.SQLHandler, eventAPI api.EventAPI, translator translator.ProductTranslator) usecases.ProductInteractor {
	return interactor.NewProductInteractor(r.NewProductFinderGateway(restAPI), r.NewSqlProductRepositoryPostgres(sqlHandler), r.NewProductSenderGateway(eventAPI), translator)
}

func (r *Registry) NewRestSqlEventProductTraslatedInteractorPostgres(sqlHandler database.SQLHandler) usecases.ProductTranslatedInteractor {
	return interactor.NewProductTranslatedInteractor(r.NewSqlProductTranslatedRepositoryPostgres(sqlHandler))
}

func (r *Registry) NewRestGormEventProductInteractor(restAPI api.RestAPI, sqlGorm database.SQLGorm, eventAPI api.EventAPI, translator translator.ProductTranslator) usecases.ProductInteractor {
	return interactor.NewProductInteractor(r.NewProductFinderGateway(restAPI), r.NewGormProductRepository(sqlGorm), r.NewProductSenderGateway(eventAPI), translator)
}

func (r *Registry) NewRestGormEventProductTraslatedInteractor(sqlGorm database.SQLGorm) usecases.ProductTranslatedInteractor {
	return interactor.NewProductTranslatedInteractor(r.NewGormProductTranslatedRepository(sqlGorm))
}
