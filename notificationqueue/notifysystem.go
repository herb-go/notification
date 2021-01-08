package notificationqueue

import "github.com/herb-go/notification"

var ExecuteStatusFail = int32(0)
var ExecuteStatusSuccess = int32(1)
var ExecuteStatusAbort = int32(2)

type NotifySystem struct {
	DraftReviewer DraftReviewer
	Draftbox      Draftbox
	*Notifier
}

func (notifysystem *NotifySystem) Notify(n *notification.Notification) (bool, error) {
	ok, err := notifysystem.DraftReviewer.ReviewDraft(n)
	if err != nil {
		return false, err
	}
	if ok {
		return false, notifysystem.Draftbox.Draft(n)
	}
	err = notifysystem.Notifier.Notify(n)
	return err == nil, err
}

func (notifysystem *NotifySystem) PublishDraft(nid string) (*notification.Notification, error) {
	n, err := notifysystem.Draftbox.Discard(nid)
	if err != nil {
		return nil, err
	}
	return n, notifysystem.Notifier.Notify(n)
}

func NewNotifySystem() *NotifySystem {
	return &NotifySystem{}
}
