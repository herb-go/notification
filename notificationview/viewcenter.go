package notificationview

import "sync/atomic"

//ViewCenter render center interface
type ViewCenter interface {
	//Get get view by given name.
	//ErrViewNotFound should be returned if given name not found.
	Get(name string) (View, error)
}

//PlainViewCenter plain render center struct
type PlainViewCenter map[string]View

//Get get view by given name.
//ErrViewNotFound should be returned if given name not found.
func (c *PlainViewCenter) Get(name string) (View, error) {
	r, ok := (*c)[name]
	if !ok {
		return nil, NewErrViewNotFound(name)
	}
	return r, nil
}

//Set set view with given name.
func (c *PlainViewCenter) Set(name string, r View) {
	(*c)[name] = r
}

//NewViewCenter create new plain render center
func NewViewCenter() *PlainViewCenter {
	return &PlainViewCenter{}
}

//AtomicViewCenter render center which use atomic.Value to implement concurrently update,
type AtomicViewCenter struct {
	data atomic.Value
}

//SetViewCenter atomicly update render center.
func (c *AtomicViewCenter) SetViewCenter(center ViewCenter) {
	c.data.Store(center)
}

//ViewCenter return render center actually used.
func (c *AtomicViewCenter) ViewCenter() ViewCenter {
	return c.data.Load().(ViewCenter)
}

//Get get view by given name.
//ErrViewNotFound should be returned if given name not found.
func (c *AtomicViewCenter) Get(name string) (View, error) {
	return c.ViewCenter().Get(name)
}

//NewAtomicViewCenter create new atomic render center.
func NewAtomicViewCenter() *AtomicViewCenter {
	c := &AtomicViewCenter{}
	c.SetViewCenter(NewViewCenter())
	return c
}
