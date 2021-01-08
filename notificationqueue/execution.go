package notificationqueue

import "github.com/herb-go/notification"

var ExecuteStatusFail = int32(0)
var ExecuteStatusSuccess = int32(1)
var ExecuteStatusAbort = int32(2)
var ExecuteStatusRetryTooMany = int32(3)

type Execution struct {
	ExecutionID    string
	Notification   *notification.Notification
	RetryCount     int32
	StartTime      int64
	RetryAfterTime int64
	ReturnReceipt  func(nid string, eid string, status int32, msg string) error
}

func (e *Execution) Success(receipt string) error {
	return e.ReturnReceipt(e.Notification.ID, e.ExecutionID, ExecuteStatusSuccess, receipt)
}

func (e *Execution) Fail(reason string) error {
	return e.ReturnReceipt(e.Notification.ID, e.ExecutionID, ExecuteStatusSuccess, reason)
}

func (e *Execution) Abort(reason string) error {
	return e.ReturnReceipt(e.Notification.ID, e.ExecutionID, ExecuteStatusSuccess, reason)
}
