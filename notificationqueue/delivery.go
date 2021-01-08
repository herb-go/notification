package notificationqueue

import "github.com/herb-go/notification"

type DeliveryCenter interface {
	List() ([]notification.DeliveryServer, error)
	Get(name string) (notification.DeliveryServer, error)
}
