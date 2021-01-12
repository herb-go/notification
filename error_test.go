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
	ok = IsInvalidContentError(err)
	if ok {
		t.Fatal(ok)
	}
	c := NewContent()
	c.Set("field1", "")
	c.Set("field3", "value3")
	err = CheckRequiredContentError(c, []string{"field1", "field2"})
	ok = IsInvalidContentError(err)
	if !ok {
		t.Fatal(ok)
	}
	msg = err.Error()
	if !strings.Contains(msg, "field1") || !strings.Contains(msg, "field2") || !strings.Contains(msg, "required") {
		t.Fatal(msg)
	}
	ok = IsInvalidContentError(errFoo)
	if ok {
		t.Fatal(ok)
	}
	err = CheckRequiredContentError(c, []string{"field3"})
	if err != nil {
		t.Fatal(err)
	}
	ok = IsInvalidContentError(err)
	if ok {
		t.Fatal(ok)
	}

	ok = IsErrNotificationIDNotFound(errFoo)
	if ok {
		t.Fatal(ok)
	}

	err = ErrNotificationIDNotFound("nid")
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
	err = ErrDeliveryNotFound("did")
	ok = IsErrDeliveryNotFound(nil)
	if ok {
		t.Fatal(ok)
	}
	ok = IsErrDeliveryNotFound(err)
	if !ok {
		t.Fatal(ok)
	}
	msg = err.Error()
	if !strings.Contains(msg, "did") || !strings.Contains(msg, "not found") {
		t.Fatal(msg)
	}

}
