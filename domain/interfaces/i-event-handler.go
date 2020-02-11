package interfaces

// IEventHandler is the interface for eventhandlers
type IEventHandler interface {
	Handle(event IEvent) error
}
