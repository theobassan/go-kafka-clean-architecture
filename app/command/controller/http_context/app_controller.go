package http_context

type AppController struct {
	ProductController           ProductController
	ProductTranslatedController ProductTranslatedController
}

func NewAppController(ProductController ProductController, ProductTranslatedController ProductTranslatedController) *AppController {
	return &AppController{ProductController, ProductTranslatedController}
}
