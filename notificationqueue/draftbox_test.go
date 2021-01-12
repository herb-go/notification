package notificationqueue_test

import (
	"strconv"
	"sync"
	"testing"

	"github.com/herb-go/notification"
	"github.com/herb-go/notification/notificationqueue"
)

type testDraft struct {
	locker sync.Mutex
	data   []*notification.Notification
}

func (d *testDraft) Open() error {
	return nil
}
func (d *testDraft) Close() error {
	return nil
}
func (d *testDraft) Draft(notification *notification.Notification) error {
	locker.Lock()
	defer locker.Unlock()
	for k := range d.data {
		if d.data[k].ID == notification.ID {
			d.data[k] = notification
			return nil
		}
	}
	d.data = append(d.data, notification)
	return nil
}
func (d *testDraft) List(condition []*notificationqueue.Condition, iter string, asc bool, count int) (result []*notification.Notification, newiter string, err error) {
	locker.Lock()
	defer locker.Unlock()
	var start int
	var step int
	var end int
	var batch = ""
	for _, v := range condition {
		if v.Keyword == notificationqueue.ConditionBatch {
			batch = v.Value
		} else {
			return nil, "", notificationqueue.NewErrConditionNotSupported(v.Keyword)
		}
	}
	if asc {
		start = 0
		step = 1
		end = len(d.data)
	} else {
		start = len(d.data) - 1
		step = -1
		end = -1
	}
	result = []*notification.Notification{}
	var i = start
	iterpos, _ := strconv.Atoi(iter)
	for {
		var skiped bool
		if iter != "" {
			if asc {
				if iterpos >= i {
					skiped = true
				}
			} else {
				if iterpos <= i {
					skiped = true
				}
			}
		}
		if !skiped {
			data := d.data[i]
			if batch != "" {
				if data.Header.Get(notification.HeaderNameBatch) == batch {
					result = append(result, data)
				}
			} else {
				result = append(result, data)
			}
			if count > 0 && len(result) == count {
				return result, strconv.Itoa(i), nil
			}
		}
		i = i + step
		if i == end {
			break
		}
	}
	return result, "", nil
}
func (d *testDraft) Count(condition []*notificationqueue.Condition) (int, error) {
	locker.Lock()
	defer locker.Unlock()
	var batch = ""
	for _, v := range condition {
		if v.Keyword == notificationqueue.ConditionBatch {
			batch = v.Value
		} else {
			return 0, notificationqueue.NewErrConditionNotSupported(v.Keyword)
		}
	}
	var count int
	for k := range d.data {
		if batch == "" {
			count = count + 1
		} else {
			if d.data[k].Header.Get(notification.HeaderNameBatch) == batch {
				count = count + 1
			}
		}
	}
	return count, nil
}
func (d *testDraft) SupportedConditions() ([]string, error) {
	return []string{notificationqueue.ConditionBatch}, nil
}
func (d *testDraft) Eject(id string) (*notification.Notification, error) {
	for k := range d.data {
		if d.data[k].ID == id {
			n := d.data[k]
			d.data = append(d.data[:k], d.data[k:]...)
			return n, nil
		}
	}
	return nil, notification.NewErrNotificationIDNotFound(id)
}

func newTestDraft() *testDraft {
	return &testDraft{}
}

func TestCondition(t *testing.T) {
	var n *notification.Notification
	initLog()
	p := newTestPublisher()
	err := p.Start()
	if err != nil {
		panic(err)
	}
	defer func() {
		err := p.Stop()
		if err != nil {
			panic(err)
		}
	}()
	n = notification.New()
	n.Header.Set(notification.HeaderNameDraftMode, "1")
	n.ID = mustID()
	ok, err := p.PublishNotification(n)
	if ok || err != nil {
		t.Fatal(ok, err)
	}
	n = notification.New()
	n.Header.Set(notification.HeaderNameDraftMode, "1")
	n.Header.Set(notification.HeaderNameBatch, "12345")
	n.ID = mustID()
	ok, err = p.PublishNotification(n)
	if ok || err != nil {
		t.Fatal(ok, err)
	}
	n = notification.New()
	n.Header.Set(notification.HeaderNameDraftMode, "1")
	n.Header.Set(notification.HeaderNameBatch, "12345")
	n.ID = mustID()
	ok, err = p.PublishNotification(n)
	if ok || err != nil {
		t.Fatal(ok, err)
	}
	count, err := p.Draftbox.Count(nil)
	if count != 3 || err != nil {
		t.Fatal(count, err)
	}
	result, iter, err := p.Draftbox.List(nil, "", true, 0)
	if len(result) != 3 || iter != "" || err != nil {
		t.Fatal(result, iter, err)
	}
	result, iter, err = p.Draftbox.List(nil, "", true, 2)
	if len(result) != 2 || iter != "1" || err != nil {
		t.Fatal(result, iter, err)
	}
	result, iter, err = p.Draftbox.List(nil, "", false, 1)
	if len(result) != 1 || iter != "2" || err != nil {
		t.Fatal(result, iter, err)
	}
	result, iter, err = p.Draftbox.List(nil, "2", false, 1)
	if len(result) != 1 || iter != "1" || err != nil {
		t.Fatal(result, iter, err)
	}
	result, iter, err = p.Draftbox.List(nil, "1", false, 1)
	if len(result) != 1 || iter != "0" || err != nil {
		t.Fatal(result, iter, err)
	}
	result, iter, err = p.Draftbox.List(nil, "0", false, 1)
	if len(result) != 0 || iter != "" || err != nil {
		t.Fatal(result, iter, err)
	}
	cond := []*notificationqueue.Condition{&notificationqueue.Condition{
		Keyword: notificationqueue.ConditionBatch,
		Value:   "12345",
	}}
	count, err = p.Draftbox.Count(cond)
	if count != 2 || err != nil {
		t.Fatal(count, err)
	}
	result, iter, err = p.Draftbox.List(cond, "", true, 0)
	if len(result) != 2 || iter != "" || err != nil {
		t.Fatal(result, iter, err)
	}
	cond = []*notificationqueue.Condition{&notificationqueue.Condition{
		Keyword: notificationqueue.ConditionBatch,
		Value:   "notfound",
	}}
	count, err = p.Draftbox.Count(cond)
	if count != 0 || err != nil {
		t.Fatal(count, err)
	}
	result, iter, err = p.Draftbox.List(cond, "", true, 0)
	if len(result) != 0 || iter != "" || err != nil {
		t.Fatal(result, iter, err)
	}
	cond = []*notificationqueue.Condition{&notificationqueue.Condition{
		Keyword: "notfound",
		Value:   "notfound",
	}}
	count, err = p.Draftbox.Count(cond)
	if !notificationqueue.IsErrConditionNotSupported(err) {
		t.Fatal(result, iter, err)
	}
	result, iter, err = p.Draftbox.List(cond, "", true, 0)
	if !notificationqueue.IsErrConditionNotSupported(err) {
		t.Fatal(result, iter, err)
	}
}
