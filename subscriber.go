package notification

type Subscriber interface {
	OnNotification(*Notification) error
}
