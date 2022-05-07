package api

type EventApiBind func(topic string, value []byte) interface{}
type EventApiWriteMessage func(i interface{}) error
