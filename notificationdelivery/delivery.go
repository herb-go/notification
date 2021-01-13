package notificationdelivery

import "github.com/herb-go/notification"

//DeliveryStatus delivery status type
type DeliveryStatus int64

const (
	//DeliveryStatusFail stands for delivery fail
	DeliveryStatusFail = DeliveryStatus(0)
	//DeliveryStatusSuccess stands for delivery success
	DeliveryStatusSuccess = DeliveryStatus(1)
	//DeliveryStatusAbort stands for delivery abort
	DeliveryStatusAbort = DeliveryStatus(2)
	//DeliveryStatusExpired stands for delivery expired
	DeliveryStatusExpired = DeliveryStatus(3)
	//DeliveryStatusDisabled stands for delivery expired
	DeliveryStatusDisabled = DeliveryStatus(4)
)

//DeliveryServer delivery server struct
type DeliveryServer struct {
	//Delivery delivery id
	Delivery string
	//Disabled is delivery disabled
	Disabled bool
	//Description delivery server description
	Description string
	//Driver delivery driver
	DeliveryDriver
}

//NewDeliveryServer create new delivery server
func NewDeliveryServer() *DeliveryServer {
	return &DeliveryServer{}
}

//DeliveryDriver Delivery driver
type DeliveryDriver interface {
	//DeliveryType Delivery type
	DeliveryType() string
	//MustEscape delivery escape helper
	MustEscape(string) string
	//Deliver send give content.
	//Return delivery status and any receipt if returned,and any error if raised.
	Deliver(notification.Content) (status DeliveryStatus, receipt string, err error)
}
