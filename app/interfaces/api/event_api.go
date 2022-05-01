package api

type EventAPI interface {
	Bind(topic string, value []byte) interface{}
	WriteMessage(i interface{}) error
}
