package interfaces

type INotificationService interface {
	Notify(INotification) error
}
