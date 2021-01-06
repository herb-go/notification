package notification

type DeliveryServer interface {
	DeliveryType() string
	MustEscape(string) string
	Deliver(Content) (staus DeliveryStatus, receipt string, err error)
}

type DeliveryStatus int64

const (
	DeliveryStatusFail = DeliveryStatus(iota)
	DeliveryStatusSuccess
	DeliveryStatusAbort
)
