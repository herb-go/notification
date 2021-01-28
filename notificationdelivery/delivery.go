package notificationdelivery

import "github.com/herb-go/notification"

//DeliveryStatus delivery status type
type DeliveryStatus int64

const (
	//DeliveryStatusFail status for delivery fail
	DeliveryStatusFail = DeliveryStatus(0)
	//DeliveryStatusSuccess status for delivery success
	DeliveryStatusSuccess = DeliveryStatus(1)
	//DeliveryStatusAbort status for delivery abort
	DeliveryStatusAbort = DeliveryStatus(2)
	//DeliveryStatusExpired status for delivery expired
	DeliveryStatusExpired = DeliveryStatus(3)
	//DeliveryStatusDisabled status for delivery disabled
	DeliveryStatusDisabled = DeliveryStatus(4)
	//DeliveryStatusTimeout status for delivery timeout
	DeliveryStatusTimeout = DeliveryStatus(5)
	//DeliveryStatusRetryTooMany status for delivery too many
	DeliveryStatusRetryTooMany = DeliveryStatus(6)
)

//IsStatusRetryable chech if status retryable
func IsStatusRetryable(s DeliveryStatus) bool {
	return s == DeliveryStatusFail || s == DeliveryStatusTimeout
}

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

//Deliver send give content.
//Return delivery status and any receipt if returned,and any error if raised.
//DeliveryStatusDisabled will be returned if DeliveryServer is disabled.
func (s *DeliveryServer) Deliver(c notification.Content) (status DeliveryStatus, receipt string, err error) {
	if s.Disabled {
		return DeliveryStatusDisabled, "", nil
	}
	return s.DeliveryDriver.Deliver(c)
}

//NewDeliveryServer create new delivery server
func NewDeliveryServer() *DeliveryServer {
	return &DeliveryServer{}
}

//DeliveryDriver Delivery driver
type DeliveryDriver interface {
	//DeliveryType Delivery type
	DeliveryType() string
	//CheckInvalidContent check if given content invalid
	//Return invalid fields and any error raised
	CheckInvalidContent(notification.Content) ([]string, error)
	//Deliver send give content.
	//Return delivery status and any receipt if returned,and any error if raised.
	Deliver(notification.Content) (status DeliveryStatus, receipt string, err error)
}
