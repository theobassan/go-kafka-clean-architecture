package json_context

import (
	"github.com/gin-gonic/gin"
)

type GinContext struct {
	ginContext *gin.Context
}

func (context GinContext) JSON(code int, i interface{}) error {
	context.ginContext.JSON(code, i)
	return nil
}

func (context GinContext) Bind(i interface{}) error {
	return context.ginContext.Bind(i)
}

func (context GinContext) Query(key string) string {
	return context.ginContext.Query(key)
}
