package json_context

import "github.com/labstack/echo"

type EchoContext struct {
	echoContext echo.Context
}

func (context EchoContext) JSON(code int, i interface{}) error {
	context.echoContext.JSON(code, i)
	return nil
}

func (context EchoContext) Bind(i interface{}) error {
	return context.echoContext.Bind(i)
}

func (context EchoContext) Query(key string) string {
	return context.echoContext.QueryParam(key)
}
