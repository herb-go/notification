package notificationqueue

import "github.com/herb-go/notification"

var ExecuteStatusFail = int32(0)
var ExecuteStatusSuccess = int32(1)
var ExecuteStatusAbort = int32(2)

type Execution struct {
	ExecutionID    string
	Notification   *notification.Notification
	RetryCount     int32
	StartTime      int64
	RetryAfterTime int64
	queue          Queue
}

func (e *Execution) Success(receipt string) error {
	return e.queue.ReturnReceipt(e.Notification.ID, e.ExecutionID, ExecuteStatusSuccess, receipt)
}

func (e *Execution) Fail(reason string) error {
	return e.queue.ReturnReceipt(e.Notification.ID, e.ExecutionID, ExecuteStatusSuccess, reason)
}

func (e *Execution) Abort(reason string) error {
	return e.queue.ReturnReceipt(e.Notification.ID, e.ExecutionID, ExecuteStatusSuccess, reason)
}
