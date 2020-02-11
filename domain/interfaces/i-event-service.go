package interfaces

// IEventService is the interface for event-services
type IEventService interface {
	ListenForEvents() error
	Subscribe(topic string, action string, lastID string, handler IEventHandler) error
}
