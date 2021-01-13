package notificationdelivery

var NullFactory = func(loader func(v interface{}) error) (DeliveryDriver, error) {
	return NullDelivery{}, nil
}

func registerNull() {
	Register("null", NullFactory)
}

func init() {
	registerNull()
}
