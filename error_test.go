package notification

import (
	"errors"
	"strings"
	"testing"
)

var errFoo = errors.New("nop")

func TestError(t *testing.T) {
	var err error
	var ok bool
	var msg string
	ok = IsErrInvalidContent(err)
	if ok {
		t.Fatal(ok)
	}
	c := NewContent()
	c.Set("field1", "")
	c.Set("field3", "value3")
	err = CheckRequiredContentError(c, []string{"field1", "field2"})
	ok = IsErrInvalidContent(err)
	if !ok {
		t.Fatal(ok)
	}
	msg = err.Error()
	if !strings.Contains(msg, "field1") || !strings.Contains(msg, "field2") || !strings.Contains(msg, "required") {
		t.Fatal(msg)
	}
	ok = IsErrInvalidContent(errFoo)
	if ok {
		t.Fatal(ok)
	}
	err = CheckRequiredContentError(c, []string{"field3"})
	if err != nil {
		t.Fatal(err)
	}
	ok = IsErrInvalidContent(err)
	if ok {
		t.Fatal(ok)
	}

	ok = IsErrNotificationIDNotFound(errFoo)
	if ok {
		t.Fatal(ok)
	}

	err = NewErrNotificationIDNotFound("nid")
	ok = IsErrNotificationIDNotFound(nil)
	if ok {
		t.Fatal(ok)
	}
	ok = IsErrNotificationIDNotFound(err)
	if !ok {
		t.Fatal(ok)
	}
	msg = err.Error()
	if !strings.Contains(msg, "nid") || !strings.Contains(msg, "not found") {
		t.Fatal(msg)
	}

}

func TestConditionError(t *testing.T) {
	var err error
	var ok bool
	var msg string
	ok = IsErrConditionNotSupported(nil)
	if ok {
		t.Fatal(ok)
	}
	ok = IsErrConditionNotSupported(errFoo)
	if ok {
		t.Fatal(ok)
	}
	err = NewErrConditionNotSupported("test")
	ok = IsErrConditionNotSupported(err)
	if !ok {
		t.Fatal(ok)
	}
	msg = err.Error()
	if !strings.Contains(msg, "test") || !strings.Contains(msg, "not supported") {
		t.Fatal(msg)
	}
	ok = IsErrInvalidConditionValue(nil)
	if ok {
		t.Fatal(ok)
	}

	err = NewErrInvalidConditionValue("testvalue")
	ok = IsErrInvalidConditionValue(err)
	if !ok {
		t.Fatal(ok)
	}
	msg = err.Error()
	if !strings.Contains(msg, "testvalue") || !strings.Contains(msg, "value invalid") {
		t.Fatal(msg)
	}
	ok = IsErrConditionNotSupported(err)
	if ok {
		t.Fatal(ok)
	}
}
