package notification

type Publisher interface {
	PublishMessage(*Message) error
}
