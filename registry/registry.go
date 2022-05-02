package registry

import (
	"go-kafka-clean-architecture/app/command/controller/event_context"
	"go-kafka-clean-architecture/app/command/controller/http_context"
	"go-kafka-clean-architecture/app/interfaces/api"
	"go-kafka-clean-architecture/app/interfaces/database"
)

type Registry struct {
}

func NewRegistry() *Registry {
	return &Registry{}
}

func (r *Registry) NewHttpContextRestSqlEventAppControllerMySql(restAPI api.RestAPI, sqlHandler database.SQLHandler, eventAPI api.EventAPI) *http_context.AppController {
	//interactor := r.NewRestSqlEventProductInteractor(restAPI, sqlHandler, eventAPI, r.NewProductChileTranslator())
	interactor := r.NewRestSqlEventProductInteractorMySql(restAPI, sqlHandler, eventAPI, r.NewProductBrasilTranslator())
	interactorTranslated := r.NewRestSqlEventProductTraslatedInteractorMySql(sqlHandler)

	productController := http_context.NewProductController(interactor)
	productTranslatedController := http_context.NewProductTranslatedController(interactorTranslated)

	return http_context.NewAppController(productController, productTranslatedController)
}

func (r *Registry) NewEventContextRestSqlEventAppControllerMySql(restAPI api.RestAPI, sqlHandler database.SQLHandler, eventAPI api.EventAPI) *event_context.AppController {
	interactorTranslated := r.NewRestSqlEventProductTraslatedInteractorMySql(sqlHandler)
	productTranslatedController := event_context.NewProductTranslatedController(interactorTranslated)

	return event_context.NewAppController(productTranslatedController)
}

func (r *Registry) NewHttpContextRestSqlEventAppControllerPostgres(restAPI api.RestAPI, sqlHandler database.SQLHandler, eventAPI api.EventAPI) *http_context.AppController {
	//interactor := r.NewRestSqlEventProductInteractor(restAPI, sqlHandler, eventAPI, r.NewProductChileTranslator())
	interactor := r.NewRestSqlEventProductInteractorPostgres(restAPI, sqlHandler, eventAPI, r.NewProductBrasilTranslator())
	interactorTranslated := r.NewRestSqlEventProductTraslatedInteractorPostgres(sqlHandler)

	productController := http_context.NewProductController(interactor)
	productTranslatedController := http_context.NewProductTranslatedController(interactorTranslated)

	return http_context.NewAppController(productController, productTranslatedController)
}

func (r *Registry) NewEventContextRestSqlEventAppControllerPostgres(restAPI api.RestAPI, sqlHandler database.SQLHandler, eventAPI api.EventAPI) *event_context.AppController {
	interactorTranslated := r.NewRestSqlEventProductTraslatedInteractorPostgres(sqlHandler)
	productTranslatedController := event_context.NewProductTranslatedController(interactorTranslated)

	return event_context.NewAppController(productTranslatedController)
}

func (r *Registry) NewHttpContextRestGormEventAppController(restAPI api.RestAPI, sqlGorm database.SQLGorm, eventAPI api.EventAPI) *http_context.AppController {
	//interactor := r.NewRestGormEventProductInteractor(restAPI, sqlGorm, eventAPI, r.NewProductChileTranslator())
	interactor := r.NewRestGormEventProductInteractor(restAPI, sqlGorm, eventAPI, r.NewProductBrasilTranslator())
	interactorTranslated := r.NewRestGormEventProductTraslatedInteractor(sqlGorm)

	productController := http_context.NewProductController(interactor)
	productTranslatedController := http_context.NewProductTranslatedController(interactorTranslated)

	return http_context.NewAppController(productController, productTranslatedController)
}

func (r *Registry) NewEventContextRestGormEventAppController(restAPI api.RestAPI, sqlGorm database.SQLGorm, eventAPI api.EventAPI) *event_context.AppController {
	interactorTranslated := r.NewRestGormEventProductTraslatedInteractor(sqlGorm)
	productTranslatedController := event_context.NewProductTranslatedController(interactorTranslated)

	return event_context.NewAppController(productTranslatedController)
}
