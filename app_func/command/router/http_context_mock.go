package router

import (
	"net/http"
)

type HttpContextMock struct {
	responseWriter http.ResponseWriter
	request        *http.Request
}

func NewHttpContextMock(responseWriter http.ResponseWriter, request *http.Request) HttpContextMock {
	return HttpContextMock{responseWriter, request}
}
func (context HttpContextMock) ResponseWriter() http.ResponseWriter {
	return context.responseWriter
}

func (context HttpContextMock) Request() *http.Request {
	return context.request
}
