package http_context

import (
	"net/http"

	"github.com/labstack/echo"
)

type EchoContext struct {
	echoContext echo.Context
}

func NewEchoContext(echoContext echo.Context) EchoContext {
	return EchoContext{echoContext}
}

func (context EchoContext) ResponseWriter() http.ResponseWriter {
	return context.echoContext.Response().Writer
}

func (context EchoContext) Request() *http.Request {
	return context.echoContext.Request()
}
