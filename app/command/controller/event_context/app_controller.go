package event_context

import "go-kafka-clean-architecture/app/infrastructure/command/event_context"

type AppController struct {
	ProductTranslatedController event_context.ProductTranslatedController
}

func NewAppController(productTranslatedController event_context.ProductTranslatedController) *AppController {
	return &AppController{productTranslatedController}
}
