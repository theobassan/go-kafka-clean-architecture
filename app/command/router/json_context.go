package router

type JsonContext interface {
	JSON(code int, i interface{}) error
	Bind(i interface{}) error
	Query(key string) string
}
