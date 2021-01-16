package notification_test

import (
	"strconv"
	"sync"
	"sync/atomic"
	"testing"

	"github.com/herb-go/notification"
)

type testStore struct {
	locker sync.Mutex
	data   []*notification.Notification
}

func (d *testStore) Open() error {
	return nil
}
func (d *testStore) Close() error {
	return nil
}
func (d *testStore) Save(notification *notification.Notification) error {
	d.locker.Lock()
	defer d.locker.Unlock()
	for k := range d.data {
		if d.data[k].ID == notification.ID {
			d.data[k] = notification
			return nil
		}
	}
	d.data = append(d.data, notification)
	return nil
}
func (d *testStore) List(condition []*notification.Condition, iter string, asc bool, count int) (result []*notification.Notification, newiter string, err error) {
	d.locker.Lock()
	defer d.locker.Unlock()
	var start int
	var step int
	var end int
	var batch = ""
	for _, v := range condition {
		if v.Keyword == notification.ConditionBatch {
			batch = v.Value
		} else {
			return nil, "", notification.NewErrConditionNotSupported(v.Keyword)
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
func (d *testStore) Count(condition []*notification.Condition) (int, error) {
	d.locker.Lock()
	defer d.locker.Unlock()
	var batch = ""
	for _, v := range condition {
		if v.Keyword == notification.ConditionBatch {
			batch = v.Value
		} else {
			return 0, notification.NewErrConditionNotSupported(v.Keyword)
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
func (d *testStore) SupportedConditions() ([]string, error) {
	return []string{notification.ConditionBatch}, nil
}
func (d *testStore) Eject(id string) (*notification.Notification, error) {
	for k := range d.data {
		if d.data[k].ID == id {
			n := d.data[k]
			d.data = append(d.data[:k], d.data[k:]...)
			return n, nil
		}
	}
	return nil, notification.NewErrNotificationIDNotFound(id)
}

func newTestStore() *testStore {
	return &testStore{}
}

var current int64

func mustID() string {
	c := atomic.AddInt64(&current, 1)
	return strconv.FormatInt(c, 10)
}
func TestCondition(t *testing.T) {
	var store = newTestStore()
	var n *notification.Notification
	n = notification.New()
	n.Header.Set(notification.HeaderNameDraftMode, "1")
	n.ID = mustID()
	err := store.Save(n)
	if err != nil {
		t.Fatal(err)
	}
	n = notification.New()
	n.Header.Set(notification.HeaderNameDraftMode, "1")
	n.Header.Set(notification.HeaderNameBatch, "12345")
	err = store.Save(n)
	if err != nil {
		t.Fatal(err)
	}
	n = notification.New()
	n.Header.Set(notification.HeaderNameDraftMode, "1")
	n.Header.Set(notification.HeaderNameBatch, "12345")
	n.ID = mustID()
	err = store.Save(n)
	if err != nil {
		t.Fatal(err)
	}
	count, err := store.Count(nil)
	if count != 3 || err != nil {
		t.Fatal(count, err)
	}
	result, iter, err := store.List(nil, "", true, 0)
	if len(result) != 3 || iter != "" || err != nil {
		t.Fatal(result, iter, err)
	}
	result, iter, err = store.List(nil, "", true, 2)
	if len(result) != 2 || iter != "1" || err != nil {
		t.Fatal(result, iter, err)
	}
	result, iter, err = store.List(nil, "", false, 1)
	if len(result) != 1 || iter != "2" || err != nil {
		t.Fatal(result, iter, err)
	}
	result, iter, err = store.List(nil, "2", false, 1)
	if len(result) != 1 || iter != "1" || err != nil {
		t.Fatal(result, iter, err)
	}
	result, iter, err = store.List(nil, "1", false, 1)
	if len(result) != 1 || iter != "0" || err != nil {
		t.Fatal(result, iter, err)
	}
	result, iter, err = store.List(nil, "0", false, 1)
	if len(result) != 0 || iter != "" || err != nil {
		t.Fatal(result, iter, err)
	}
	cond := []*notification.Condition{&notification.Condition{
		Keyword: notification.ConditionBatch,
		Value:   "12345",
	}}
	count, err = store.Count(cond)
	if count != 2 || err != nil {
		t.Fatal(count, err)
	}
	result, iter, err = store.List(cond, "", true, 0)
	if len(result) != 2 || iter != "" || err != nil {
		t.Fatal(result, iter, err)
	}
	cond = []*notification.Condition{&notification.Condition{
		Keyword: notification.ConditionBatch,
		Value:   "notfound",
	}}
	count, err = store.Count(cond)
	if count != 0 || err != nil {
		t.Fatal(count, err)
	}
	result, iter, err = store.List(cond, "", true, 0)
	if len(result) != 0 || iter != "" || err != nil {
		t.Fatal(result, iter, err)
	}
	cond = []*notification.Condition{&notification.Condition{
		Keyword: "notfound",
		Value:   "notfound",
	}}
	count, err = store.Count(cond)
	if !notification.IsErrConditionNotSupported(err) {
		t.Fatal(result, iter, err)
	}
	result, iter, err = store.List(cond, "", true, 0)
	if !notification.IsErrConditionNotSupported(err) {
		t.Fatal(result, iter, err)
	}
}

func TestNopStore(t *testing.T) {
	var err error
	d := &notification.NopStore{}
	err = d.Open()
	if err != nil {
		t.Fatal(err)
	}
	err = d.Close()
	if err != nil {
		t.Fatal(err)
	}
	_, err = d.Count(nil)
	if err != notification.ErrStoreFeatureNotSupported {
		t.Fatal(err)
	}
	_, _, err = d.List(nil, "", true, 0)
	if err != notification.ErrStoreFeatureNotSupported {
		t.Fatal(err)
	}
	err = d.Save(notification.New())
	if err != notification.ErrStoreFeatureNotSupported {
		t.Fatal(err)
	}
	_, err = d.Remove("notexsit")
	if err != notification.ErrStoreFeatureNotSupported {
		t.Fatal(err)
	}
	_, err = d.SupportedConditions()
	if err != notification.ErrStoreFeatureNotSupported {
		t.Fatal(err)
	}
}
