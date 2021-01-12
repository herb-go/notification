package notificationqueue

import (
	"testing"

	"github.com/herb-go/notification"
)

func TestFilter(t *testing.T) {
	var n *notification.Notification
	var err error
	var ok bool
	var cond []*Condition
	f := NewFilter()
	if f.BatchID != "" || f.Before != 0 || f.After != 0 || f.Delivery != "" || f.InContent != "" || f.NotificationID != "" || f.Target != "" || f.Topic != "" {
		t.Fatal(f)
	}
	n = notification.New()
	ok, err = f.FilterNotification(n, 100000)
	if !ok || err != nil {
		t.Fatal(ok, err)
	}
	cond = []*Condition{&Condition{Keyword: "unknown"}}
	err = ApplyToFilter(f, cond)
	if !IsErrConditionNotSupported(err) || err.(*ErrConditionNotSupported).Condition != "unknown" {
		t.Fatal(err)
	}
	f = NewFilter()
	err = ApplyToFilter(f, []*Condition{&Condition{Keyword: ConditionAfterTimestamp, Value: "a100000"}})
	if !IsErrInvalidConditionValue(err) || err.(*ErrInvalidConditionValue).Condition != ConditionAfterTimestamp {
		t.Fatal(err)
	}
	err = ApplyToFilter(f, []*Condition{&Condition{Keyword: ConditionBeforeTimestamp, Value: "a100000"}})
	if !IsErrInvalidConditionValue(err) || err.(*ErrInvalidConditionValue).Condition != ConditionBeforeTimestamp {
		t.Fatal(err)
	}
	f = NewFilter()
	err = ApplyToFilter(f, []*Condition{&Condition{Keyword: ConditionAfterTimestamp, Value: "100000"}})
	if err != nil {
		t.Fatal(err)
	}
	if f.After != 100000 {
		t.Fatal(f)
	}
	n = notification.New()
	ok, err = f.FilterNotification(n, 99999)
	if !ok || err != nil {
		t.Fatal(ok, err)
	}
	ok, err = f.FilterNotification(n, 100000)
	if ok || err != nil {
		t.Fatal(ok, err)
	}
	ok, err = f.FilterNotification(n, 100001)
	if ok || err != nil {
		t.Fatal(ok, err)
	}

	f = NewFilter()
	err = ApplyToFilter(f, []*Condition{&Condition{Keyword: ConditionBeforeTimestamp, Value: "100000"}})
	if err != nil {
		t.Fatal(err)
	}
	if f.Before != 100000 {
		t.Fatal(f)
	}
	n = notification.New()
	ok, err = f.FilterNotification(n, 99999)
	if ok || err != nil {
		t.Fatal(ok, err)
	}
	ok, err = f.FilterNotification(n, 100000)
	if ok || err != nil {
		t.Fatal(ok, err)
	}
	ok, err = f.FilterNotification(n, 100001)
	if !ok || err != nil {
		t.Fatal(ok, err)
	}

	f = NewFilter()
	err = ApplyToFilter(f, []*Condition{&Condition{Keyword: ConditionInContent, Value: "searchfor"}})
	if err != nil {
		t.Fatal(err)
	}
	if f.InContent != "searchfor" {
		t.Fatal(f)
	}
	n = notification.New()
	ok, err = f.FilterNotification(n, 100000)
	if ok || err != nil {
		t.Fatal(ok, err)
	}
	n.Content.Set("content1", "value1")
	ok, err = f.FilterNotification(n, 100000)
	if ok || err != nil {
		t.Fatal(ok, err)
	}
	n.Content.Set("content1", "searchfor")
	ok, err = f.FilterNotification(n, 100000)
	if !ok || err != nil {
		t.Fatal(ok, err)
	}
	n.Content.Set("content1", "asearchfora")
	ok, err = f.FilterNotification(n, 100000)
	if !ok || err != nil {
		t.Fatal(ok, err)
	}
	n.Content.Set("content2", "asearchfora")
	ok, err = f.FilterNotification(n, 100000)
	if !ok || err != nil {
		t.Fatal(ok, err)
	}
}
