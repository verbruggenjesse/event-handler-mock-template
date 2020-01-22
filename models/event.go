package models

// Event is the model for incoming events
type Event struct {
	Topic    string
	Action   string
	Payload  map[string]interface{}
	Metadata map[string]string
	Hash     string
}
