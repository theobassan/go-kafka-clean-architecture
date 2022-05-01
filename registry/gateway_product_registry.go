package registry

import (
	"go-kafka-clean-architecture/app/interfaces/api"
	"go-kafka-clean-architecture/app/interfaces/gateway/event_api"
	"go-kafka-clean-architecture/app/interfaces/gateway/rest_api"
	"go-kafka-clean-architecture/app/usecases/gateway"
)

func (r *Registry) NewProductFinderGateway(restAPI api.RestAPI) gateway.ProductFinderGateway {
	return rest_api.NewProductFinderGateway(restAPI)
}

func (r *Registry) NewProductSenderGateway(eventAPI api.EventAPI) gateway.ProductSenderGateway {
	return event_api.NewProductSenderGateway(eventAPI)
}
