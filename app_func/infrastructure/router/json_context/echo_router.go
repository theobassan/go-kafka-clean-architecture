package json_context

import (
	"go-kafka-clean-architecture/app_func/infrastructure/command/json_context"
	"go-kafka-clean-architecture/app_func/logger"
	"strconv"

	"github.com/go-errors/errors"
	"github.com/labstack/echo"
)

func StartEchoRouter(port int) func(logError logger.LoggerError) func(createFunc json_context.ProductControllerCreate) func(findAllFunc json_context.ProductControllerFindAll) func(getFunc json_context.ProductControllerGet) func(findAllTranslatedFunc json_context.ProductTranslatedControllerFindAll) error {
	return func(logError logger.LoggerError) func(createFunc json_context.ProductControllerCreate) func(findAllFunc json_context.ProductControllerFindAll) func(getFunc json_context.ProductControllerGet) func(findAllTranslatedFunc json_context.ProductTranslatedControllerFindAll) error {
		return func(createFunc json_context.ProductControllerCreate) func(findAllFunc json_context.ProductControllerFindAll) func(getFunc json_context.ProductControllerGet) func(findAllTranslatedFunc json_context.ProductTranslatedControllerFindAll) error {
			return func(findAllFunc json_context.ProductControllerFindAll) func(getFunc json_context.ProductControllerGet) func(findAllTranslatedFunc json_context.ProductTranslatedControllerFindAll) error {
				return func(getFunc json_context.ProductControllerGet) func(findAllTranslatedFunc json_context.ProductTranslatedControllerFindAll) error {
					return func(findAllTranslatedFunc json_context.ProductTranslatedControllerFindAll) error {
						echoRouter := echo.New()
						//echoRouter.Use(middleware.Logger())
						//echoRouter.Use(middleware.Recover())

						err := startEcho(echoRouter, port)(logError)(createFunc)(findAllFunc)(getFunc)(findAllTranslatedFunc)
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

func startEcho(echoRouter *echo.Echo, port int) func(logError logger.LoggerError) func(createFunc json_context.ProductControllerCreate) func(findAllFunc json_context.ProductControllerFindAll) func(getFunc json_context.ProductControllerGet) func(findAllTranslatedFunc json_context.ProductTranslatedControllerFindAll) error {
	return func(logError logger.LoggerError) func(createFunc json_context.ProductControllerCreate) func(findAllFunc json_context.ProductControllerFindAll) func(getFunc json_context.ProductControllerGet) func(findAllTranslatedFunc json_context.ProductTranslatedControllerFindAll) error {
		return func(createFunc json_context.ProductControllerCreate) func(findAllFunc json_context.ProductControllerFindAll) func(getFunc json_context.ProductControllerGet) func(findAllTranslatedFunc json_context.ProductTranslatedControllerFindAll) error {
			return func(findAllFunc json_context.ProductControllerFindAll) func(getFunc json_context.ProductControllerGet) func(findAllTranslatedFunc json_context.ProductTranslatedControllerFindAll) error {
				return func(getFunc json_context.ProductControllerGet) func(findAllTranslatedFunc json_context.ProductTranslatedControllerFindAll) error {
					return func(findAllTranslatedFunc json_context.ProductTranslatedControllerFindAll) error {
						echoRouter.POST("/product", func(echoContext echo.Context) error { return createEcho(echoContext)(logError)(createFunc) })
						echoRouter.GET("/products", func(echoContext echo.Context) error { return findAllEcho(echoContext)(logError)(findAllFunc) })
						echoRouter.GET("/product", func(echoContext echo.Context) error { return getEcho(echoContext)(logError)(getFunc) })

						echoRouter.GET("/productstranslated", func(echoContext echo.Context) error {
							return findAllTranslatedEcho(echoContext)(logError)(findAllTranslatedFunc)
						})

						err := echoRouter.Start(":" + strconv.Itoa(port))
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

func createEcho(echoContext echo.Context) func(logError logger.LoggerError) func(create json_context.ProductControllerCreate) error {
	return func(logError logger.LoggerError) func(create json_context.ProductControllerCreate) error {
		return func(create json_context.ProductControllerCreate) error {
			context := EchoContext{echoContext}

			err := create(context.Bind, context.JSON)
			if !errors.Is(err, nil) {
				logError(err)
			}
			return nil
		}
	}
}

func findAllEcho(echoContext echo.Context) func(logError logger.LoggerError) func(findAll json_context.ProductControllerFindAll) error {
	return func(logError logger.LoggerError) func(findAll json_context.ProductControllerFindAll) error {
		return func(findAll json_context.ProductControllerFindAll) error {
			context := EchoContext{echoContext}

			err := findAll(context.JSON)
			if !errors.Is(err, nil) {
				logError(err)
			}
			return nil
		}
	}
}

func getEcho(echoContext echo.Context) func(logError logger.LoggerError) func(get json_context.ProductControllerGet) error {
	return func(logError logger.LoggerError) func(get json_context.ProductControllerGet) error {
		return func(get json_context.ProductControllerGet) error {
			context := EchoContext{echoContext}

			err := get(context.Query, context.JSON)
			if !errors.Is(err, nil) {
				logError(err)
			}
			return nil
		}
	}
}

func findAllTranslatedEcho(echoContext echo.Context) func(logError logger.LoggerError) func(findAll json_context.ProductTranslatedControllerFindAll) error {
	return func(logError logger.LoggerError) func(findAll json_context.ProductTranslatedControllerFindAll) error {
		return func(findAll json_context.ProductTranslatedControllerFindAll) error {
			context := EchoContext{echoContext}

			err := findAll(context.JSON)
			if !errors.Is(err, nil) {
				logError(err)
			}
			return nil
		}
	}
}
