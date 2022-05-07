package registry

import (
	"go-kafka-clean-architecture/app_func/infrastructure/api/event_api"
	"go-kafka-clean-architecture/app_func/infrastructure/api/rest_api"
	"go-kafka-clean-architecture/app_func/interfaces/api"
)

func EventApiWriteMessage() (api.EventApiBind, api.EventApiWriteMessage) {
	bind, writeMessage := event_api.NewKafkaApi("localhost:9092")

	return bind, writeMessage
}

func RestApiGet() api.RestApiGet {
	return rest_api.Get("http://localhost:8080")
}
