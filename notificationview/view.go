package notificationview

import (
	"github.com/herb-go/herbtext"
	"github.com/herb-go/notification"
)

//View view interface
type View interface {
	//Render render notification with given message.
	//Nil should be returned if notification should not be send.
	Render(Message) (*notification.Notification, error)
}

//CloneView view clone message to notification content
type CloneView struct {
	Delivery string
}

//Render render notification with given message.
//Nil should be returned if notification should not be send.
func (v *CloneView) Render(m Message) (*notification.Notification, error) {
	n := notification.New()
	n.Delivery = v.Delivery
	herbtext.MergeSet(n.Content, m)
	return n, nil
}

//CloneViewFactoryName clone view factory name
const CloneViewFactoryName = "clone"

//CloneViewFactory clone view factory
func CloneViewFactory(loader func(v interface{}) error) (View, error) {
	v := &CloneView{}
	err := loader(&v)
	if err != nil {
		return nil, err
	}
	return v, nil
}
