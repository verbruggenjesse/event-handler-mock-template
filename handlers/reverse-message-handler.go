package handlers

import (
	"fmt"

	"github.com/mitchellh/mapstructure"
	"github.com/verbruggenjesse/event-handler-mock-template/extensions"
	"github.com/verbruggenjesse/event-handler-mock-template/mock"
	"github.com/verbruggenjesse/event-handler-mock-template/models"
)

// ReverseMessageHandler is the event handler for topic "message" and action "reverse"
func ReverseMessageHandler(payload map[string]interface{}, hash string, metadata map[string]string) {
	var message models.Message

	mapstructure.Decode(payload, &message)

	reversed := extensions.Reverse(message.Content)

	fmt.Println(reversed)

	echoNotification := &models.Notification{
		Type:   "global",
		Target: "any",
		Payload: map[string]interface{}{
			"message": reversed,
		},
		Hash: fmt.Sprintf("n:%s", hash),
	}

	var ns mock.NotificationService

	ns.Publish(echoNotification)
}
