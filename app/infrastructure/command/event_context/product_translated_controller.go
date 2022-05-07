package event_context

import "go-kafka-clean-architecture/app/command/router"

type ProductTranslatedController interface {
	Create(context router.EventContext) error
}
