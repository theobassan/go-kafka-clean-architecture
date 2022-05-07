package api

type EventApi interface {
	Bind(topic string, value []byte) interface{}
	WriteMessage(i interface{}) error
}
