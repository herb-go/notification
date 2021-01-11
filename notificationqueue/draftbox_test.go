package notificationqueue_test

import (
	"sync"

	"github.com/herb-go/notification"
	"github.com/herb-go/notification/notificationqueue"
)

type testDraft struct {
	locker sync.Mutex
	data   []*notification.Notification
}

func (d *testDraft) Draft(notification *notification.Notification) error {
	locker.Lock()
	defer locker.Unlock()
	for k := range d.data {
		if d.data[k].ID == notification.ID {
			d.data[k] = notification
			return nil
		}
	}
	d.data = append(d.data, notification)
	return nil
}
func (d *testDraft) List(condition []notificationqueue.Condition, start string, asc bool, count int) (result []*notification.Notification, iter string, err error) {
	return nil, "", nil
}
func (d *testDraft) SupportedConditions() ([]string, error) {
	return nil, nil
}
func (d *testDraft) Eject(id string) (*notification.Notification, error) {
	for k := range d.data {
		if d.data[k].ID == id {
			n := d.data[k]
			d.data = append(d.data[:k], d.data[k:]...)
			return n, nil
		}
	}
	return nil, notification.ErrNofitactionIDNotFound(id)
}

func newTestDraft() *testDraft {
	return &testDraft{}
}

var testDraftReviewer = notificationqueue.FuncDraftReviewer(func(n *notification.Notification) (publishable bool, err error) {
	return n.Header[notification.HeaderNameDraftMode] != "", nil
})
