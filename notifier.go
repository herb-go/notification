package notification

//Notifier notifier interafce
type Notifier interface {
	//Notify asynchronous delivery given notification and return any error if raised.
	Notify(*Notification) error
}
