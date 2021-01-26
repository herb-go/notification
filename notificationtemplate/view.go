package notificationtemplate

import (
	"github.com/herb-go/notification"
)

type View interface {
	Render(Collection) (*notification.Notification, error)
}
