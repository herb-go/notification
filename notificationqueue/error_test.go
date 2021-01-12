package notificationqueue

import (
	"errors"
	"strings"
	"testing"
)

var errFoo = errors.New("foo")

func TestError(t *testing.T) {
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
	err = ErrConditionNotSupported("test")
	ok = IsErrConditionNotSupported(err)
	if !ok {
		t.Fatal(ok)
	}
	msg = err.Error()
	if !strings.Contains(msg, "test") || !strings.Contains(msg, "not supported") {
		t.Fatal(msg)
	}
}
