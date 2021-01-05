package notification

import (
	"log"
)

var Debug bool

type ErrorHandler func(error)

var NopErrorHandler = func(err error) {
	log.Println(err)
}

type NotificationHandler func(*Notification)

var NopNotificationHanlder = func(n *Notification) {
}
