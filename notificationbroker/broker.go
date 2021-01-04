package notificationbroker

import (
	"github.com/herb-go/notification"
)

type Broker interface {
	Start() error
	Stop() error
	Reset() error
	Subscribe(notification.Topic, notification.Subscriber) error
	PublishMessage(*notification.Message) error
	HandleErrors(notification.ErrorHandler)
	HandleRecords(notification.RecordHandler)
	HandleNotifcations(notification.NotificationHandler)
}
