package mock

import (
	"errors"
	"time"

	"github.com/verbruggenjesse/event-handler-mock-template/models"
)

// EventService is the mock implementation of the actual event service
type EventService struct {
	EventChannel chan models.Event
	MockedQueue  []models.Event
}

// ListenForEvents will start emptying the queue of mock events to the eventchannel after one second
func (e *EventService) ListenForEvents() error {
	time.Sleep(100 * time.Millisecond)

	ticker := time.NewTicker(100 * time.Millisecond)

	// Every second, send next event in the mockedQueue to the eventchannel
	for range ticker.C {
		event := e.MockedQueue[0]
		e.MockedQueue = e.MockedQueue[1:]
		e.EventChannel <- event
		if len(e.MockedQueue) == 0 {
			time.Sleep(10 * time.Second)
			return errors.New("No events left to process")
		}
	}

	return nil
}

// Subscribe will register eventhandlers for certain topics and actions
func (e *EventService) Subscribe(topic string, action string, eventHandler func(map[string]interface{}, string, map[string]string)) {
	go func() {
		for event := range e.EventChannel {
			if event.Topic == topic && event.Action == action {
				eventHandler(event.Payload, event.Hash, event.Metadata)
			} else {
				// If it could not be processed, put it back in the back of the queue
				e.EventChannel <- event
			}
		}
	}()
}
