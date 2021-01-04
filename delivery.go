package notification

type DeliveryServer interface {
	DeliveryNotificationInstance(*NotificationInstance) error
}
