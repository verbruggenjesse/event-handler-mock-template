package handlers

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
	"github.com/verbruggenjesse/event-handler-mock-template/mock"
	"github.com/verbruggenjesse/event-handler-mock-template/models"
)

// EchoMessageHandler is the event handler for topic "message" and action "echo"
func EchoMessageHandler(payload map[string]interface{}, hash string, metadata map[string]string) {
	var message models.Message

	mapstructure.Decode(payload, &message)

	fmt.Println(message.Content)

	echoNotification := &models.Notification{
		Type:   "global",
		Target: "any",
		Payload: map[string]interface{}{
			"message": message.Content,
		},
		Hash: fmt.Sprintf("n:%s", hash),
	}

	var ns mock.NotificationService

	ns.Publish(echoNotification)
}
