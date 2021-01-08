package notificationqueue

import (
	"time"

	"github.com/herb-go/notification"
)

var NopReceiptHanlder = func(nid string, eid string, reason string, status int32) {}

type Notifier struct {
	DeliveryCenter
	queue       Queue
	c           chan int
	OnExecution func(*Execution)
	OnReceipt   func(nid string, eid string, status notification.DeliveryStatus, msg string)
	Recovery    func()
}

func (n *Notifier) SetQueue(q Queue) {
	n.queue = q
}
func (notifier *Notifier) handleReceipt(nid string, eid string, status notification.DeliveryStatus, msg string) {
	defer notifier.Recovery()
	notifier.OnReceipt(nid, eid, status, msg)
}

func (notifier *Notifier) execute(e *Execution) {
	defer notifier.Recovery()
	notifier.OnExecution(e)

}
func (notifier *Notifier) deliver(e *Execution) {
	defer notifier.Recovery()
	status, msg, err := notifier.DeliverNotification(e.Notification)
	if err != nil {
		status = notification.DeliveryStatusFail
		msg = err.Error()
	}
	nid := e.Notification.ID
	eid := e.ExecutionID
	go notifier.handleReceipt(nid, eid, status, msg)
	if status == notification.DeliveryStatusFail {
		return
	}
	err = notifier.queue.Remove(nid)
	if err != nil {
		panic(err)
	}
}
func (notifier *Notifier) listen(c chan *Execution) {
	for {
		select {
		case e := <-c:
			go notifier.execute(e)
		case _ = <-notifier.c:
			return
		}
	}
}
func (notifier *Notifier) Notify(n *notification.Notification) error {
	return notifier.queue.Push(n)
}

func (notifier *Notifier) Start() error {
	c, err := notifier.queue.PopChan()
	if err != nil {
		return err
	}
	go notifier.listen(c)
	return notifier.queue.Start()
}

func (notifier *Notifier) Stop() error {
	close(notifier.c)
	notifier.c = nil
	return notifier.queue.Stop()
}

func (notifier *Notifier) DeliverNotification(n *notification.Notification) (status notification.DeliveryStatus, receipt string, err error) {
	if n.ExpiredTime >= time.Now().Unix() {
		return notification.DeliveryStatusExpired, "", nil
	}
	d, err := notifier.DeliveryCenter.Get(n.Delivery)
	if err != nil {
		return 0, "", err
	}
	return d.Deliver(n.Content)
}

func NewNotifier() *Notifier {
	return &Notifier{}
}
