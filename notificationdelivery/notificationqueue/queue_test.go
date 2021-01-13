package notificationqueue_test

import (
	"strconv"
	"sync"

	"github.com/herb-go/notification"
	"github.com/herb-go/notification/notificationdelivery/notificationqueue"
)

var locker sync.Mutex

var currentid int

func testIDGenerator() (string, error) {
	locker.Lock()
	defer locker.Unlock()
	currentid = currentid + 1
	return strconv.Itoa(currentid), nil
}

func mustID() string {
	id, _ := testIDGenerator()
	return id
}

type testQueue chan *notificationqueue.Execution

func (q testQueue) PopChan() (chan *notificationqueue.Execution, error) {
	return q, nil
}
func (q testQueue) Push(n *notification.Notification) error {
	e := notificationqueue.NewExecution()
	e.ExecutionID = mustID()
	e.Notification = n
	go func() {
		q <- e
	}()
	return nil
}
func (q testQueue) Remove(nid string) error {
	return nil
}
func (q testQueue) Start() error {
	return nil
}
func (q testQueue) Stop() error {
	return nil
}

func newTestQueue() testQueue {
	return make(testQueue)
}
