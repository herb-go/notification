package notificationdelivery

import (
	"strconv"
	"sync"
	"testing"

	"github.com/herb-go/notification"
)

type testDelivery struct {
	locker sync.Mutex
	data   []notification.Content
}

func (d *testDelivery) DeliveryType() string {
	return "test"
}
func (d *testDelivery) MustEscape(u string) string {
	return "escaped:" + u
}
func (d *testDelivery) Deliver(c notification.Content) (status DeliveryStatus, receipt string, err error) {
	d.locker.Lock()
	defer d.locker.Unlock()
	d.data = append(d.data, c)
	return DeliveryStatusSuccess, strconv.Itoa(len(d.data)), nil
}
func (d *testDelivery) CheckInvalidContent(notification.Content) ([]string, error) {
	return []string{}, nil
}
func newTestDelivery(id string) *DeliveryServer {
	s := NewDeliveryServer()
	s.Delivery = id
	s.DeliveryDriver = &testDelivery{}
	return s
}

func TestDeliveryCenter(t *testing.T) {
	c := NewPlainDeliveryCenter()
	c["test1"] = newTestDelivery("test1")
	c["test2"] = newTestDelivery("test2")
	ac := NewAtomicDeliveryCenter()
	l, err := ac.List()
	if err != nil || len(l) != 0 {
		t.Fatal(l, err)
	}
	d, err := ac.Get("test")
	if !IsErrDeliveryNotFound(err) || d != nil {
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

func newTestDeliveryCenter() DeliveryCenter {
	c := NewPlainDeliveryCenter()
	c["test1"] = newTestDelivery("test1")
	c["test2"] = newTestDelivery("test2")
	ac := NewAtomicDeliveryCenter()
	ac.SetDeliveryCenter(c)
	return ac
}
