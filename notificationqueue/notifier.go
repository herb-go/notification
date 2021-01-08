package notificationqueue

import "github.com/herb-go/notification"

type Notifier struct {
	DeliveryCenter DeliveryCenter
	Stream         Stream
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
	return notifier.Stream.Push(n)
}

func (notifier *Notifier) Start() error {
	c, err := notifier.Stream.PopChan()
	if err != nil {
		return err
	}
	go notifier.listen(c)
	return notifier.Stream.Start()
}

func (notifier *Notifier) Stop() error {
	close(notifier.c)
	notifier.c = nil
	return notifier.Stream.Stop()
}

func (notifier *Notifier) DeliverNotification(n *notification.Notification) (status notification.DeliveryStatus, receipt string, err error) {
	d, err := notifier.DeliveryCenter.Get(n.Delivery)
	if err != nil {
		return 0, "", err
	}
	return d.Deliver(n.Content)
}
