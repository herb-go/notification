package notificationqueue

import (
	"sync/atomic"

	"github.com/herb-go/notification"
)

type DeliveryCenter interface {
	List() ([]notification.DeliveryServer, error)
	Get(name string) (notification.DeliveryServer, error)
}

type PlainDeliveryCenter struct {
	data atomic.Value
}
