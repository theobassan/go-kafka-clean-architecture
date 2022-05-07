package json_context

import "go-kafka-clean-architecture/app/command/router"

type ProductTranslatedController interface {
	FindAll(context router.JsonContext) error
}
