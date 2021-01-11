package notificationqueue

import (
	"github.com/herb-go/notification"
)

const ConditionTopic = "topic"
const ConditionNotificationID = "notificationid"
const ConditionTarget = "target"
const ConditionBatch = "batch"
const ConditionInContent = "incontent"
const ConditionDelivery = "delivery"
const ConditionBeforeTimestamp = "beforetimestamp"
const ConditionAeforeTimestamp = "aftertimestamp"

type Condition struct {
	Keyword string
	Value   string
}

type Draftbox interface {
	Draft(notification *notification.Notification) error
	List(condition []Condition, start string, asc bool, count int) (result []*notification.Notification, iter string, err error)
	SupportedConditions() ([]string, error)
	Eject(id string) (*notification.Notification, error)
}

type DraftReviewer interface {
	ReviewDraft(*notification.Notification) (publishable bool, err error)
}

type FuncDraftReviewer func(*notification.Notification) (publishable bool, err error)

func (r FuncDraftReviewer) ReviewDraft(n *notification.Notification) (publishable bool, err error) {
	return r(n)
}
