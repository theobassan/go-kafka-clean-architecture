package http_context

import (
	"go-kafka-clean-architecture/app/command/controller/http_context"
	"go-kafka-clean-architecture/app/logger"
	"strconv"

	"github.com/go-errors/errors"
	"github.com/labstack/echo"
)

type echoHandler struct {
	appController *http_context.AppController
	logger        logger.Logger
}

func StartEchoRouter(appController *http_context.AppController, port int, logger logger.Logger) error {
	handler := &echoHandler{
		appController: appController,
		logger:        logger,
	}

	echoRouter := echo.New()
	//echoRouter.Use(middleware.Logger())
	//echoRouter.Use(middleware.Recover())

	err := handler.start(echoRouter, port)
	if !errors.Is(err, nil) {
		return errors.Wrap(err, 1)
	}
	return nil
}

func (handler *echoHandler) start(echoRouter *echo.Echo, port int) error {
	echoRouter.GET("/products", handler.findAll)
	echoRouter.POST("/product", handler.create)
	echoRouter.GET("/product", handler.get)

	echoRouter.GET("/productstranslated", handler.findAllTranslated)

	err := echoRouter.Start(":" + strconv.Itoa(port))
	if !errors.Is(err, nil) {
		return errors.Wrap(err, 1)
	}
	return nil
}

func (handler *echoHandler) create(echoContext echo.Context) error {
	context := &EchoContext{echoContext}

	err := handler.appController.ProductController.Create(context)
	if !errors.Is(err, nil) {
		handler.logger.Error(err)
	}

	return nil
}

func (handler *echoHandler) findAll(echoContext echo.Context) error {
	context := &EchoContext{echoContext}

	err := handler.appController.ProductController.FindAll(context)
	if !errors.Is(err, nil) {
		handler.logger.Error(err)
	}

	return nil
}

func (handler *echoHandler) get(echoContext echo.Context) error {
	context := &EchoContext{echoContext}

	err := handler.appController.ProductController.Get(context)
	if !errors.Is(err, nil) {
		handler.logger.Error(err)
	}
	return nil
}

func (handler *echoHandler) findAllTranslated(echoContext echo.Context) error {
	context := &EchoContext{echoContext}

	err := handler.appController.ProductTranslatedController.FindAll(context)
	if !errors.Is(err, nil) {
		handler.logger.Error(err)
	}

	return nil
}
