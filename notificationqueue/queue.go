package notificationqueue

import "github.com/herb-go/notification"

var NopReceiptHanlder = func(nid string, eid string, reason string, status int) {}

type Queue interface {
	PopChan() (chan *Execution, error)
	Push(*notification.Notification) error
	ReturnSuccessReceipt(nid string, eid string, receipt string) error
	ReturnFailReceipt(nid string, eid string, reason string) error
	ReturnAbortReceipt(nid string, eid string, reason string) error
	SetReceiptHandler(func(nid string, eid string, reason string, status int))
}
