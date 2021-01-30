package notificationrender

import (
	"strings"
	"testing"
)

func TestError(t *testing.T) {
	var err error
	var ok bool
	var msg string
	err = NewErrRendererNotFound("rid")
	ok = IsErrRendererNotFound(nil)
	if ok {
		t.Fatal(ok)
	}
	ok = IsErrRendererNotFound(err)
	if !ok {
		t.Fatal(ok)
	}
	msg = err.Error()
	if !strings.Contains(msg, "rid") || !strings.Contains(msg, "not found") {
		t.Fatal(msg)
	}
}
