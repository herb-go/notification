package notificationqueue

import (
	"time"

	"github.com/herb-go/notification"
)

var NopReceiptHanlder = func(nid string, eid string, reason string, status int32) {}

type Notifier struct {
	DeliveryCenter
	Workers        int
	queue          Queue
	c              chan int
	OnNotification func(*notification.Notification)
	OnExecution    func(*Execution)
	OnReceipt      func(*Receipt)
	OnError        func(error)
}

func (n *Notifier) SetQueue(q Queue) {
	n.queue = q
}
func (notifier *Notifier) Recovery() {
	r := recover()
	if r != nil {
		err := r.(error)
		notifier.OnError(err)
	}
}
func (notifier *Notifier) handleReceipt(r *Receipt) {
	defer notifier.Recovery()
	notifier.OnReceipt(r)
}

func (notifier *Notifier) execute(e *Execution) {
	go notifier.OnExecution(e)
	notifier.deliver(e)

}
func (notifier *Notifier) deliver(e *Execution) {
	defer notifier.Recovery()
	status, msg, err := notifier.deliverNotification(e.Notification)
	if err != nil {
		go notifier.OnError(err)
		status = notification.DeliveryStatusFail
		msg = err.Error()
	}
	nid := e.Notification.ID
	eid := e.ExecutionID
	r := NewReceipt()
	r.NotificationID = nid
	r.ExecutionID = eid
	r.Status = status
	r.Message = msg
	go notifier.handleReceipt(r)
	if status == notification.DeliveryStatusFail {
		return
	}
	err = notifier.queue.Remove(nid)
	if err != nil {
		go notifier.OnError(err)
	}
}
func (notifier *Notifier) listen(c chan *Execution) {
	for {
		select {
		case e := <-c:
			notifier.execute(e)
		case <-notifier.c:
			return
		}
	}
}

func (notifier *Notifier) onNotification(n *notification.Notification) {
	defer notifier.Recovery()
	notifier.OnNotification(n)
}
func (notifier *Notifier) Notify(n *notification.Notification) error {
	go notifier.onNotification(n)
	return notifier.queue.Push(n)
}

func (notifier *Notifier) Start() error {
	notifier.c = make(chan int)
	c, err := notifier.queue.PopChan()
	if err != nil {
		return err
	}
	workers := notifier.Workers
	if workers < 1 {
		workers = 1
	}
	for i := 0; i < workers; i++ {
		go notifier.listen(c)
	}
	return notifier.queue.Start()
}

func (notifier *Notifier) Stop() error {
	close(notifier.c)
	notifier.c = nil
	return notifier.queue.Stop()
}

func (notifier *Notifier) deliverNotification(n *notification.Notification) (status notification.DeliveryStatus, receipt string, err error) {
	if n.ExpiredTime > 0 && n.ExpiredTime <= time.Now().Unix() {
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
