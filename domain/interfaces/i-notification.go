package interfaces

// INotification is the interface for notifications
type INotification interface {
	Type() string
	Targets() []string
	Payload() map[string]interface{}
	ID() string
}
