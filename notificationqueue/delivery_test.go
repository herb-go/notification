package notificationqueue_test

import (
	"sync"

	"github.com/herb-go/notification"
)

type testDelivery struct {
	locker sync.Mutex
	data   []*notification.Notification
}

func newTestDelivery() *testDelivery {
	return &testDelivery{}
}
