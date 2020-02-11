package models

import "fmt"

// GlobalNotification is the model for notifications aimed at all connected users
type GlobalNotification struct {
	payload map[string]interface{}
	eventID string
}

// NewGlobalNotification is the constructor for a global notification
func NewGlobalNotification(payload map[string]interface{}, eventID string) *GlobalNotification {
	return &GlobalNotification{
		payload: payload,
		eventID: eventID,
	}
}

// Type returns the type of notification this is
func (n *GlobalNotification) Type() string {
	return "global"
}

// Targets returns targets for this notification
func (n *GlobalNotification) Targets() []string {
	return make([]string, 0)
}

// Payload returns the payload for this notification
func (n *GlobalNotification) Payload() map[string]interface{} {
	return n.payload
}

// ID returns the ID generated for this notification
func (n *GlobalNotification) ID() string {
	return fmt.Sprintf("gn-%s", n.eventID)
}
