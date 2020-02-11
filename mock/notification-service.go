package mock

import (
	"fmt"

	"github.com/verbruggenjesse/event-handler-mock-template/domain/interfaces"
)

// NotificationService is the mock implementation of the actual notification service
type NotificationService struct{}

// Notify sends out notifications
func (n *NotificationService) Notify(notification interfaces.INotification) error {
	fmt.Printf("Sending %s notification to %s [%s] with payload %v\n", notification.Type(), notification.Targets(), notification.ID(), notification.Payload())
	return nil
}
