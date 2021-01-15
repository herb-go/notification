package notificationqueue

import "github.com/herb-go/notification"

//NopExecutionHandler nop execution handelr
func NopExecutionHandler(e *Execution) {}

//NopRecover noop recover
func NopRecover() {}

//NopNotificationHandler nop notification heandler
func NopNotificationHandler(n *notification.Notification) {}

//NopReceiptHanlder nop receipt hanlder
func NopReceiptHanlder(*Receipt) {}

//NopIDGenerator nop id gererator
func NopIDGenerator() (string, error) {
	return "", ErrIDGeneratorRequired
}
