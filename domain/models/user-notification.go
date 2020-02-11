package models

import "fmt"

// UserNotification is the model for notifications aimed at all connected users
type UserNotification struct {
	payload map[string]interface{}
	eventID string
	targets []string
}

// NewUserNotification is the constructor for a global notification
func NewUserNotification(payload map[string]interface{}, eventID string, targets []string) *UserNotification {
	if targets == nil {
		targets = make([]string, 0)
	}

	return &UserNotification{
		payload: payload,
		eventID: eventID,
		targets: targets,
	}
}

// Type returns the type of notification this is
func (n *UserNotification) Type() string {
	return "user"
}

// Targets returns targets for this notification
func (n *UserNotification) Targets() []string {
	return n.targets
}

// Payload returns the payload for this notification
func (n *UserNotification) Payload() map[string]interface{} {
	return n.payload
}

// ID returns the ID generated for this notification
func (n *UserNotification) ID() string {
	return fmt.Sprintf("un-%s", n.eventID)
}
