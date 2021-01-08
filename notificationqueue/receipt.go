package notificationqueue

import "github.com/herb-go/notification"

type Receipt struct {
	NotificationID string
	ExecutionID    string
	Status         notification.DeliveryStatus
	Message        string
}

func NewReceipt() *Receipt {
	return &Receipt{}
}
