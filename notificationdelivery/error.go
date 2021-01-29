package notificationdelivery

import "fmt"

//ErrDeliveryNotFound error raised if given delivery not found
type ErrDeliveryNotFound struct {
	Delivery string
}

//Error return error message
func (e *ErrDeliveryNotFound) Error() string {
	return fmt.Sprintf("notification delivery: delivery not found [%s]", e.Delivery)
}

//NewErrDeliveryNotFound create new ErrDeliveryNotFound
func NewErrDeliveryNotFound(delivery string) *ErrDeliveryNotFound {
	return &ErrDeliveryNotFound{
		Delivery: delivery,
	}
}

//IsErrDeliveryNotFound check if given error is ErrDeliveryNotFound.
func IsErrDeliveryNotFound(err error) bool {
	if err == nil {
		return false
	}
	_, ok := err.(*ErrDeliveryNotFound)
	return ok
}
