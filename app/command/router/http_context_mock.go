package router

import (
	"net/http"
)

func NewHttpContextMock(responseWriter http.ResponseWriter, request *http.Request) HttpContext {
	return &HttpContextMock{responseWriter, request}
}

type HttpContextMock struct {
	responseWriter http.ResponseWriter
	request        *http.Request
}

func (context *HttpContextMock) ResponseWriter() http.ResponseWriter {
	return context.responseWriter
}

func (context *HttpContextMock) Request() *http.Request {
	return context.request
}
