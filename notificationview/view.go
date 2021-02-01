package notificationview

import (
	"github.com/herb-go/notification"
)

//View view interface
type View interface {
	//Render render notification with given message.
	//Nil should be returned if notification should not be send.
	Render(Message) (*notification.Notification, error)
}
