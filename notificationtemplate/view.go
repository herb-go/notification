package notificationtemplate

import (
	"github.com/herb-go/notification"
)

type View interface {
	Render(Model) (*notification.Notification, error)
	Supported() (directives []string, err error)
}
