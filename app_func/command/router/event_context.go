package router

type EventContextBind func(v any) error
type EventContextAcknowledge func() error
