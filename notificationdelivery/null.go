package notificationdelivery

//NullFactory null factory
var NullFactory = func(loader func(v interface{}) error) (DeliveryDriver, error) {
	return NullDelivery{}, nil
}

//RegisterNullFactory register null factory.
func RegisterNullFactory() {
	Register("null", NullFactory)
}

func init() {
	RegisterNullFactory()
}
