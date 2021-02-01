package notificationrender

import (
	"github.com/herb-go/notification"
)

//Renderer renderer interface
type Renderer interface {
	//Render render notification with given message
	Render(notification.Message) (*notification.Notification, error)
	//Supported return supported directives.
	Supported() (directives []string, err error)
}
