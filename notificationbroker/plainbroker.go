package notificationbroker

import (
	"sync"

	"github.com/herb-go/notification"
)

type PlainBroker struct {
	locker              sync.Mutex
	subscribers         map[notification.Topic][]notification.Subscriber
	IDGenerator         func() (string, error)
	errorHandler        notification.ErrorHandler
	recordHandler       notification.RecordHandler
	notificationHandler notification.NotificationHandler
}

func (b *PlainBroker) Start() error {
	return nil
}
func (b *PlainBroker) Stop() error {
	return nil
}
func (b *PlainBroker) Reset() error {
	b.locker.Lock()
	defer b.locker.Unlock()
	b.subscribers = map[notification.Topic][]notification.Subscriber{}
	return nil
}
func (b *PlainBroker) Subscribe(notification.Topic, notification.Subscriber) error {
	b.locker.Lock()
	defer b.locker.Unlock()
	return nil
}
func (b *PlainBroker) PublishMessage(*notification.Message) error {
	b.locker.Lock()
	defer b.locker.Unlock()
	return nil
}
func (b *PlainBroker) HandleErrors(h notification.ErrorHandler) {
	b.errorHandler = h
}
func (b *PlainBroker) HandleRecords(h notification.RecordHandler) {
	b.recordHandler = h
}
func (b *PlainBroker) HandleNotifcations(h notification.NotificationHandler) {
	b.notificationHandler = h
}

func New() *PlainBroker {
	return &PlainBroker{
		subscribers: map[notification.Topic][]notification.Subscriber{},
	}
}
