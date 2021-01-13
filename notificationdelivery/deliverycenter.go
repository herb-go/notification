package notificationdelivery

import (
	"sync/atomic"
	"time"

	"github.com/herb-go/notification"
)

//DeliveryCenter delivery center interface.
//Delivery center manages delivery servers by delivery keyword.
type DeliveryCenter interface {
	//List all delivery servers in delivery center and any error if raised.
	List() ([]*notification.DeliveryServer, error)
	//Get get delivery server by keyword and return any error if rasied.
	//Notification.ErrDeliveryNotFound should be returned if give keyword not found.
	Get(keyword string) (*notification.DeliveryServer, error)
}

//PlainDeliveryCenter plain delivery center type
type PlainDeliveryCenter map[string]*notification.DeliveryServer

//List all delivery servers in delivery center and any error if raised.
func (c PlainDeliveryCenter) List() ([]*notification.DeliveryServer, error) {
	result := []*notification.DeliveryServer{}
	for k := range c {
		result = append(result, c[k])
	}
	return result, nil
}

//Get get delivery server by keyword and return any error if rasied.
//Notification.ErrDeliveryNotFound should be returned if give keyword not found.
func (c PlainDeliveryCenter) Get(id string) (*notification.DeliveryServer, error) {
	s, ok := c[id]
	if !ok || s == nil {
		return nil, notification.NewErrDeliveryNotFound(id)
	}
	return s, nil
}

//Insert insert delivery server to c
func (c PlainDeliveryCenter) Insert(d *notification.DeliveryServer) {
	c[d.DeliveryType()] = d
}

//NewPlainDeliveryCenter create new plain delivery center
func NewPlainDeliveryCenter() PlainDeliveryCenter {
	return PlainDeliveryCenter{}
}

//AtomicDeliveryCenter delivery center which use atomic.Value to  implement concurrently update
type AtomicDeliveryCenter struct {
	data atomic.Value
}

//SetDeliveryCenter atomicly update delivery center.
func (c *AtomicDeliveryCenter) SetDeliveryCenter(pc DeliveryCenter) {
	c.data.Store(pc)
}

//DeliveryCenter returm delivery center actually used.
func (c *AtomicDeliveryCenter) DeliveryCenter() DeliveryCenter {
	return c.data.Load().(DeliveryCenter)
}

//List all delivery servers in delivery center and any error if raised.
func (c *AtomicDeliveryCenter) List() ([]*notification.DeliveryServer, error) {
	return c.DeliveryCenter().List()
}

//Get get delivery server by keyword and return any error if rasied.
//Notification.ErrDeliveryNotFound should be returned if give keyword not found.
func (c *AtomicDeliveryCenter) Get(id string) (*notification.DeliveryServer, error) {
	return c.DeliveryCenter().Get(id)
}

//NewAtomicDeliveryCenter create new atomic delivery cemter.
func NewAtomicDeliveryCenter() *AtomicDeliveryCenter {
	c := &AtomicDeliveryCenter{}
	c.SetDeliveryCenter(NewPlainDeliveryCenter())
	return c
}

func Deliver(c DeliveryCenter, delivery string, content notification.Content) (status notification.DeliveryStatus, receipt string, err error) {
	d, err := c.Get(delivery)
	if err != nil {
		return 0, "", err
	}
	return d.Deliver(content)
}
func DeliverNotification(c DeliveryCenter, n *notification.Notification) (status notification.DeliveryStatus, receipt string, err error) {
	if n.ExpiredTime > 0 && n.ExpiredTime <= time.Now().Unix() {
		return notification.DeliveryStatusExpired, "", nil
	}
	return Deliver(c, n.Delivery, n.Content)
}