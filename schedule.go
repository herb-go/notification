package notification

type Schedule struct {
	DeliveryTime int64
	TTLInSecond  int64
}

type NotificationScheduler interface {
	LoadNotificationSchedule(*Message) (*Schedule, error)
}

func NewSchedule() *Schedule {
	return &Schedule{}
}
