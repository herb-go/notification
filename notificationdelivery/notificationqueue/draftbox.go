package notificationqueue

import (
	"github.com/herb-go/notification"
)

//DefaultDraftboxListLimit default draftbox list limit
const DefaultDraftboxListLimit = 10
const (
	//ConditionTopic draftbox serach condition keyword for notification topic
	ConditionTopic = "topic"

	//ConditionNotificationID draftbox serach condition keyword for notification id
	ConditionNotificationID = "notificationid"

	//ConditionTarget draftbox serach condition keyword for notification target
	ConditionTarget = "target"

	//ConditionBatch draftbox serach condition keyword for notification batch id
	ConditionBatch = "batch"

	//ConditionInContent draftbox serach condition keyword for text in notification content
	ConditionInContent = "incontent"

	//ConditionDelivery draftbox serach condition keyword for text in notification delivery keyword
	ConditionDelivery = "delivery"

	//ConditionBeforeTimestamp draftbox serach condition keyword for notification created-time before given timestamp
	ConditionBeforeTimestamp = "beforetimestamp"

	//ConditionAfterTimestamp draftbox serach condition keyword for notification created-time after given timestamp
	ConditionAfterTimestamp = "aftertimestamp"
)

//Condition draftbox search condition
type Condition struct {
	//Keyword condition keyword
	Keyword string
	//Value condition value to filter notification
	Value string
}

//Draftbox notification draftbox interface
type Draftbox interface {
	//Open open draftbox and return any error if raised
	Open() error
	//Close close draftbox and return any error if raised
	Close() error
	//Draft save given notificaiton to draft box.
	//Notification with same id will be overwritten.
	Draft(notification *notification.Notification) error
	//List list no more than count notifactions in draftbox with given search conditions form start position .
	//Count should be greater than 0.
	//Found notifications and next list position iter will be returned.
	//Return largest id notification if asc is false.
	List(condition []*Condition, start string, asc bool, count int) (result []*notification.Notification, iter string, err error)
	//Count draft box with given search conditions
	Count(condition []*Condition) (int, error)
	//SupportedConditions return supported condition keyword list
	SupportedConditions() ([]string, error)
	//Eject remove notification by given id and return removed notification.
	Eject(id string) (*notification.Notification, error)
}

var CheckerDraftModeHeader = notification.HasHeaderChecker(notification.HeaderNameDraftMode)

//NopDraftbox nop draft box
type NopDraftbox struct{}

//Open open draftbox and return any error if raised
func (b *NopDraftbox) Open() error {
	return ErrDraftBoxRequired
}

//Close close draftbox and return any error if raised
func (b *NopDraftbox) Close() error {
	return ErrDraftBoxRequired
}

//Draft save given notificaiton to draft box.
//Notification with same id will be overwritten.
func (b *NopDraftbox) Draft(notification *notification.Notification) error {
	return ErrDraftBoxRequired
}

//List list no more than count notifactions in draftbox with given search conditions form start position .
//Count should be greater than 0.
//Found notifications and next list position iter will be returned.
//Return largest id notification if asc is false.
func (b *NopDraftbox) List(condition []*Condition, start string, asc bool, count int) (result []*notification.Notification, iter string, err error) {
	return nil, "", ErrDraftBoxRequired
}

//Count draft box with given search conditions
func (b *NopDraftbox) Count(condition []*Condition) (int, error) {
	return 0, ErrDraftBoxRequired
}

//SupportedConditions return supported condition keyword list
func (b *NopDraftbox) SupportedConditions() ([]string, error) {
	return nil, ErrDraftBoxRequired
}

//Eject remove notification by given id and return removed notification.
func (b *NopDraftbox) Eject(id string) (*notification.Notification, error) {
	return nil, ErrDraftBoxRequired
}
