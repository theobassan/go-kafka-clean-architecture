package http_context

import (
	"go-kafka-clean-architecture/app/input/router"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GinContext struct {
	ginContext *gin.Context
}

func NewGinContextt(ginContext *gin.Context) router.HttpContext {
	return &GinContext{ginContext}
}

func (context *GinContext) ResponseWriter() http.ResponseWriter {
	return context.ginContext.Writer
}

func (context *GinContext) Request() *http.Request {
	return context.ginContext.Request
}
