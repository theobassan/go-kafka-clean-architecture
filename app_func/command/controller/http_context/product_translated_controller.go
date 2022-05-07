package http_context

import (
	"net/http"

	"github.com/BooleanCat/go-functional/iter"
	"github.com/go-errors/errors"

	"go-kafka-clean-architecture/app_func/command/controller/http_context/model"
	"go-kafka-clean-architecture/app_func/command/router"
	"go-kafka-clean-architecture/app_func/command/usecases"
	"go-kafka-clean-architecture/app_func/entities"
	"go-kafka-clean-architecture/app_func/infrastructure/command/http_context"
)

func FindAllProductsTranslated() func(findAll usecases.ProductTranslatedInteractorFindAll) http_context.ProductTranslatedControllerFindAll {
	return func(findAll usecases.ProductTranslatedInteractorFindAll) http_context.ProductTranslatedControllerFindAll {
		return func(responseWriterFunc router.HttpContextResponseWriter) error {
			responseWriter := responseWriterFunc()

			products, err := findAll()
			if !errors.Is(err, nil) {
				return ReturnError(responseWriter, err)
			}

			productsIteractor := iter.Lift(products)
			productsMapper := iter.Map[entities.Product](productsIteractor, model.MapProduct)
			modelProducts := iter.Collect[model.Product](productsMapper)

			return ReturnSuccess(responseWriter, http.StatusOK, modelProducts)
		}
	}
}
