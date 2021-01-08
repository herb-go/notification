package notificationqueue

import "github.com/herb-go/notification"

type Publisher struct {
	DraftReviewer DraftReviewer
	Draftbox      Draftbox
	*Notifier
}

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

func (publisher *Publisher) PublishDraft(nid string) (*notification.Notification, error) {
	n, err := publisher.Draftbox.Discard(nid)
	if err != nil {
		return nil, err
	}
	return n, publisher.Notifier.Notify(n)
}

func NewPublisher() *Publisher {
	return &Publisher{}
}
