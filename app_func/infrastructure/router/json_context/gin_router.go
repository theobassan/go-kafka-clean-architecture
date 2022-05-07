package json_context

import (
	"go-kafka-clean-architecture/app_func/infrastructure/command/json_context"
	"go-kafka-clean-architecture/app_func/logger"
	"strconv"

	"github.com/go-errors/errors"

	"github.com/gin-gonic/gin"
)

func StartGinRouter(port int) func(logError logger.LoggerError) func(createFunc json_context.ProductControllerCreate) func(findAllFunc json_context.ProductControllerFindAll) func(getFunc json_context.ProductControllerGet) func(findAllTranslatedFunc json_context.ProductTranslatedControllerFindAll) error {
	return func(logError logger.LoggerError) func(createFunc json_context.ProductControllerCreate) func(findAllFunc json_context.ProductControllerFindAll) func(getFunc json_context.ProductControllerGet) func(findAllTranslatedFunc json_context.ProductTranslatedControllerFindAll) error {
		return func(createFunc json_context.ProductControllerCreate) func(findAllFunc json_context.ProductControllerFindAll) func(getFunc json_context.ProductControllerGet) func(findAllTranslatedFunc json_context.ProductTranslatedControllerFindAll) error {
			return func(findAllFunc json_context.ProductControllerFindAll) func(getFunc json_context.ProductControllerGet) func(findAllTranslatedFunc json_context.ProductTranslatedControllerFindAll) error {
				return func(getFunc json_context.ProductControllerGet) func(findAllTranslatedFunc json_context.ProductTranslatedControllerFindAll) error {
					return func(findAllTranslatedFunc json_context.ProductTranslatedControllerFindAll) error {
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

func startGin(ginRouter *gin.Engine, port int) func(logError logger.LoggerError) func(createFunc json_context.ProductControllerCreate) func(findAllFunc json_context.ProductControllerFindAll) func(getFunc json_context.ProductControllerGet) func(findAllTranslatedFunc json_context.ProductTranslatedControllerFindAll) error {
	return func(logError logger.LoggerError) func(createFunc json_context.ProductControllerCreate) func(findAllFunc json_context.ProductControllerFindAll) func(getFunc json_context.ProductControllerGet) func(findAllTranslatedFunc json_context.ProductTranslatedControllerFindAll) error {
		return func(createFunc json_context.ProductControllerCreate) func(findAllFunc json_context.ProductControllerFindAll) func(getFunc json_context.ProductControllerGet) func(findAllTranslatedFunc json_context.ProductTranslatedControllerFindAll) error {
			return func(findAllFunc json_context.ProductControllerFindAll) func(getFunc json_context.ProductControllerGet) func(findAllTranslatedFunc json_context.ProductTranslatedControllerFindAll) error {
				return func(getFunc json_context.ProductControllerGet) func(findAllTranslatedFunc json_context.ProductTranslatedControllerFindAll) error {
					return func(findAllTranslatedFunc json_context.ProductTranslatedControllerFindAll) error {
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

func createGin(ginContext *gin.Context) func(logError logger.LoggerError) func(create json_context.ProductControllerCreate) error {
	return func(logError logger.LoggerError) func(create json_context.ProductControllerCreate) error {
		return func(create json_context.ProductControllerCreate) error {
			context := GinContext{ginContext}

			err := create(context.Bind, context.JSON)
			if !errors.Is(err, nil) {
				logError(err)
			}
			return nil
		}
	}
}

func findAllGin(ginContext *gin.Context) func(logError logger.LoggerError) func(findAll json_context.ProductControllerFindAll) error {
	return func(logError logger.LoggerError) func(findAll json_context.ProductControllerFindAll) error {
		return func(findAll json_context.ProductControllerFindAll) error {
			context := GinContext{ginContext}

			err := findAll(context.JSON)
			if !errors.Is(err, nil) {
				logError(err)
			}
			return nil
		}
	}
}

func getGin(ginContext *gin.Context) func(logError logger.LoggerError) func(get json_context.ProductControllerGet) error {
	return func(logError logger.LoggerError) func(get json_context.ProductControllerGet) error {
		return func(get json_context.ProductControllerGet) error {
			context := GinContext{ginContext}

			err := get(context.Query, context.JSON)
			if !errors.Is(err, nil) {
				logError(err)
			}
			return nil
		}
	}
}

func findAllTranslatedGin(ginContext *gin.Context) func(logError logger.LoggerError) func(findAll json_context.ProductTranslatedControllerFindAll) error {
	return func(logError logger.LoggerError) func(findAll json_context.ProductTranslatedControllerFindAll) error {
		return func(findAll json_context.ProductTranslatedControllerFindAll) error {
			context := GinContext{ginContext}

			err := findAll(context.JSON)
			if !errors.Is(err, nil) {
				logError(err)
			}
			return nil
		}
	}
}
