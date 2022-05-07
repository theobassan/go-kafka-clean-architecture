package router

type JsonContextBind func(i interface{}) error
type JsonContextJSON func(code int, i interface{}) error
type JsonContextQuery func(key string) string
