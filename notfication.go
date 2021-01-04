package notification

import "fmt"

type Notification struct {
	Record
	InstanceID       string
	DeliveryServerID string
}

func (n *Notification) String() string {
	return fmt.Sprintf("[%s|%s] target %s @ topic %s", n.InstanceID, n.RecordID, n.Target, n.Topic)
}

type NotificationInstance struct {
	Notification
	Schedule *Schedule
	Content  *Content
}
