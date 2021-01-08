package notification

type DeliveryStatus int64

type DeliveryServer interface {
	DeliveryName() string
	Driver
}

type Driver interface {
	DeliveryType() string
	MustEscape(string) string
	Deliver(Content) (status DeliveryStatus, receipt string, err error)
}

const (
	DeliveryStatusFail    = DeliveryStatus(0)
	DeliveryStatusSuccess = DeliveryStatus(1)
	DeliveryStatusAbort   = DeliveryStatus(2)
)
