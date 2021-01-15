package notificationqueue

import "github.com/herb-go/notification"

//Publisher publisher struct
type Publisher struct {
	//DraftReviewer default value DraftReviewerHeader
	DraftReviewer DraftReviewer
	Draftbox      Draftbox
	*Notifier
}

//PublishNotification publish given notification
//Return if notfication is published.
func (publisher *Publisher) PublishNotification(n *notification.Notification) (bool, error) {
	ok, err := publisher.DraftReviewer.ReviewDraft(n)
	if err != nil {
		return false, err
	}
	if ok {
		return false, publisher.Draftbox.Draft(n)
	}
	err = publisher.Notifier.Notify(n)
	return err == nil, err
}

//PublishDraft publish notificaiton by given id.
//Notification.ErrNotificationIDNotFound will be returned if nid not found.
func (publisher *Publisher) PublishDraft(nid string) (*notification.Notification, error) {
	n, err := publisher.Draftbox.Eject(nid)
	if err != nil {
		return nil, err
	}
	return n, publisher.Notifier.Notify(n)
}

//Start start publisher
func (publisher *Publisher) Start() error {
	err := publisher.Draftbox.Open()
	if err != nil {
		return err
	}
	return publisher.Notifier.Start()
}

//Stop stop publisher
func (publisher *Publisher) Stop() error {
	err := publisher.Draftbox.Close()
	if err != nil {
		go func() {
			defer publisher.Notifier.Recover()
			panic(err)
		}()
	}
	return publisher.Notifier.Stop()

}

//NewPublisher create new publisher
func NewPublisher() *Publisher {
	return &Publisher{
		DraftReviewer: DraftReviewerHeader,
		Notifier:      NewNotifier(),
	}
}
