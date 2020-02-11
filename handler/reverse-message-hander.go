package handler

import (
	"fmt"

	"github.com/mitchellh/mapstructure"
	"github.com/verbruggenjesse/event-handler-mock-template/domain/interfaces"
	"github.com/verbruggenjesse/event-handler-mock-template/domain/models"
	"github.com/verbruggenjesse/event-handler-mock-template/extensions"
	"github.com/verbruggenjesse/event-handler-mock-template/mock"
)

type ReverseMessageHandler struct {
	notificationService interfaces.INotificationService
}

func NewReverseMessageHandler(ns interfaces.INotificationService) *ReverseMessageHandler {
	return &ReverseMessageHandler{
		notificationService: ns,
	}
}

func (h *ReverseMessageHandler) Handle(event interfaces.IEvent) error {
	var message models.Message

	mapstructure.Decode(event.Payload(), &message)

	fmt.Println(extensions.Reverse(message.Content))

	response := map[string]interface{}{
		"message": extensions.Reverse(message.Content),
	}

	reversNotification := models.NewGlobalNotification(response, event.ID())

	var ns mock.NotificationService

	ns.Notify(reversNotification)

	return nil
}
