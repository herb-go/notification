package notificationqueue_test

import (
	"strconv"
	"sync"

	"github.com/herb-go/notification"
	"github.com/herb-go/notification/notificationdelivery"
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
func (d *testDelivery) Deliver(c notification.Content) (status notification.DeliveryStatus, receipt string, err error) {
	d.locker.Lock()
	defer d.locker.Unlock()
	d.data = append(d.data, c)
	return notification.DeliveryStatusSuccess, strconv.Itoa(len(d.data)), nil
}

func newTestDelivery(id string) *notification.DeliveryServer {
	s := notification.NewDeliveryServer()
	s.Delivery = id
	s.DeliveryDriver = &testDelivery{}
	return s
}

func newTestDeliveryCenter() notificationdelivery.DeliveryCenter {
	c := notificationdelivery.NewPlainDeliveryCenter()
	c["test1"] = newTestDelivery("test1")
	c["test2"] = newTestDelivery("test2")
	ac := notificationdelivery.NewAtomicDeliveryCenter()
	ac.SetDeliveryCenter(c)
	return ac
}
