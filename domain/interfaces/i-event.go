package interfaces

// IEvent is the interface for an event
type IEvent interface {
	ID() string
	Topic() string
	Action() string
	Payload() map[string]interface{}
	Metadata() map[string]string
}
