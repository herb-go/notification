package notificationqueue

import "github.com/herb-go/notification"

type Queue interface {
	PopChan() (chan *Execution, error)
	Push(*notification.Notification) error
	Remove(nid string) error
	Start() error
	Stop() error
}
