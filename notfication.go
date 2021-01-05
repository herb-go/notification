package notification

import "fmt"

type Notification struct {
	ID          string
	Delivery    string
	ExpiredTime int64
	Header      Header
	Content     Content
}

func (n *Notification) String() string {
	return fmt.Sprintf("%s @ %s [ %s ]", n.ID, n.Delivery, n.Header.String())
}
