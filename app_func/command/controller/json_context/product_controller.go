package json_context

import (
	"net/http"
	"strconv"

	"github.com/BooleanCat/go-functional/iter"
	"github.com/go-errors/errors"

	"go-kafka-clean-architecture/app_func/command/controller/json_context/model"
	"go-kafka-clean-architecture/app_func/command/router"
	"go-kafka-clean-architecture/app_func/command/usecases"
	"go-kafka-clean-architecture/app_func/entities"
	"go-kafka-clean-architecture/app_func/infrastructure/command/json_context"
)

func CreateProduct() func(create usecases.ProductInteractorCreate) json_context.ProductControllerCreate {
	return func(create usecases.ProductInteractorCreate) json_context.ProductControllerCreate {
		return func(bind router.JsonContextBind, json router.JsonContextJSON) error {
			var params model.Product

			err := bind(&params)
			if !errors.Is(err, nil) {
				return ReturnError(json, err)
			}

			id, err := create(params.ID)
			if !errors.Is(err, nil) {
				return ReturnError(json, err)
			}

			return ReturnSuccess(json, http.StatusCreated, id)
		}
	}
}

func FindAllProducts() func(findAll usecases.ProductInteractorFindAll) json_context.ProductControllerFindAll {
	return func(findAll usecases.ProductInteractorFindAll) json_context.ProductControllerFindAll {
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

func GetProduct() func(get usecases.ProductInteractorGet) json_context.ProductControllerGet {
	return func(get usecases.ProductInteractorGet) json_context.ProductControllerGet {
		return func(query router.JsonContextQuery, json router.JsonContextJSON) error {
			id := query("id")
			productID, err := strconv.ParseInt(id, 10, 64)
			if !errors.Is(err, nil) {
				return ReturnError(json, err)
			}
			product, err := get(productID)
			if !errors.Is(err, nil) {
				return ReturnError(json, err)
			}

			modelProduct := model.MapProduct(product)

			return ReturnSuccess(json, http.StatusOK, modelProduct)
		}
	}
}
