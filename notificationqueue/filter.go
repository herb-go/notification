package notificationqueue

import (
	"strconv"
	"strings"

	"github.com/herb-go/notification"
)

//Filter notification filter interface
type Filter interface {
	//FilterNotification filter notification with given timestamp
	//Return if notification is valid
	FilterNotification(n *notification.Notification, ts int64) (bool, error)
	//ApplyCondition apply search condition to filter
	//ErrConditionNotSupported should be returned if condition keyword is not supported
	ApplyCondition(cond *Condition) error
}

//PlainFilter plain filter struct
type PlainFilter struct {
	BatchID        string
	NotificationID string
	Delivery       string
	Target         string
	Topic          string
	InContent      string
	After          int64
	Before         int64
}

//ApplyCondition apply search condition to filter
//ErrConditionNotSupported should be returned if condition keyword is not supported
func (c *PlainFilter) ApplyCondition(cond *Condition) error {
	switch cond.Keyword {
	case ConditionBatch:
		c.BatchID = cond.Value
	case ConditionNotificationID:
		c.NotificationID = cond.Value
	case ConditionDelivery:
		c.Delivery = cond.Value
	case ConditionTarget:
		c.Target = cond.Value
	case ConditionTopic:
		c.Topic = cond.Value
	case ConditionAfterTimestamp:
		ts, err := strconv.ParseInt(cond.Value, 10, 64)
		if err != nil {
			return NewErrInvalidConditionValue(ConditionAfterTimestamp)
		}
		c.After = ts
	case ConditionBeforeTimestamp:
		ts, err := strconv.ParseInt(cond.Value, 10, 64)
		if err != nil {
			return NewErrInvalidConditionValue(ConditionBeforeTimestamp)
		}
		c.Before = ts
	default:
		return NewErrConditionNotSupported(cond.Keyword)
	}
	return nil
}

//FilterNotification filter notification with given timestamp
//Return if notification is valid
func (c *PlainFilter) FilterNotification(n *notification.Notification, ts int64) (bool, error) {
	if c.BatchID != "" && n.Header.Get(notification.HeaderNameBatch) != c.BatchID {
		return false, nil
	}
	if c.NotificationID != "" && n.ID != c.NotificationID {
		return false, nil
	}
	if c.Delivery != "" && n.Delivery != c.Delivery {
		return false, nil
	}
	if c.Target != "" && n.Header.Get(notification.HeaderNameTarget) != c.Target {
		return false, nil
	}
	if c.Topic != "" && n.Header.Get(notification.HeaderNameTopic) != c.Topic {
		return false, nil
	}
	if c.InContent != "" {
		var found bool
		for k := range n.Content {
			if strings.Contains(n.Content[k], c.InContent) {
				found = true
				break
			}
		}
		if !found {
			return false, nil
		}
	}
	if c.After > 0 && c.After <= ts {
		return false, nil
	}
	if c.Before > 0 && c.Before >= ts {
		return false, nil
	}
	return true, nil
}

//NewFilter create new plain filter
func NewFilter() *PlainFilter {
	return &PlainFilter{}
}

//ApplyToFilter apply condiitons to filter.
func ApplyToFilter(f Filter, conds []*Condition) error {
	for k := range conds {
		err := f.ApplyCondition(conds[k])
		if err != nil {
			return err
		}
	}
	return nil
}
