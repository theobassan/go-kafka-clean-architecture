package json_context

import (
	"go-kafka-clean-architecture/app/command/controller/json_context"
	"go-kafka-clean-architecture/app/logger"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-errors/errors"
)

type ginHandler struct {
	appController *json_context.AppController
	logger        logger.Logger
}

func StarGinRouter(appController *json_context.AppController, port int, logger logger.Logger) error {
	handler := &ginHandler{
		appController: appController,
		logger:        logger,
	}

	ginRouter := gin.Default()

	err := handler.start(ginRouter, port)
	if !errors.Is(err, nil) {
		return errors.Wrap(err, 1)
	}
	return nil
}

func (handler *ginHandler) start(ginRouter *gin.Engine, port int) error {
	ginRouter.GET("/products", handler.findAll)
	ginRouter.POST("/product", handler.create)
	ginRouter.GET("/product", handler.get)

	ginRouter.GET("/productstranslated", handler.findAllTranslated)

	err := ginRouter.Run(":" + strconv.Itoa(port))
	if !errors.Is(err, nil) {
		return errors.Wrap(err, 1)
	}
	return nil
}

func (handler *ginHandler) create(ginContext *gin.Context) {
	context := GinContext{ginContext}

	err := handler.appController.ProductController.Create(context)
	if !errors.Is(err, nil) {
		handler.logger.Error(err)
	}
}

func (handler *ginHandler) findAll(ginContext *gin.Context) {
	context := GinContext{ginContext}

	err := handler.appController.ProductController.FindAll(context)
	if !errors.Is(err, nil) {
		handler.logger.Error(err)
	}
}

func (handler *ginHandler) get(ginContext *gin.Context) {
	context := GinContext{ginContext}

	err := handler.appController.ProductController.Get(context)
	if !errors.Is(err, nil) {
		handler.logger.Error(err)
	}
}

func (handler *ginHandler) findAllTranslated(ginContext *gin.Context) {
	context := GinContext{ginContext}

	err := handler.appController.ProductTranslatedController.FindAll(context)
	if !errors.Is(err, nil) {
		handler.logger.Error(err)
	}
}
