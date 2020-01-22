package main

import (
	"fmt"
	"github.com/verbruggenjesse/event-handler-mock-template/handlers"
	"github.com/verbruggenjesse/event-handler-mock-template/mock"
	"github.com/verbruggenjesse/event-handler-mock-template/models"
)

func main() {
	mockEvents := []models.Event{
		models.Event{
			Topic:  "message",
			Action: "echo",
			Payload: map[string]interface{}{
				"content": "Hello world!",
			},
			Hash: "HelloWorld1",
		},
		models.Event{
			Topic:  "message",
			Action: "echo",
			Payload: map[string]interface{}{
				"content": "Hello people!",
			},
			Hash: "HelloPeople1",
		},
		models.Event{
			Topic:  "message",
			Action: "reverse",
			Payload: map[string]interface{}{
				"content": "Hello world!",
			},
			Hash: "HelloWorld2",
		},
		models.Event{
			Topic:  "message",
			Action: "reverse",
			Payload: map[string]interface{}{
				"content": "Hello people!",
			},
			Hash: "HelloPeople2",
		},
	}

	eventService := mock.EventService{
		MockedQueue:  mockEvents,
		EventChannel: make(chan models.Event),
	}

	eventService.Subscribe("message", "echo", handlers.EchoMessageHandler)
	eventService.Subscribe("message", "reverse", handlers.ReverseMessageHandler)

	err := eventService.ListenForEvents()

	if err != nil {
		fmt.Println(err)
	}
}
