package notification

//Sender sender interafce
type Sender interface {
	//Send send notification and return any error if raised.
	Send(*Notification) error
}

//SenderFunc sender func interface
type SenderFunc func(*Notification) error

//Send send notification and return any error if raised.
func (f SenderFunc) Send(n *Notification) error {
	return f(n)
}
