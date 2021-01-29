package notificationrender

import (
	"github.com/herb-go/notification"
)

//Renderer renderer interface
type Renderer interface {
	//Render render notification with given data
	Render(map[string]string) (*notification.Notification, error)
	//Supported return supported directives.
	Supported() (directives []string, err error)
}
