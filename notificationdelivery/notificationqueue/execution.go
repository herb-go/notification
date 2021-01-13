package notificationqueue

import (
	"github.com/herb-go/notification"
)

//Execution notification Execution
type Execution struct {
	//ExecutionID execition id
	ExecutionID string
	//Notification notification to execute
	Notification *notification.Notification
	//RetryCount retry count
	RetryCount int32
	//StartTime execution start time
	StartTime int64
	//RetryAfterTime execution retry after time
	RetryAfterTime int64
}

//NewExecution create new execution
func NewExecution() *Execution {
	return &Execution{}
}
