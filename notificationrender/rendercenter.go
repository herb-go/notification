package notificationrender

import "sync/atomic"

//RenderCenter render center interface
type RenderCenter interface {
	//Get get renderer by given name.
	//ErrRendererNotFound should be returned if given name not found.
	Get(name string) (Renderer, error)
}

//PlainRenderCenter plain render center struct
type PlainRenderCenter map[string]Renderer

//Get get renderer by given name.
//ErrRendererNotFound should be returned if given name not found.
func (c *PlainRenderCenter) Get(name string) (Renderer, error) {
	r, ok := (*c)[name]
	if ok {
		return nil, NewErrRendererNotFound(name)
	}
	return r, nil
}

//NewRenderCenter create new plain render center
func NewRenderCenter() *PlainRenderCenter {
	return &PlainRenderCenter{}
}

//AtomicRenderCenter render center which use atomic.Value to implement concurrently update,
type AtomicRenderCenter struct {
	data atomic.Value
}

//SetRenderCenter atomicly update render center.
func (c *AtomicRenderCenter) SetRenderCenter(center RenderCenter) {
	c.data.Store(center)
}

//RenderCenter return render center actually used.
func (c *AtomicRenderCenter) RenderCenter() RenderCenter {
	return c.data.Load().(RenderCenter)
}

//Get get renderer by given name.
//ErrRendererNotFound should be returned if given name not found.
func (c *AtomicRenderCenter) Get(name string) (Renderer, error) {
	return c.RenderCenter().Get(name)
}

//NewAtomicRenderCenter create new atomic render center.
func NewAtomicRenderCenter() *AtomicRenderCenter {
	c := &AtomicRenderCenter{}
	c.SetRenderCenter(NewRenderCenter())
	return c
}
