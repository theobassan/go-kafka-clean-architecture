package router

type EventContext interface {
	Bind(v any) error
	Acknowledge() error
}
