package registry

import (
	"context"
	"go-kafka-clean-architecture/app/interfaces/api"
	"go-kafka-clean-architecture/app/interfaces/broker"
	"go-kafka-clean-architecture/app/interfaces/controller/rest_context"
	"go-kafka-clean-architecture/app/interfaces/database"
	"go-kafka-clean-architecture/app/interfaces/repository/sql_gorm"
	"go-kafka-clean-architecture/app/interfaces/repository/sql_handler"
	"go-kafka-clean-architecture/app/usecases/interactor"
	"go-kafka-clean-architecture/app/usecases/repository"
)

func (r *Registry) NewSqlBrokerBrasilProductController(restAPI api.RestAPI, sqlHandler database.SQLHandler, productWriter broker.EventWriter, productReader broker.EventReader) rest_context.ProductController {
	interactor := r.NewSqlBrokerProductInteractor(restAPI, sqlHandler, productWriter, r.NewProductBrasilTranslatorInteractor())

	subscriber := r.NewBrokerProductSubscriber(productReader, interactor)
	go subscriber.Subscribe(context.Background())

	return rest_context.NewProductController(interactor)
}

func (r *Registry) NewSqlBrokerChileProductController(restAPI api.RestAPI, sqlHandler database.SQLHandler, productWriter broker.EventWriter, productReader broker.EventReader) rest_context.ProductController {
	interactor := r.NewSqlBrokerProductInteractor(restAPI, sqlHandler, productWriter, r.NewProductChileTranslatorInteractor())

	subscriber := r.NewBrokerProductSubscriber(productReader, interactor)
	go subscriber.Subscribe(context.Background())

	return rest_context.NewProductController(interactor)
}

func (r *Registry) NewSqlBrokerProductInteractor(restAPI api.RestAPI, sqlHandler database.SQLHandler, productWriter broker.EventWriter, translator interactor.ProductTranslatorInteractor) interactor.ProductInteractor {
	return interactor.NewProductInteractor(r.NewProductGateway(restAPI), r.NewSqlProductRepository(sqlHandler), r.NewBrokerProductPublisher(productWriter), translator)
}

func (r *Registry) NewSqlProductRepository(sqlHandler database.SQLHandler) repository.ProductRepository {
	return sql_handler.NewProductRepository(sqlHandler)
}

func (r *Registry) NewGormBrokerBrasilProductController(restAPI api.RestAPI, sqlGorm database.SQLGorm, productWriter broker.EventWriter, productReader broker.EventReader) rest_context.ProductController {
	interactor := r.NewGormBrokerProductInteractor(restAPI, sqlGorm, productWriter, r.NewProductBrasilTranslatorInteractor())

	subscriber := r.NewBrokerProductSubscriber(productReader, interactor)
	go subscriber.Subscribe(context.Background())

	return rest_context.NewProductController(interactor)
}

func (r *Registry) NewGormBrokerChileProductController(restAPI api.RestAPI, sqlGorm database.SQLGorm, productWriter broker.EventWriter, productReader broker.EventReader) rest_context.ProductController {
	interactor := r.NewGormBrokerProductInteractor(restAPI, sqlGorm, productWriter, r.NewProductChileTranslatorInteractor())

	subscriber := r.NewBrokerProductSubscriber(productReader, interactor)
	go subscriber.Subscribe(context.Background())

	return rest_context.NewProductController(interactor)
}

func (r *Registry) NewGormBrokerProductInteractor(restAPI api.RestAPI, sqlGorm database.SQLGorm, productWriter broker.EventWriter, translator interactor.ProductTranslatorInteractor) interactor.ProductInteractor {
	return interactor.NewProductInteractor(r.NewProductGateway(restAPI), r.NewGormProductRepository(sqlGorm), r.NewBrokerProductPublisher(productWriter), translator)
}

func (r *Registry) NewGormProductRepository(sqlGorm database.SQLGorm) repository.ProductRepository {
	return sql_gorm.NewProductRepository(sqlGorm)
}
