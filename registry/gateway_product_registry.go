package registry

import (
	"go-kafka-clean-architecture/app/interfaces/api"
	"go-kafka-clean-architecture/app/interfaces/gateway/rest_api"
	"go-kafka-clean-architecture/app/usecases/gateway"
)

func (r *Registry) NewProductGateway(restAPI api.RestAPI) gateway.ProductGateway {
	return rest_api.NewProductGateway(restAPI)
}
