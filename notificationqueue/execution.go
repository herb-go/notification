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
	return e.queue.ReturnSuccessReceipt(e.Notification.ID, e.ExecutionID, receipt)
}

func (e *Execution) Fail(reason string) error {
	return e.queue.ReturnFailReceipt(e.Notification.ID, e.ExecutionID, reason)
}

func (e *Execution) Abort(reason string) error {
	return e.queue.ReturnFailReceipt(e.Notification.ID, e.ExecutionID, reason)
}
