package notificationqueue_test

import (
	"testing"
	"time"

	"github.com/herb-go/notification/notificationdelivery"

	"github.com/herb-go/notification"
	"github.com/herb-go/notification/notificationdelivery/notificationqueue"
)

var loggedNotifications []*notification.Notification

func testOnNotification(n *notification.Notification) {
	loggedNotifications = append(loggedNotifications, n)
}

var loggedErrors []error

func testRecover() {
	r := recover()
	if r != nil {
		err := r.(error)
		loggedErrors = append(loggedErrors, err)
	}
}

var loggedtestReceipt []*notificationqueue.Receipt

func testReceipt(r *notificationqueue.Receipt) {
	loggedtestReceipt = append(loggedtestReceipt, r)
}
func initLog() {
	loggedNotifications = []*notification.Notification{}
	loggedErrors = []error{}
	loggedtestReceipt = []*notificationqueue.Receipt{}
}

func newTestPublisher() *notificationqueue.Publisher {
	p := notificationqueue.NewPublisher()
	p.DraftReviewer = notificationqueue.DraftReviewerHeader
	n := notificationqueue.NewNotifier()
	p.Notifier = n
	n.DeliveryCenter = newTestDeliveryCenter()
	n.SetQueue(newTestQueue())
	p.OnNotification = testOnNotification
	p.Draftbox = newTestDraft()
	p.Recover = testRecover
	p.OnReceipt = testReceipt
	return p
}

func TestPublisher(t *testing.T) {
	initLog()
	p := newTestPublisher()
	q := p.Queue()
	_, ok := q.(testQueue)
	if !ok {
		t.Fatal(q)
	}
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
	n.Header[notification.HeaderNameTarget] = "1"
	ok, err = p.PublishNotification(n)
	if ok || err != nil {
		t.Fatal(ok, err)
	}
	time.Sleep(100 * time.Millisecond)

	n2, err := p.PublishDraft(n.ID)
	if n2.ID != n.ID || err != nil {
		t.Fatal(n2, err)
	}
	n = notification.New()
	n.Header[notification.HeaderNameTarget] = "2"
	n.ID = mustID()
	ok, err = p.PublishNotification(n)
	if !ok || err != nil {
		t.Fatal(ok, err)
	}
	time.Sleep(100 * time.Millisecond)
	n = notification.New()
	n.Header[notification.HeaderNameTarget] = "3"
	n.ID = mustID()
	n.Delivery = "test1"
	ok, err = p.PublishNotification(n)
	if !ok || err != nil {
		t.Fatal(ok, err)
	}
	time.Sleep(500 * time.Millisecond)
	if len(loggedErrors) != 2 || !notificationdelivery.IsErrDeliveryNotFound(loggedErrors[0]) || !notificationdelivery.IsErrDeliveryNotFound(loggedErrors[1]) {
		t.Fatal(len(loggedErrors))
	}
	if len(loggedNotifications) != 3 {
		t.Fatal(len(loggedNotifications))
	}
	if len(loggedtestReceipt) != 3 ||
		loggedtestReceipt[0].Status != notificationdelivery.DeliveryStatusFail ||
		loggedtestReceipt[1].Status != notificationdelivery.DeliveryStatusFail ||
		loggedtestReceipt[2].Status != notificationdelivery.DeliveryStatusSuccess {
		t.Fatal(len(loggedNotifications))
	}
}
