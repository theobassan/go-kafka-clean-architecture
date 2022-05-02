package json_context

import (
	"go-kafka-clean-architecture/app/command/controller/json_context"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ginHandler struct {
	appController *json_context.AppController
}

func StarGinRouter(appController *json_context.AppController, port int) {
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

func (handler *ginHandler) create(ginContext *gin.Context) {
	context := GinContext{ginContext}

	handler.appController.ProductController.Create(context)
}

func (handler *ginHandler) findAll(ginContext *gin.Context) {
	context := GinContext{ginContext}

	handler.appController.ProductController.FindAll(context)
}

func (handler *ginHandler) get(ginContext *gin.Context) {
	context := GinContext{ginContext}

	handler.appController.ProductController.Get(context)
}
