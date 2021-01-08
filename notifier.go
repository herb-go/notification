package notification

type Notifier interface {
	Notify(*Notification) error
}
