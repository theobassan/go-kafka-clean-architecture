package registry

import (
	"go-kafka-clean-architecture/app/infrastructure/api/event_api"
	"go-kafka-clean-architecture/app/infrastructure/api/rest_api"
	"go-kafka-clean-architecture/app/interfaces/api"
)

func (r *Registry) NewEventApi() api.EventApi {
	eventApi := event_api.NewKafkaApi("localhost:9092")

	return eventApi
}

func (r *Registry) NewRestApi() api.RestApi {
	return rest_api.NewHttpApi("http://localhost:8080")
}
