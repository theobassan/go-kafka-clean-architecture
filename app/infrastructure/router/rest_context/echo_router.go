package rest_context

import (
	"fmt"
	"go-kafka-clean-architecture/app/interfaces/controller/rest_context"
	"log"
	"strconv"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type echoHandler struct {
	appController *rest_context.AppController
}

func NewEchoRouter(appController *rest_context.AppController, port int) {

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

func (handler *echoHandler) create(context echo.Context) error {

	err := handler.appController.ProductController.Create(context)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}

func (handler *echoHandler) findAll(context echo.Context) error {

	err := handler.appController.ProductController.FindAll(context)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}

func (handler *echoHandler) get(context echo.Context) error {

	err := handler.appController.ProductController.Get(context)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}
