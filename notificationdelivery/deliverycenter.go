package notificationdelivery

import (
	"sort"
	"sync/atomic"
	"time"

	"github.com/herb-go/notification"
)

//DeliveryCenter delivery center interface.
//Delivery center manages delivery servers by delivery keyword.
type DeliveryCenter interface {
	//List all delivery servers in delivery center and any error if raised.
	List() ([]*DeliveryServer, error)
	//Get get delivery server by keyword and return any error if rasied.
	//Notification.ErrDeliveryNotFound should be returned if give keyword not found.
	Get(keyword string) (*DeliveryServer, error)
}

//PlainDeliveryCenter plain delivery center type
type PlainDeliveryCenter map[string]*DeliveryServer

//List all delivery servers in delivery center and any error if raised.
//Result will be sorted by delivery.
func (c PlainDeliveryCenter) List() ([]*DeliveryServer, error) {
	var keys = make([]string, 0, len(c))
	for k := range c {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	var result = make([]*DeliveryServer, len(keys))

	for k, v := range keys {
		result[k] = c[v]
	}
	return result, nil
}

//Get get delivery server by keyword and return any error if rasied.
//Notification.ErrDeliveryNotFound should be returned if give keyword not found.
func (c PlainDeliveryCenter) Get(id string) (*DeliveryServer, error) {
	s, ok := c[id]
	if !ok || s == nil {
		return nil, NewErrDeliveryNotFound(id)
	}
	return s, nil
}

//Insert insert delivery server to c
func (c PlainDeliveryCenter) Insert(d *DeliveryServer) {
	c[d.Delivery] = d
}

//NewPlainDeliveryCenter create new plain delivery center
func NewPlainDeliveryCenter() PlainDeliveryCenter {
	return PlainDeliveryCenter{}
}

//AtomicDeliveryCenter delivery center which use atomic.Value to  implement concurrently update.
type AtomicDeliveryCenter struct {
	data atomic.Value
}

//SetDeliveryCenter atomicly update delivery center.
func (c *AtomicDeliveryCenter) SetDeliveryCenter(pc DeliveryCenter) {
	c.data.Store(pc)
}

//DeliveryCenter return delivery center actually used.
func (c *AtomicDeliveryCenter) DeliveryCenter() DeliveryCenter {
	return c.data.Load().(DeliveryCenter)
}

//List all delivery servers in delivery center and any error if raised.
func (c *AtomicDeliveryCenter) List() ([]*DeliveryServer, error) {
	return c.DeliveryCenter().List()
}

//Get get delivery server by keyword and return any error if rasied.
//Notification.ErrDeliveryNotFound should be returned if give keyword not found.
func (c *AtomicDeliveryCenter) Get(id string) (*DeliveryServer, error) {
	return c.DeliveryCenter().Get(id)
}

//NewAtomicDeliveryCenter create new atomic delivery cemter.
func NewAtomicDeliveryCenter() *AtomicDeliveryCenter {
	c := &AtomicDeliveryCenter{}
	c.SetDeliveryCenter(NewPlainDeliveryCenter())
	return c
}

//Deliver delivery content to delivery center with given keyword
//Return delivery status,receipt and any error if raised.
func Deliver(c DeliveryCenter, delivery string, content notification.Content) (status DeliveryStatus, receipt string, err error) {
	d, err := c.Get(delivery)
	if err != nil {
		if IsErrDeliveryNotFound(err) {
			return DeliveryStatusAbort, err.Error(), nil
		}
		return 0, "", err
	}
	return d.Deliver(content)
}

//DeliverNotification notification to delivery center with given keyword
//Return delivery status,receipt and any error if raised.
func DeliverNotification(c DeliveryCenter, n *notification.Notification) (status DeliveryStatus, receipt string, err error) {
	if n.ExpiredTime > 0 && n.ExpiredTime <= time.Now().Unix() {
		return DeliveryStatusExpired, "", nil
	}
	return Deliver(c, n.Delivery, n.Content)
}
