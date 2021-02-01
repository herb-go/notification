package notificationview

import (
	"testing"

	"github.com/herb-go/notification"
)

type nopView struct{}

//View render notification with given data
func (r nopView) Render(Message) (*notification.Notification, error) {
	return nil, nil
}

func TestViewCenter(t *testing.T) {
	ac := NewAtomicViewCenter()
	r, err := ac.Get("test")
	if err == nil || !IsErrViewNotFound(err) || r != nil {
		t.Fatal(r, err)
	}
	c := NewViewCenter()
	r, err = c.Get("test")
	if err == nil || !IsErrViewNotFound(err) || r != nil {
		t.Fatal(r, err)
	}
	c.Set("test", nopView{})
	r, err = c.Get("test")
	if err != nil || r == nil {
		t.Fatal(r, err)
	}
	ac.SetViewCenter(c)
	r, err = ac.Get("test")
	if err != nil || r == nil {
		t.Fatal(r, err)
	}
	current := ac.ViewCenter()
	pc, ok := current.(*PlainViewCenter)
	if !ok || pc != c {
		t.Fatal(current)
	}
}
