package http_context

import "go-kafka-clean-architecture/app/infrastructure/command/http_context"

type AppController struct {
	ProductController           http_context.ProductController
	ProductTranslatedController http_context.ProductTranslatedController
}

func NewAppController(productController http_context.ProductController, productTranslatedController http_context.ProductTranslatedController) *AppController {
	return &AppController{productController, productTranslatedController}
}
