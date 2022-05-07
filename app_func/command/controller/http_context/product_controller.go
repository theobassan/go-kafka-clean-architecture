package http_context

import (
	"encoding/json"
	"net/http"

	"github.com/BooleanCat/go-functional/iter"
	"github.com/go-errors/errors"

	"go-kafka-clean-architecture/app_func/command/controller/http_context/model"
	"go-kafka-clean-architecture/app_func/command/router"
	"go-kafka-clean-architecture/app_func/command/usecases"
	"go-kafka-clean-architecture/app_func/entities"
	"go-kafka-clean-architecture/app_func/infrastructure/command/http_context"
)

func CreateProduct() func(create usecases.ProductInteractorCreate) http_context.ProductControllerCreate {
	return func(create usecases.ProductInteractorCreate) http_context.ProductControllerCreate {
		return func(responseWriterFunc router.HttpContextResponseWriter, requestFunc router.HttpContextRequest) error {
			responseWriter := responseWriterFunc()
			request := requestFunc()

			product := model.Product{}
			err := json.NewDecoder(request.Body).Decode(&product)
			if !errors.Is(err, nil) {
				return ReturnError(responseWriter, err)
			}

			id, err := create(product.ID)
			if !errors.Is(err, nil) {
				return ReturnError(responseWriter, err)
			}

			return ReturnSuccess(responseWriter, http.StatusCreated, id)
		}
	}
}

func FindAllProducts() func(findAll usecases.ProductInteractorFindAll) http_context.ProductControllerFindAll {
	return func(findAll usecases.ProductInteractorFindAll) http_context.ProductControllerFindAll {
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

func GetProduct() func(get usecases.ProductInteractorGet) http_context.ProductControllerGet {
	return func(get usecases.ProductInteractorGet) http_context.ProductControllerGet {
		return func(responseWriterFunc router.HttpContextResponseWriter, requestFunc router.HttpContextRequest) error {
			responseWriter := responseWriterFunc()
			request := requestFunc()

			productID, err := GetIntParamFromRequest(request, "id")
			if !errors.Is(err, nil) {
				return ReturnError(responseWriter, err)
			}

			product, err := get(productID)
			if !errors.Is(err, nil) {
				return ReturnError(responseWriter, err)
			}

			modelProduct := model.MapProduct(product)

			return ReturnSuccess(responseWriter, http.StatusOK, modelProduct)
		}
	}
}
