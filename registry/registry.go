package registry

import (
	"go-kafka-clean-architecture/app/interfaces/api"
	"go-kafka-clean-architecture/app/interfaces/broker"
	"go-kafka-clean-architecture/app/interfaces/controller/rest_context"
	"go-kafka-clean-architecture/app/interfaces/database"
)

type Registry struct {
}

func NewRegistry() *Registry {
	return &Registry{}
}

func (r *Registry) NewSqlBrokerAppController(restAPI api.RestAPI, sqlHandler database.SQLHandler, productWriter broker.EventWriter, productReader broker.EventReader) *rest_context.AppController {
	return &rest_context.AppController{
		//ProductController: r.NewSqlBrokerBrasilProductController(restAPI, sqlHandler, productWriter, productReader),
		ProductController: r.NewSqlBrokerChileProductController(restAPI, sqlHandler, productWriter, productReader),
	}
}

func (r *Registry) NewGormBrokerAppController(restAPI api.RestAPI, sqlGorm database.SQLGorm, productWriter broker.EventWriter, productReader broker.EventReader) *rest_context.AppController {
	return &rest_context.AppController{
		//ProductController: r.NewGormBrokerBrasilProductController(restAPI, sqlGorm, productWriter, productReader),
		ProductController: r.NewGormBrokerChileProductController(restAPI, sqlGorm, productWriter, productReader),
	}
}
