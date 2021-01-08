package notificationqueue

import "github.com/herb-go/notification"

var NopReceiptHanlder = func(nid string, eid string, reason string, status int32) {}

type Stream struct {
	Queue
	ReceiptHandler func(nid string, eid string, status int32, msg string)
}

func (s *Stream) ReturnReceipt(nid string, eid string, status int32, msg string) error {
	go s.ReceiptHandler(nid, eid, status, msg)
	return s.Remove(nid)
}
func (s *Stream) SetReceiptHandler(h func(nid string, eid string, status int32, msg string)) {
	s.ReceiptHandler = h
}

type Queue interface {
	PopChan() (chan *Execution, error)
	Push(*notification.Notification) error
	Remove(nid string) error
	Start() error
	Stop() error
}
