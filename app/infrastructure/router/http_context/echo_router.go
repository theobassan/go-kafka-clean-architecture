package http_context

import (
	"errors"
	"go-kafka-clean-architecture/app/input/controller/http_context"
	"log"
	"strconv"

	"github.com/labstack/echo"
)

type echoHandler struct {
	appController *http_context.AppController
}

func StartEchoRouter(appController *http_context.AppController, port int) {

	handler := &echoHandler{
		appController: appController,
	}

	echoRouter := echo.New()
	//echoRouter.Use(middleware.Logger())
	//echoRouter.Use(middleware.Recover())

	handler.start(echoRouter, port)
}

func (handler *echoHandler) start(echoRouter *echo.Echo, port int) {
	echoRouter.GET("/products", handler.findAll)
	echoRouter.POST("/product", handler.create)
	echoRouter.GET("/product", handler.get)

	echoRouter.GET("/productstranslated", handler.findAllTranslated)

	err := echoRouter.Start(":" + strconv.Itoa(port))
	if !errors.Is(err, nil) {
		log.Fatalln(err)
	}
}

func (handler *echoHandler) create(echoContext echo.Context) error {
	context := &EchoContext{echoContext}

	handler.appController.ProductController.Create(context)
	return nil
}

func (handler *echoHandler) findAll(echoContext echo.Context) error {
	context := &EchoContext{echoContext}

	handler.appController.ProductController.FindAll(context)

	return nil
}

func (handler *echoHandler) get(echoContext echo.Context) error {
	context := &EchoContext{echoContext}

	handler.appController.ProductController.Get(context)
	return nil
}

func (handler *echoHandler) findAllTranslated(echoContext echo.Context) error {
	context := &EchoContext{echoContext}

	handler.appController.ProductTranslatedController.FindAll(context)

	return nil
}
