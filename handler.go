package notification

import (
	"fmt"
	"log"
)

var Debug bool

type ErrorHandler func(error)

var NopErrorHandler = func(err error) {
	log.Println(err)
}

type RecordHandler func(*Record)

var NopRecordHandler = func(r *Record) {
	if Debug {
		fmt.Println(fmt.Sprintf("message record created: %s", r.String()))
	}
}

type NotificationHandler func(*Notification)

var NopNotificationHanlder = func(n *Notification) {
	if Debug {
		fmt.Println(fmt.Sprintf("notification created: %s", n.String()))
	}
}
