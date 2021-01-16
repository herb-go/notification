package notification

const (
	//ConditionTopic store serach condition keyword for notification topic
	ConditionTopic = "topic"

	//ConditionNotificationID store serach condition keyword for notification id
	ConditionNotificationID = "notificationid"

	//ConditionTarget store serach condition keyword for notification target
	ConditionTarget = "target"

	//ConditionBatch store serach condition keyword for notification batch id
	ConditionBatch = "batch"

	//ConditionInContent store serach condition keyword for text in notification content
	ConditionInContent = "incontent"

	//ConditionSender store serach condition keyword for text in notification delivery keyword
	ConditionSender = "sender"

	//ConditionDelivery store serach condition keyword for text in notification delivery keyword
	ConditionDelivery = "delivery"

	//ConditionBeforeTimestamp store serach condition keyword for notification created-time before given timestamp
	ConditionBeforeTimestamp = "beforetimestamp"

	//ConditionAfterTimestamp store serach condition keyword for notification created-time after given timestamp
	ConditionAfterTimestamp = "aftertimestamp"
)

//Condition store search condition
type Condition struct {
	//Keyword condition keyword
	Keyword string
	//Value condition value to filter notification
	Value string
}
