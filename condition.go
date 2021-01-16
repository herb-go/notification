package notification

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
