package notificationqueue

import (
	"github.com/herb-go/notification"
	"github.com/herb-go/notification/notificationdelivery"
)

//NopReceiptHanlder nop receipt hanlder
var NopReceiptHanlder = func(nid string, eid string, reason string, status int32) {}

//Notifier notifier struct
type Notifier struct {
	//DeliveryCenter
	notificationdelivery.DeliveryCenter
	//Workers push workers num
	Workers int
	queue   Queue
	c       chan int
	//OnNotification notification handler
	OnNotification func(*notification.Notification)
	//OnExecution execution handler
	OnExecution func(*Execution)
	//OnReceipt receipt handler
	OnReceipt func(*Receipt)
	//OnError error handler
	OnError func(error)
}

//SetQueue set queue to notifier.
//SetQueue should be called before start
func (notifier *Notifier) SetQueue(q Queue) {
	notifier.queue = q
}

//Recover recover with notifier.OnError
func (notifier *Notifier) Recover() {
	r := recover()
	if r != nil {
		err := r.(error)
		go notifier.OnError(err)
	}
}
func (notifier *Notifier) handleReceipt(r *Receipt) {
	defer notifier.Recover()
	notifier.OnReceipt(r)
}

func (notifier *Notifier) execute(e *Execution) {
	go notifier.OnExecution(e)
	notifier.deliver(e)

}
func (notifier *Notifier) deliver(e *Execution) {
	defer notifier.Recover()
	status, msg, err := notifier.deliveryNotification(e.Notification)
	if err != nil {
		go notifier.OnError(err)
		status = notificationdelivery.DeliveryStatusFail
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
	if status == notificationdelivery.DeliveryStatusFail {
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
	defer notifier.Recover()
	notifier.OnNotification(n)
}

//Notify delivery notifiction
func (notifier *Notifier) Notify(n *notification.Notification) error {
	go notifier.onNotification(n)
	return notifier.queue.Push(n)
}

//Start start notifier
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

//Stop stop notifier
func (notifier *Notifier) Stop() error {
	close(notifier.c)
	notifier.c = nil
	return notifier.queue.Stop()
}

func (notifier *Notifier) deliveryNotification(n *notification.Notification) (status notificationdelivery.DeliveryStatus, receipt string, err error) {
	return notificationdelivery.DeliverNotification(notifier.DeliveryCenter, n)

}

//NewNotifier create new notifier
func NewNotifier() *Notifier {
	return &Notifier{
		DeliveryCenter: notificationdelivery.NewAtomicDeliveryCenter(),
	}
}
