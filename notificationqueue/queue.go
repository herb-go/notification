package notificationqueue

import "github.com/herb-go/notification"

var NopReceiptHanlder = func(nid string, eid string, reason string, status int) {}

type Queue interface {
	PopChan() (chan *Execution, error)
	Push(*notification.Notification) error
	ReturnReceipt(nid string, eid string, status int32, content string) error
	SetReceiptHandler(func(nid string, eid string, reason string, status int))
}
