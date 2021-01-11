package notificationqueue_test

import (
	"strconv"
	"sync"
	"testing"

	"github.com/herb-go/notification"
	"github.com/herb-go/notification/notificationqueue"
)

type testDelivery struct {
	locker sync.Mutex
	data   []notification.Content
}

func (d *testDelivery) DeliveryName() string {
	return "test"
}
func (d *testDelivery) DeliveryType() string {
	return "test"
}
func (d *testDelivery) MustEscape(u string) string {
	return "escaped:" + u
}
func (d *testDelivery) Deliver(c notification.Content) (status notification.DeliveryStatus, receipt string, err error) {
	d.locker.Lock()
	defer d.locker.Unlock()
	d.data = append(d.data, c)
	return notification.DeliveryStatusSuccess, strconv.Itoa(len(d.data)), nil
}

func newTestDelivery() *testDelivery {
	return &testDelivery{}
}

func TestDeliveryCenter(t *testing.T) {
	c := notificationqueue.NewPlainDeliveryCenter()
	c["test1"] = newTestDelivery()
	c["test2"] = newTestDelivery()
	ac := notificationqueue.NewAtomicDeliveryCenter()
	l, err := ac.List()
	if err != nil || len(l) != 0 {
		t.Fatal(l, err)
	}
	d, err := ac.Get("test")
	if !notification.IsErrDeliveryNotFound(err) || d != nil {
		t.Fatal(d, err)
	}
	ac.SetDeliveryCenter(c)
	l, err = ac.List()
	if err != nil || len(l) != 2 {
		t.Fatal(l, err)
	}
	d, err = ac.Get("test1")
	if err != nil || d == nil {
		t.Fatal(d, err)
	}
}

func newTestDeliveryCenter() notificationqueue.DeliveryCenter {
	c := notificationqueue.NewPlainDeliveryCenter()
	c["test1"] = newTestDelivery()
	c["test2"] = newTestDelivery()
	ac := notificationqueue.NewAtomicDeliveryCenter()
	ac.SetDeliveryCenter(c)
	return ac
}
