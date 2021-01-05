package notificationmessage

import (
	"fmt"

	"github.com/herb-go/notification"
)

type Message struct {
	Topic  Topic
	Header notification.Header
	Model  Model
}

type Record struct {
	RecordID    string
	CreatedTime int64
	Message
}

func (r *Record) String() string {
	return fmt.Sprintf("[%s] target %s @ topic %s", r.RecordID)
}
