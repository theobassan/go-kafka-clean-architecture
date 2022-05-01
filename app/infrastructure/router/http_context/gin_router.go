package http_context

import (
	"errors"
	"go-kafka-clean-architecture/app/input/controller/http_context"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ginHandler struct {
	appController *http_context.AppController
}

func StartGinRouter(appController *http_context.AppController, port int) {
	handler := &ginHandler{
		appController: appController,
	}

	ginRouter := gin.Default()
	handler.start(ginRouter, port)
}

func (handler *ginHandler) start(ginRouter *gin.Engine, port int) {
	ginRouter.GET("/products", handler.findAll)
	ginRouter.POST("/product", handler.create)
	ginRouter.GET("/product", handler.get)

	ginRouter.GET("/productstranslated", handler.findAllTranslated)

	err := ginRouter.Run(":" + strconv.Itoa(port))
	if !errors.Is(err, nil) {
		log.Fatalln(err)
	}
}

func (handler *ginHandler) create(ginContext *gin.Context) {
	context := &GinContext{ginContext}

	handler.appController.ProductController.Create(context)
}

func (handler *ginHandler) findAll(ginContext *gin.Context) {
	context := &GinContext{ginContext}

	handler.appController.ProductController.FindAll(context)
}

func (handler *ginHandler) get(ginContext *gin.Context) {
	context := &GinContext{ginContext}

	handler.appController.ProductController.Get(context)
}

func (handler *ginHandler) findAllTranslated(ginContext *gin.Context) {
	context := &GinContext{ginContext}

	handler.appController.ProductTranslatedController.FindAll(context)
}
