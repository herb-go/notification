package notificationqueue

import "github.com/herb-go/notification"

//Publisher publisher struct
type Publisher struct {
	//DraftReviewer checker that check if given notification should be published or put in draft box.
	//Return true if notification should sendt to draftbox
	//Default value is CheckerDraftModeHeader
	DraftReviewer notification.Checker
	Draftbox      notification.Store
	*Notifier
}

//PublishNotification generate notification id and publish given notification
//Return notification id and if notification is published.
func (publisher *Publisher) PublishNotification(n *notification.Notification) (string, bool, error) {
	err := publisher.Notifier.InitNotification(n)
	if err != nil {
		return "", false, err
	}
	ok, err := publisher.DraftReviewer.Check(n)
	if err != nil {
		return "", false, err
	}
	if ok {
		return n.ID, false, publisher.Draftbox.Save(n)
	}
	err = publisher.Notifier.Notify(n)
	return n.ID, err == nil, err
}

//PublishDraft publish notificaiton by given id.
//Notification.ErrNotificationIDNotFound will be returned if nid not found.
func (publisher *Publisher) PublishDraft(nid string) (*notification.Notification, error) {
	n, err := publisher.Draftbox.Remove(nid)
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

//DraftboxLoader draftbox loader
func (publisher *Publisher) DraftboxLoader() notification.Store {
	return publisher.Draftbox
}

//NewPublisher create new publisher
func NewPublisher() *Publisher {
	return &Publisher{
		DraftReviewer: CheckerDraftModeHeader,
		Notifier:      NewNotifier(),
	}
}
