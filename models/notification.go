package models

// Notification is the model for notifications
type Notification struct {
	Type    string
	Target  string
	Payload map[string]interface{}
	Hash    string
}
