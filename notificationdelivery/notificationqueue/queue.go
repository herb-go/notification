package notificationqueue

import "github.com/herb-go/notification"

//Queue notification deliver queue
type Queue interface {
	//PopChan return execution chan
	PopChan() (<-chan *Execution, error)
	//Push push notification to queue
	Push(*notification.Notification) error
	//Remove remove notification by given id
	Remove(nid string) error
	//Start start queue
	Start() error
	//Stop stop queue
	Stop() error
	//AttachTo attach queue to notifier
	AttachTo(*Notifier) error
	//Detach detach queue.
	Detach() error
}

//NopQueue struct
type NopQueue struct {
}

//PopChan return execution chan
func (*NopQueue) PopChan() (<-chan *Execution, error) {
	return nil, ErrQueueDriverRequired
}

//Push push notification to queue
func (*NopQueue) Push(*notification.Notification) error {
	return ErrQueueDriverRequired
}

//Remove remove notification by given id
func (*NopQueue) Remove(nid string) error {
	return ErrQueueDriverRequired
}

//Start start queue
func (*NopQueue) Start() error {
	return ErrQueueDriverRequired
}

//Stop stop queue
func (*NopQueue) Stop() error {
	return ErrQueueDriverRequired
}

//AttachTo attach queue to notifier
func (*NopQueue) AttachTo(*Notifier) error {
	return ErrQueueDriverRequired
}

//Detach detach queue.
func (*NopQueue) Detach() error {
	return ErrQueueDriverRequired
}
