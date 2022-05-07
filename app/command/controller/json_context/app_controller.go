package json_context

import "go-kafka-clean-architecture/app/infrastructure/command/json_context"

type AppController struct {
	ProductController           json_context.ProductController
	ProductTranslatedController json_context.ProductTranslatedController
}

func NewAppController(productController json_context.ProductController, productTranslatedController json_context.ProductTranslatedController) *AppController {
	return &AppController{productController, productTranslatedController}
}
