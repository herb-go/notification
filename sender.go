package notification

//Sender sender interafce
type Sender interface {
	//Send send notification and return any error if raised.
	Send(*Notification) error
}
