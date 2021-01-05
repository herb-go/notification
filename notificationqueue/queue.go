package notificationqueue

import "github.com/herb-go/notification"

type Queue interface {
	Notfiy(*notification.Notification) error
	ListDeliveryServers() ([]notification.DeliveryServer, error)
}
