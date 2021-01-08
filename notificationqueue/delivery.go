package notificationqueue

import (
	"sync/atomic"

	"github.com/herb-go/notification"
)

type DeliveryCenter interface {
	List() ([]notification.DeliveryServer, error)
	Get(name string) (notification.DeliveryServer, error)
}

type PlainDeliveryCenter map[string]notification.DeliveryServer

func (c PlainDeliveryCenter) List() ([]notification.DeliveryServer, error) {
	result := []notification.DeliveryServer{}
	for k := range c {
		result = append(result, c[k])
	}
	return result, nil
}
func (c PlainDeliveryCenter) Get(name string) (notification.DeliveryServer, error) {
	s, ok := c[name]
	if !ok || s == nil {
		return nil, notification.ErrorDeliveryNotFound(name)
	}
	return s, nil
}
func NewPlainDeliveryCenter() PlainDeliveryCenter {
	return PlainDeliveryCenter{}
}

type AtomicDeliveryCenter struct {
	data atomic.Value
}

func (c *AtomicDeliveryCenter) SetDeliveryCenter(pc PlainDeliveryCenter) {
	c.data.Store(pc)
}
func (c *AtomicDeliveryCenter) DeliveryCenter() PlainDeliveryCenter {
	return c.data.Load().(PlainDeliveryCenter)
}

func (c *AtomicDeliveryCenter) List() ([]notification.DeliveryServer, error) {
	return c.DeliveryCenter().List()
}
func (c *AtomicDeliveryCenter) Get(name string) (notification.DeliveryServer, error) {
	return c.DeliveryCenter().Get(name)
}

func NewAtomicDeliveryCenter() *AtomicDeliveryCenter {
	c := &AtomicDeliveryCenter{}
	c.SetDeliveryCenter(NewPlainDeliveryCenter())
	return c
}
