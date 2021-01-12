package notificationqueue

import "github.com/herb-go/notification"

//Queue notification deliver queue
type Queue interface {
	//PopChan return execution chan
	PopChan() (chan *Execution, error)
	//Push push notification to queue
	Push(*notification.Notification) error
	//Remove remove notification by given id
	Remove(nid string) error
	//Start start queue
	Start() error
	//Stop stop queue
	Stop() error
}
