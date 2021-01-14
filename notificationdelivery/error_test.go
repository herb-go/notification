package notificationdelivery

import (
	"strings"
	"testing"
)

func TestError(t *testing.T) {
	var err error
	var ok bool
	var msg string
	err = NewErrDeliveryNotFound("did")
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
