package registry

import (
	"go-kafka-clean-architecture/app/input/controller/event_context"
	"go-kafka-clean-architecture/app/input/controller/http_context"
	"go-kafka-clean-architecture/app/interfaces/api"
	"go-kafka-clean-architecture/app/interfaces/database"
)

type Registry struct {
}

func NewRegistry() *Registry {
	return &Registry{}
}

func (r *Registry) NewHttpContextRestSqlEventAppController(restAPI api.RestAPI, sqlHandler database.SQLHandler, eventAPI api.EventAPI) *http_context.AppController {
	//interactor := r.NewRestSqlEventProductInteractor(restAPI, sqlHandler, eventAPI, r.NewProductChileTranslator())
	interactor := r.NewRestSqlEventProductInteractor(restAPI, sqlHandler, eventAPI, r.NewProductBrasilTranslator())
	interactorTranslated := r.NewRestSqlEventProductTraslatedInteractor(sqlHandler)

	productController := http_context.NewProductController(interactor)
	productTranslatedController := http_context.NewProductTranslatedController(interactorTranslated)

	return http_context.NewAppController(productController, productTranslatedController)
}

func (r *Registry) NewHttpContextRestGormEventAppController(restAPI api.RestAPI, sqlGorm database.SQLGorm, eventAPI api.EventAPI) *http_context.AppController {
	//interactor := r.NewRestGormEventProductInteractor(restAPI, sqlGorm, eventAPI, r.NewProductChileTranslator())
	interactor := r.NewRestGormEventProductInteractor(restAPI, sqlGorm, eventAPI, r.NewProductBrasilTranslator())
	interactorTranslated := r.NewRestGormEventProductTraslatedInteractor(sqlGorm)

	productController := http_context.NewProductController(interactor)
	productTranslatedController := http_context.NewProductTranslatedController(interactorTranslated)

	return http_context.NewAppController(productController, productTranslatedController)
}

func (r *Registry) NewEventContextRestSqlEventAppController(restAPI api.RestAPI, sqlHandler database.SQLHandler, eventAPI api.EventAPI) *event_context.AppController {
	interactorTranslated := r.NewRestSqlEventProductTraslatedInteractor(sqlHandler)
	productTranslatedController := event_context.NewProductTranslatedController(interactorTranslated)

	return event_context.NewAppController(productTranslatedController)
}

func (r *Registry) NewEventContextRestGormEventAppController(restAPI api.RestAPI, sqlGorm database.SQLGorm, eventAPI api.EventAPI) *event_context.AppController {
	interactorTranslated := r.NewRestGormEventProductTraslatedInteractor(sqlGorm)
	productTranslatedController := event_context.NewProductTranslatedController(interactorTranslated)

	return event_context.NewAppController(productTranslatedController)
}
