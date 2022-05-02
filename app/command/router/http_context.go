package router

import "net/http"

type HttpContext interface {
	ResponseWriter() http.ResponseWriter
	Request() *http.Request
}
