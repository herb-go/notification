package receiptstore

import (
	"time"

	"github.com/herb-go/notification"
	"github.com/herb-go/notification/notificationdelivery/notificationqueue"
)

//ReceiptStore receipt stroce interace
type ReceiptStore interface {
	//Open open store and return any error if raised
	Open() error
	//Close close store and return any error if raised
	Close() error
	//Save save given notificaiton to store.
	//Receipt with same notification id will be overwritten.
	Save(receipt *notificationqueue.Receipt) error
	//List list no more than count notifactions in store with given search conditions form start position .
	//Count should be greater than 0.
	//Found receipts and next list position iter will be returned.
	//Return largest id receipts if asc is false.
	List(condition []*notification.Condition, start string, asc bool, count int) (result []*notificationqueue.Receipt, iter string, err error)
	//Count count store with given search conditions
	Count(condition []*notification.Condition) (int, error)
	//SupportedConditions return supported condition keyword list
	SupportedConditions() ([]string, error)
	//LogRetentionPeriod log retention period.
	RetentionPeriod() (time.Duration, error)
	//Clear clear outdate log
	Clear() error
}
