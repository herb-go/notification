package notificationdelivery

import "github.com/herb-go/notification"

var NullFactory = func(loader func(v interface{}) error) (notification.DeliveryDriver, error) {
	return notification.NullDelivery{}, nil
}

func registerNull() {
	Register("null", NullFactory)
}

func init() {
	registerNull()
}
