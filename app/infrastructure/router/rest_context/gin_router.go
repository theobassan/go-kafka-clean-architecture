package rest_context

import (
	"fmt"
	"go-kafka-clean-architecture/app/interfaces/controller/rest_context"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ginHandler struct {
	appController *rest_context.AppController
}

func NewGinRouter(appController *rest_context.AppController, port int) {
	handler := &ginHandler{
		appController: appController,
	}

	ginRouter := gin.Default()
	ginRouter.GET("/products", handler.findAll)
	ginRouter.POST("/product", handler.create)
	ginRouter.GET("/product/:id", handler.get)

	if err := ginRouter.Run(":" + strconv.Itoa(port)); err != nil {
		log.Fatalln(err)
	}
}

type Context struct {
	ginContext *gin.Context
}

func (context Context) JSON(code int, i interface{}) error {
	context.ginContext.JSON(code, i)
	return nil
}

func (context Context) Bind(i interface{}) error {
	return context.ginContext.Bind(i)
}

func (context Context) Param(key string) string {
	return context.ginContext.Param(key)
}

func (handler *ginHandler) create(ginContext *gin.Context) {
	context := Context{ginContext}

	err := handler.appController.ProductController.Create(context)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

func (handler *ginHandler) findAll(ginContext *gin.Context) {
	context := Context{ginContext}

	err := handler.appController.ProductController.FindAll(context)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

func (handler *ginHandler) get(ginContext *gin.Context) {
	context := Context{ginContext}

	err := handler.appController.ProductController.Get(context)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
