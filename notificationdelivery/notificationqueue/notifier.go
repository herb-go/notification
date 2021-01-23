package notificationqueue

import (
	"time"

	"github.com/herb-go/notification"
	"github.com/herb-go/notification/notificationdelivery"
)

//Notifier notifier struct
type Notifier struct {
	//DeliveryCenter
	notificationdelivery.DeliveryCenter
	//Workers push workers num
	Workers int
	queue   Queue
	c       chan int
	//IDGenerator id generator
	IDGenerator func() (string, error)
	//OnNotification notification handler
	OnNotification func(*notification.Notification)
	//OnReceipt receipt handler
	OnReceipt func(*Receipt)
	//OnExecution execution handler
	OnExecution func(*Execution)
	//Recover recover handler
	Recover func()
}

//SetQueue set queue to notifier.
//SetQueue should be called before start
func (notifier *Notifier) SetQueue(q Queue) {
	notifier.queue = q
}

//HandleDeliverTimeout deliver timeout handler
func (notifier *Notifier) HandleDeliverTimeout(e *Execution) {
	notifier.newReceipt(e, notificationdelivery.DeliveryStatusTimeout, "")
}

//HandleRetryTooMany retry toomany handler
func (notifier *Notifier) HandleRetryTooMany(e *Execution) {
	notifier.newReceipt(e, notificationdelivery.DeliveryStatusRetryTooMany, "")

}

//Queue return notifier queue.
func (notifier *Notifier) Queue() Queue {
	return notifier.queue
}
func (notifier *Notifier) handleReceipt(r *Receipt) {
	defer notifier.Recover()
	notifier.OnReceipt(r)
}
func (notifier *Notifier) newReceipt(e *Execution, status notificationdelivery.DeliveryStatus, msg string) {
	eid := e.ExecutionID
	r := NewReceipt()
	r.Notification = e.Notification
	r.ExecutionID = eid
	r.Status = status
	r.Message = msg
	go notifier.handleReceipt(r)
}
func (notifier *Notifier) deliver(e *Execution) {
	defer notifier.Recover()
	status, msg, err := notifier.deliveryNotification(e.Notification)
	if err != nil {
		status = notificationdelivery.DeliveryStatusFail
		msg = err.Error()
		go func() {
			defer notifier.Recover()
			panic(err)
		}()
	}
	notifier.newReceipt(e, status, msg)
	if notificationdelivery.IsStatusRetryable(status) {
		return
	}
	err = notifier.queue.Remove(e.Notification.ID)
	if err != nil {
		panic(err)
	}
}
func (notifier *Notifier) listen(c <-chan *Execution) {
	for {
		select {
		case e := <-c:
			notifier.deliver(e)
		case <-notifier.c:
			return
		}
	}
}

func (notifier *Notifier) onNotification(n *notification.Notification) {
	defer notifier.Recover()
	notifier.OnNotification(n)
}
func (notifier *Notifier) InitNotification(n *notification.Notification) error {
	id, err := notifier.IDGenerator()
	if err != nil {
		return err
	}
	n.ID = id
	n.CreatedTime = time.Now().Unix()
	return nil
}

//Notify delivery notification
//Notification id  will be returned
func (notifier *Notifier) Notify(n *notification.Notification) error {
	err := notifier.queue.Push(n)
	if err == nil {
		go notifier.onNotification(n)
	}
	return err
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

//Reset reset notifier handlers
func (notifier *Notifier) Reset() {
	notifier.OnNotification = NopNotificationHandler
	notifier.OnReceipt = NopReceiptHanlder
	// notifier.OnDeliverTimeout = NopExecutionHandler
	// notifier.OnRetryTooMany = NopExecutionHandler
	notifier.OnExecution = NopExecutionHandler
	notifier.IDGenerator = NopIDGenerator
}

//NewNotifier create new notifier
func NewNotifier() *Notifier {
	n := &Notifier{
		queue:          &NopQueue{},
		DeliveryCenter: notificationdelivery.NewAtomicDeliveryCenter(),
	}
	n.Reset()
	return n
}
