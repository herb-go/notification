package notificationqueue

import "github.com/herb-go/notification"

var ExecuteStatusFail = int32(0)
var ExecuteStatusSuccess = int32(1)
var ExecuteStatusAbort = int32(2)

type Notifier struct {
	DraftReviewer  DraftReviewer
	Queue          Queue
	Draftbox       Draftbox
	DeliveryCenter DeliveryCenter
	OnExecution    func(*Execution)
	OnError        func(error)
	c              chan int
}

func (notifier *Notifier) Notify(n *notification.Notification) (bool, error) {
	ok, err := notifier.DraftReviewer.ReviewDraft(n)
	if err != nil {
		return false, err
	}
	if ok {
		return false, notifier.Draftbox.Draft(n)
	}
	err = notifier.Queue.Push(n)
	return err == nil, err
}

func (notifier *Notifier) PublishDraft(nid string) (*notification.Notification, error) {
	n, err := notifier.Draftbox.Discard(nid)
	if err != nil {
		return nil, err
	}
	return n, notifier.Queue.Push(n)
}
func (notifier *Notifier) listen(c chan *Execution) {
	for {
		select {
		case e := <-c:
			go notifier.OnExecution(e)
		case _ = <-notifier.c:
			return
		}
	}
}
func (notifier *Notifier) Start() error {
	c, err := notifier.Queue.PopChan()
	if err != nil {
		return err
	}
	go notifier.listen(c)
	return nil
}

func (notifier *Notifier) Stop() error {
	close(notifier.c)
	notifier.c = nil
	return nil
}

func (notifier *Notifier) DeliverNotification(n *notification.Notification) (status notification.DeliveryStatus, receipt string, err error) {
	d, err := notifier.DeliveryCenter.Get(n.Delivery)
	if err != nil {
		return 0, "", err
	}
	return d.Deliver(n.Content)
}
func NewNotifier() *Notifier {
	return &Notifier{}
}
