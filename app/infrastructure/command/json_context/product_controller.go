package json_context

import "go-kafka-clean-architecture/app/command/router"

type ProductController interface {
	Create(context router.JsonContext) error
	FindAll(context router.JsonContext) error
	Get(context router.JsonContext) error
}
