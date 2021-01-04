package notification

type Model map[string]string

type NotificationModelLoader interface {
	LoadNotificationModel(*Message) (*Model, error)
}
