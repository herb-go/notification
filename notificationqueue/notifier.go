package notificationqueue

import "github.com/herb-go/notification"

type Notifier struct {
	DeliveryCenter DeliveryCenter
	Queue          Queue
	c              chan int
	OnExecution    func(*Execution)
	Recovery       func()
}

func (notifier *Notifier) execute(e *Execution) {
	defer notifier.Recovery()
	notifier.OnExecution(e)
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
	return notifier.Queue.Push(n)
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
