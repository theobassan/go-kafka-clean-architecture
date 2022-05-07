package http_context

import "go-kafka-clean-architecture/app/command/router"

type ProductController interface {
	Create(context router.HttpContext) error
	FindAll(context router.HttpContext) error
	Get(context router.HttpContext) error
}
