package notification

type Notifier interface {
	Notfiy(*Notification) error
}
