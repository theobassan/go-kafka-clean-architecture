package registry

import (
	"go-kafka-clean-architecture/app_func/interfaces/api"
	"go-kafka-clean-architecture/app_func/interfaces/gateway/event_api"
	"go-kafka-clean-architecture/app_func/interfaces/gateway/rest_api"
	"go-kafka-clean-architecture/app_func/usecases/gateway"
)

func RestApiProductFinderGatewayFindById(get api.RestApiGet) gateway.ProductFinderGatewayFindById {
	return rest_api.FindProductById()(get)
}

func EventApiProductSenderGatewaySend(bind api.EventApiBind, writeMessage api.EventApiWriteMessage) gateway.ProductSenderGatewaySend {
	return event_api.SendProduct()(bind)(writeMessage)
}
