package notificationqueue

import (
	"github.com/herb-go/notification"
)

type Execution struct {
	ExecutionID    string
	Notification   *notification.Notification
	RetryCount     int32
	StartTime      int64
	RetryAfterTime int64
}

func NewExecution() *Execution {
	return &Execution{}
}

type ExecutionCreator interface {
	CreateExecution(*notification.Notification) (*Execution, error)
}
