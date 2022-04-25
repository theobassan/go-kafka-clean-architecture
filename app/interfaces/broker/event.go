package broker

type EventReader interface {
	Bind(i interface{}) []byte
	ReadMessage() (interface{}, error)
}

type EventWriter interface {
	Bind(value []byte) interface{}
	WriteMessage(i interface{}) error
}
