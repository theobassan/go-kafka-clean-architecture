package registry

import (
	"go-kafka-clean-architecture/app/interfaces/api"
	"go-kafka-clean-architecture/app/interfaces/gateway/event_api"
	"go-kafka-clean-architecture/app/interfaces/gateway/rest_api"
	"go-kafka-clean-architecture/app/usecases/gateway"
)

func (r *Registry) NewRestApiProductFinderGateway(restApi api.RestApi) gateway.ProductFinderGateway {
	return rest_api.NewProductFinderGateway(restApi)
}

func (r *Registry) NewEventApiProductSenderGateway(eventApi api.EventApi) gateway.ProductSenderGateway {
	return event_api.NewProductSenderGateway(eventApi)
}
