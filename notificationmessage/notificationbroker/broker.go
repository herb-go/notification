package notificationbroker

import (
	"github.com/herb-go/notification"
	"github.com/herb-go/notification/notificationmessage"
)

type Broker interface {
	Start() error
	Stop() error
	Reset() error
	Subscribe(notificationmessage.Topic, notificationmessage.Subscriber) error
	PublishMessage(*notificationmessage.Message) error
	HandleErrors(notification.ErrorHandler)
	HandleRecords(notificationmessage.RecordHandler)
}
