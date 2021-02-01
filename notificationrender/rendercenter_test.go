package notificationrender

import (
	"testing"

	"github.com/herb-go/notification"
)

type nopRender struct{}

//Render render notification with given data
func (r nopRender) Render(notification.Message) (*notification.Notification, error) {
	return nil, nil
}

//Supported return supported directives.
func (r nopRender) Supported() (directives []string, err error) {
	return nil, nil
}

func TestRenderCenter(t *testing.T) {
	ac := NewAtomicRenderCenter()
	r, err := ac.Get("test")
	if err == nil || !IsErrRendererNotFound(err) || r != nil {
		t.Fatal(r, err)
	}
	c := NewRenderCenter()
	r, err = c.Get("test")
	if err == nil || !IsErrRendererNotFound(err) || r != nil {
		t.Fatal(r, err)
	}
	c.Set("test", nopRender{})
	r, err = c.Get("test")
	if err != nil || r == nil {
		t.Fatal(r, err)
	}
	ac.SetRenderCenter(c)
	r, err = ac.Get("test")
	if err != nil || r == nil {
		t.Fatal(r, err)
	}
	current := ac.RenderCenter()
	pc, ok := current.(*PlainRenderCenter)
	if !ok || pc != c {
		t.Fatal(current)
	}
}
