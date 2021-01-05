package notificationmessage

type Subscriber interface {
	OnRecord(*Record) error
}
