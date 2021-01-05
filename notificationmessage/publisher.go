package notificationmessage

type Publisher interface {
	PublishMessage(*Message) error
}
