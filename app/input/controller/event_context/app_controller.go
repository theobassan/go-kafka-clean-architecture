package event_context

type AppController struct {
	ProductTranslatedController ProductTranslatedController
}

func NewAppController(ProductTranslatedController ProductTranslatedController) *AppController {
	return &AppController{ProductTranslatedController}
}
