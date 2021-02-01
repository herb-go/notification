package notificationview

import (
	"testing"
)

func TestMessage(t *testing.T) {
	c := NewMessage()
	c.Set("Test1", "testvalue")
	if c.Get("test1") != "testvalue" || c.Get("TEST1") != "testvalue" {
		t.Fatal(c)
	}
	c.Set("tEst1", "testvalue")
	if c.Get("test1") != "testvalue" || c.Get("TEST1") != "testvalue" {
		t.Fatal(c)
	}
	c.Set("test2", "testvalue2")
	if c.Get("test1") != "testvalue" || c.Get("test2") != "testvalue2" {
		t.Fatal(c)
	}

}
