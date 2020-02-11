package main

import (
	"fmt"
	"log"
	"os"

	"github.com/verbruggenjesse/event-handler-mock-template/handler"
	"github.com/verbruggenjesse/event-handler-mock-template/infrastructure"
	"github.com/verbruggenjesse/event-handler-mock-template/mock"
	"google.golang.org/grpc"
)

func main() {
	var serverAddr string

	if serverAddr = os.Getenv("EVENT_CENTRAL_ADDR"); serverAddr == "" {
		log.Fatalln("Missing required environment variable 'EVENT_CENTRAL_ADDR'")
	}

	conn, err := grpc.Dial(serverAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatal("Could not connect to central service")
	}
	defer conn.Close()

	eventService := infrastructure.NewGrpcEventService(conn)

	notificationService := &mock.NotificationService{}

	echoHandler := handler.NewEchoMessageHandler(notificationService)
	reverseHandler := handler.NewReverseMessageHandler(notificationService)

	eventService.Subscribe("message", "echo", "", echoHandler)
	eventService.Subscribe("message", "reverse", "", reverseHandler)

	if err = eventService.ListenForEvents(); err != nil {
		fmt.Println(err)
	}
}
