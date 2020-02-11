package infrastructure

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/verbruggenjesse/event-handler-mock-template/domain/interfaces"
	"github.com/verbruggenjesse/event-handler-mock-template/domain/models"
	pb "github.com/verbruggenjesse/event-handler-mock-template/gen"
	"google.golang.org/grpc"
)

// GrpcEventService is an event service that uses GRPC for transport
type GrpcEventService struct {
	client        *pb.EventCentralClient
	subscriptions map[*pb.Key]interfaces.IEventHandler
}

// NewGrpcEventService is the constructor for GrpcEventServices
func NewGrpcEventService(conn *grpc.ClientConn) *GrpcEventService {
	client := pb.NewEventCentralClient(conn)

	return &GrpcEventService{
		client:        &client,
		subscriptions: make(map[*pb.Key]interfaces.IEventHandler),
	}
}

// ListenForEvents will start listening for incoming events
func (s *GrpcEventService) ListenForEvents() error {
	keys := make([]*pb.Key, 0)

	for k := range s.subscriptions {
		keys = append(keys, k)
	}

	request := &pb.SubscribeRequest{
		Subscriptions: keys,
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	instance := *s.client
	stream, err := instance.Subscribe(ctx, request)

	if err != nil {
		return err
	}

	for {
		event, err := stream.Recv()

		if err != nil {
			return err
		}

		for key, handler := range s.subscriptions {
			if event.Topic == key.GetTopic() && event.Action == key.GetAction() {
				var payload map[string]interface{}

				json.Unmarshal([]byte(event.GetPayload()), &payload)

				parsedEvent := models.NewEvent(event.GetId(), event.GetTopic(), event.GetAction(), payload, event.GetMetadata())

				handler.Handle(parsedEvent)
			}
		}
	}

}

// Subscribe will register eventhandlers for events
// There should not be multiple eventhandlers for a single topic and action within the same application
func (s *GrpcEventService) Subscribe(topic string, action string, lastID string, handler interfaces.IEventHandler) {
	key := &pb.Key{
		Topic:  topic,
		Action: action,
		LastId: lastID,
	}

	s.subscriptions[key] = handler
}

func getStream(key *pb.Key) string {
	return fmt.Sprintf("%s:%s:%s", "events", key.GetTopic(), key.GetAction())
}
