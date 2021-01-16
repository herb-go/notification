package notificationqueue

import (
	"errors"
	"testing"
)

var errFoo = errors.New("foo")

func TestNopIDGenerator(t *testing.T) {
	s, err := NopIDGenerator()
	if s != "" || err != ErrIDGeneratorRequired {
		t.Fatal(s, err)
	}
}
