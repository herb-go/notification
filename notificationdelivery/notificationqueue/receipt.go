package notificationqueue

import (
	"github.com/herb-go/notification"
	"github.com/herb-go/notification/notificationdelivery"
)

//Receipt notification receipt struct
type Receipt struct {
	//NotificationID notification id
	Notification *notification.Notification
	//ExecutionID notification execution id
	ExecutionID string
	//Status delivery status
	Status notificationdelivery.DeliveryStatus
	//Message. receipt for successfully delivery or resofn fail fail delivery
	Message string
}

//NewReceipt create new receipt
func NewReceipt() *Receipt {
	return &Receipt{}
}
