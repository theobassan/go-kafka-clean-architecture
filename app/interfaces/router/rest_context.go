package router

type RestContext interface {
	JSON(code int, i interface{}) error
	Bind(i interface{}) error
	Param(key string) string
}
