package json_context

import (
	"net/http"

	"github.com/BooleanCat/go-functional/iter"
	"github.com/go-errors/errors"

	"go-kafka-clean-architecture/app_func/command/controller/json_context/model"
	"go-kafka-clean-architecture/app_func/command/router"
	"go-kafka-clean-architecture/app_func/command/usecases"
	"go-kafka-clean-architecture/app_func/entities"
	"go-kafka-clean-architecture/app_func/infrastructure/command/json_context"
)

func FindAllProductsTranslated() func(findAll usecases.ProductTranslatedInteractorFindAll) json_context.ProductTranslatedControllerFindAll {
	return func(findAll usecases.ProductTranslatedInteractorFindAll) json_context.ProductTranslatedControllerFindAll {
		return func(json router.JsonContextJSON) error {
			products, err := findAll()
			if !errors.Is(err, nil) {
				return ReturnError(json, err)
			}

			productsIteractor := iter.Lift(products)
			productsMapper := iter.Map[entities.Product](productsIteractor, model.MapProduct)
			modelProducts := iter.Collect[model.Product](productsMapper)

			return ReturnSuccess(json, http.StatusOK, modelProducts)
		}
	}
}
