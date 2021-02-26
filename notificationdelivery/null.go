package notificationdelivery

import "github.com/herb-go/notification"

//DeliveryNull delivery null keyword
const DeliveryNull = "null"

//NullDelivery delivery do nothing
type NullDelivery struct {
}

//DeliveryType Delivery type
func (d NullDelivery) DeliveryType() string {
	return DeliveryNull
}

//Deliver send give content.
//Return delivery status and any receipt if returned,and any error if raised.
func (d NullDelivery) Deliver(notification.Content) (status DeliveryStatus, receipt string, err error) {
	return DeliveryStatusSuccess, "", nil
}

//CheckInvalidContent check if content invalid
//Return invalid fields and any error raised
func (d NullDelivery) CheckInvalidContent(notification.Content) ([]string, error) {
	return []string{}, nil
}

//ContentFields return content fields
//Return invalid fields and any error raised
func (d NullDelivery) ContentFields() []*Field {
	return nil
}

//NullFactory null factory
var NullFactory = func(loader func(v interface{}) error) (DeliveryDriver, error) {
	return NullDelivery{}, nil
}

//RegisterNullFactory register null factory.
func RegisterNullFactory() {
	Register(DeliveryNull, NullFactory)
}

func init() {
	RegisterNullFactory()
}
