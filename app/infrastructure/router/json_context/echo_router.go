package json_context

import (
	"go-kafka-clean-architecture/app/input/controller/json_context"
	"log"
	"strconv"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type echoHandler struct {
	appController *json_context.AppController
}

func StartEchoRouter(appController *json_context.AppController, port int) {

	handler := &echoHandler{
		appController: appController,
	}

	echoRouter := echo.New()

	echoRouter.Use(middleware.Logger())
	echoRouter.Use(middleware.Recover())

	echoRouter.GET("/products", handler.findAll)
	echoRouter.POST("/product", handler.create)
	echoRouter.GET("/product/:id", handler.get)

	if err := echoRouter.Start(":" + strconv.Itoa(port)); err != nil {
		log.Fatalln(err)
	}
}

func (handler *echoHandler) create(echoContext echo.Context) error {
	context := EchoContext{echoContext}

	handler.appController.ProductController.Create(context)
	return nil
}

func (handler *echoHandler) findAll(echoContext echo.Context) error {
	context := EchoContext{echoContext}

	handler.appController.ProductController.FindAll(context)
	return nil
}

func (handler *echoHandler) get(echoContext echo.Context) error {
	context := EchoContext{echoContext}

	handler.appController.ProductController.Get(context)
	return nil
}
