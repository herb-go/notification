package notification

type DeliveryServer interface {
	DeliveryID() string
	DeliveryType() string
	Deliver(Content) (DeliveryStatus, error)
}

type DeliveryStatus int64

const (
	DeliveryStatusFail = DeliveryStatus(iota)
	DeliveryStatusSuccess
	DeliveryStatusAbort
)
