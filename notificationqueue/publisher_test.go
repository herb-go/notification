package notificationqueue_test

import (
	"testing"
	"time"

	"github.com/herb-go/notification"
	"github.com/herb-go/notification/notificationqueue"
)

var loggedNotifications []*notification.Notification

func testOnNotification(n *notification.Notification) {
	loggedNotifications = append(loggedNotifications, n)
}

var loggedErrors []error

func testOnError(err error) {
	loggedErrors = append(loggedErrors, err)
}

var loggerExecution []*notificationqueue.Execution

func testOnExecution(e *notificationqueue.Execution) {
	loggerExecution = append(loggerExecution, e)
}

func initLog() {
	loggedNotifications = []*notification.Notification{}
	loggedErrors = []error{}
	loggerExecution = []*notificationqueue.Execution{}
}

func newTestPublisher() *notificationqueue.Publisher {
	p := notificationqueue.NewPublisher()
	p.DraftReviewer = testDraftReviewer
	n := notificationqueue.NewNotifier()
	p.Notifier = n
	n.DeliveryCenter = newTestDeliveryCenter()
	n.SetQueue(newTestQueue())
	p.OnNotification = testOnNotification
	p.OnExecution = testOnExecution
	n.DeliveryCenter = newTestDeliveryCenter()
	p.Draftbox = newTestDraft()
	p.OnError = testOnError
	return p
}

func TestPublisher(t *testing.T) {
	initLog()
	p := newTestPublisher()
	err := p.Start()
	if err != nil {
		panic(err)
	}
	defer func() {
		err := p.Stop()
		if err != nil {
			panic(err)
		}
	}()
	n := notification.New()
	n.ID = mustID()
	n.Header[notification.HeaderNameDraftMode] = "1"
	ok, err := p.PublishNotification(n)
	if ok || err != nil {
		t.Fatal(ok, err)
	}
	n2, err := p.PublishDraft(n.ID)
	if n2.ID != n.ID || err != nil {
		t.Fatal(n2, err)
	}
	n = notification.New()
	n.ID = mustID()
	ok, err = p.PublishNotification(n)
	if !ok || err != nil {
		t.Fatal(ok, err)
	}
	n = notification.New()
	n.ID = mustID()
	n.Delivery = "test1"
	ok, err = p.PublishNotification(n)
	if !ok || err != nil {
		t.Fatal(ok, err)
	}
	time.Sleep(100 * time.Microsecond)
}
