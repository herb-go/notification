package notificationqueue

import "github.com/herb-go/notification"

//Receipt notification receipt struct
type Receipt struct {
	//NotificationID notification id
	NotificationID string
	//ExecutionID notification eecution id
	ExecutionID string
	//Status delivery status
	Status notification.DeliveryStatus
	//Message. receipt for successfully delivery or resofn fail fail delivery
	Message string
}

//NewReceipt create new receipt
func NewReceipt() *Receipt {
	return &Receipt{}
}
