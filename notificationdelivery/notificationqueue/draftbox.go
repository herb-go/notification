package notificationqueue

import (
	"github.com/herb-go/notification"
)

var CheckerDraftModeHeader = notification.HasHeaderChecker(notification.HeaderNameDraftMode)
