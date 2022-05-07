package http_context

import (
	"go-kafka-clean-architecture/app_func/infrastructure/command/http_context"
	"go-kafka-clean-architecture/app_func/logger"
	"strconv"

	"github.com/go-errors/errors"

	"github.com/gin-gonic/gin"
)

func StartGinRouter(port int) func(logError logger.LoggerError) func(createFunc http_context.ProductControllerCreate) func(findAllFunc http_context.ProductControllerFindAll) func(getFunc http_context.ProductControllerGet) func(findAllTranslatedFunc http_context.ProductTranslatedControllerFindAll) error {
	return func(logError logger.LoggerError) func(createFunc http_context.ProductControllerCreate) func(findAllFunc http_context.ProductControllerFindAll) func(getFunc http_context.ProductControllerGet) func(findAllTranslatedFunc http_context.ProductTranslatedControllerFindAll) error {
		return func(createFunc http_context.ProductControllerCreate) func(findAllFunc http_context.ProductControllerFindAll) func(getFunc http_context.ProductControllerGet) func(findAllTranslatedFunc http_context.ProductTranslatedControllerFindAll) error {
			return func(findAllFunc http_context.ProductControllerFindAll) func(getFunc http_context.ProductControllerGet) func(findAllTranslatedFunc http_context.ProductTranslatedControllerFindAll) error {
				return func(getFunc http_context.ProductControllerGet) func(findAllTranslatedFunc http_context.ProductTranslatedControllerFindAll) error {
					return func(findAllTranslatedFunc http_context.ProductTranslatedControllerFindAll) error {
						ginRouter := gin.Default()

						err := startGin(ginRouter, port)(logError)(createFunc)(findAllFunc)(getFunc)(findAllTranslatedFunc)
						if !errors.Is(err, nil) {
							return errors.Wrap(err, 1)
						}
						return nil
					}
				}
			}
		}
	}
}

func startGin(ginRouter *gin.Engine, port int) func(logError logger.LoggerError) func(createFunc http_context.ProductControllerCreate) func(findAllFunc http_context.ProductControllerFindAll) func(getFunc http_context.ProductControllerGet) func(findAllTranslatedFunc http_context.ProductTranslatedControllerFindAll) error {
	return func(logError logger.LoggerError) func(createFunc http_context.ProductControllerCreate) func(findAllFunc http_context.ProductControllerFindAll) func(getFunc http_context.ProductControllerGet) func(findAllTranslatedFunc http_context.ProductTranslatedControllerFindAll) error {
		return func(createFunc http_context.ProductControllerCreate) func(findAllFunc http_context.ProductControllerFindAll) func(getFunc http_context.ProductControllerGet) func(findAllTranslatedFunc http_context.ProductTranslatedControllerFindAll) error {
			return func(findAllFunc http_context.ProductControllerFindAll) func(getFunc http_context.ProductControllerGet) func(findAllTranslatedFunc http_context.ProductTranslatedControllerFindAll) error {
				return func(getFunc http_context.ProductControllerGet) func(findAllTranslatedFunc http_context.ProductTranslatedControllerFindAll) error {
					return func(findAllTranslatedFunc http_context.ProductTranslatedControllerFindAll) error {
						ginRouter.POST("/product", func(ginContext *gin.Context) { createGin(ginContext)(logError)(createFunc) })
						ginRouter.GET("/products", func(ginContext *gin.Context) { findAllGin(ginContext)(logError)(findAllFunc) })
						ginRouter.GET("/product", func(ginContext *gin.Context) { getGin(ginContext)(logError)(getFunc) })

						ginRouter.GET("/productstranslated", func(ginContext *gin.Context) {
							findAllTranslatedGin(ginContext)(logError)(findAllTranslatedFunc)
						})

						err := ginRouter.Run(":" + strconv.Itoa(port))
						if !errors.Is(err, nil) {
							return errors.Wrap(err, 1)
						}
						return nil
					}
				}
			}
		}
	}
}

func createGin(ginContext *gin.Context) func(logError logger.LoggerError) func(create http_context.ProductControllerCreate) error {
	return func(logError logger.LoggerError) func(create http_context.ProductControllerCreate) error {
		return func(create http_context.ProductControllerCreate) error {
			context := GinContext{ginContext}

			err := create(context.ResponseWriter, context.Request)
			if !errors.Is(err, nil) {
				logError(err)
			}
			return nil
		}
	}
}

func findAllGin(ginContext *gin.Context) func(logError logger.LoggerError) func(findAll http_context.ProductControllerFindAll) error {
	return func(logError logger.LoggerError) func(findAll http_context.ProductControllerFindAll) error {
		return func(findAll http_context.ProductControllerFindAll) error {
			context := GinContext{ginContext}

			err := findAll(context.ResponseWriter)
			if !errors.Is(err, nil) {
				logError(err)
			}
			return nil
		}
	}
}

func getGin(ginContext *gin.Context) func(logError logger.LoggerError) func(get http_context.ProductControllerGet) error {
	return func(logError logger.LoggerError) func(get http_context.ProductControllerGet) error {
		return func(get http_context.ProductControllerGet) error {
			context := GinContext{ginContext}

			err := get(context.ResponseWriter, context.Request)
			if !errors.Is(err, nil) {
				logError(err)
			}
			return nil
		}
	}
}

func findAllTranslatedGin(ginContext *gin.Context) func(logError logger.LoggerError) func(findAll http_context.ProductTranslatedControllerFindAll) error {
	return func(logError logger.LoggerError) func(findAll http_context.ProductTranslatedControllerFindAll) error {
		return func(findAll http_context.ProductTranslatedControllerFindAll) error {
			context := GinContext{ginContext}

			err := findAll(context.ResponseWriter)
			if !errors.Is(err, nil) {
				logError(err)
			}
			return nil
		}
	}
}
