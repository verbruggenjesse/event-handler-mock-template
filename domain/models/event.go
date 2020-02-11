package models

// Event is the model for incoming events
type Event struct {
	id       string
	topic    string
	action   string
	payload  map[string]interface{}
	metadata map[string]string
}

// NewEvent is the constructor for new events
func NewEvent(id string, topic string, action string, payload map[string]interface{}, metadata map[string]string) *Event {
	return &Event{
		id:       id,
		topic:    topic,
		action:   action,
		payload:  payload,
		metadata: metadata,
	}
}

// ID is the getter for readonly prop id
func (e *Event) ID() string {
	return e.id
}

// Topic is the getter for readonly prop topic
func (e *Event) Topic() string {
	return e.topic
}

// Action is the getter for readonly prop action
func (e *Event) Action() string {
	return e.action
}

// Payload is the getter for readonly prop payload
func (e *Event) Payload() map[string]interface{} {
	return e.payload
}

// Metadata is the getter for readonly prop metadata
func (e *Event) Metadata() map[string]string {
	return e.metadata
}
