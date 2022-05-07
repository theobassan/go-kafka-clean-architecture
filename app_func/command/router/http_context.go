package router

import "net/http"

type HttpContextResponseWriter func() http.ResponseWriter
type HttpContextRequest func() *http.Request
