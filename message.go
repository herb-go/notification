package notification

import "fmt"

type Message struct {
	Topic  Topic
	Target string
}

type Record struct {
	RecordID string
	Message
}

func (r *Record) String() string {
	return fmt.Sprintf("[%s] target %s @ topic %s", r.RecordID, r.Target, r.Topic)
}
