package notification

type Notifier interface {
	Notify(*Notification) (sent bool, err error)
}
