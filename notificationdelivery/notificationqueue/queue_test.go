package notificationqueue_test

import (
	"strconv"
	"sync"
	"testing"

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

func (q testQueue) PopChan() (<-chan *notificationqueue.Execution, error) {
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

//AttachTo attach queue to notifier
func (q testQueue) AttachTo(*notificationqueue.Notifier) error {
	return nil
}

//Detach detach queue.
func (q testQueue) Detach() error {
	return nil
}
func newTestQueue() testQueue {
	return make(testQueue)
}

func TestNopQueue(t *testing.T) {
	var err error
	n := &notificationqueue.NopQueue{}
	err = n.Start()
	if err != notificationqueue.ErrQueueDriverRequired {
		t.Fatal(err)
	}
	err = n.Stop()
	if err != notificationqueue.ErrQueueDriverRequired {
		t.Fatal(err)
	}
	_, err = n.PopChan()
	if err != notificationqueue.ErrQueueDriverRequired {
		t.Fatal(err)
	}
	err = n.Push(notification.New())
	if err != notificationqueue.ErrQueueDriverRequired {
		t.Fatal(err)
	}
	err = n.Remove("notexist")
	if err != notificationqueue.ErrQueueDriverRequired {
		t.Fatal(err)
	}
	err = n.AttachTo(notificationqueue.NewNotifier())
	if err != notificationqueue.ErrQueueDriverRequired {
		t.Fatal(err)
	}
	err = n.Detach()
	if err != notificationqueue.ErrQueueDriverRequired {
		t.Fatal(err)
	}
}
