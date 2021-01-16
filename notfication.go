package notification

import (
	"fmt"
	"time"
)

//SuggestedNotificationTTL suggrested notification ttl
var SuggestedNotificationTTL = 24 * time.Hour

//Notification notification struct
type Notification struct {
	//ID notification id
	ID string
	//Delivery notification delivery keyword
	Delivery string
	//CreatedTime notification created unix timestmamp
	CreatedTime int64
	//ExpiredTime notification expired unix timestamp
	//ExpiredTime 0 or less than 0 means never expired
	ExpiredTime int64
	//Header notification header
	Header Header
	//Content notification content
	Content Content
}

//String return notification info in string format
func (n *Notification) String() string {
	return fmt.Sprintf("%s@%s [ %s ]", n.ID, n.Delivery, n.Header.String())
}

//New create new notification
func New() *Notification {
	return &Notification{
		Header:  Header{},
		Content: Content{},
	}
}
