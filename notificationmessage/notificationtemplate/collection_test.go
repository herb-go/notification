package notificationtemplate

import "testing"

func TestCollection(t *testing.T) {
	c := NewCollection()
	c.Set("Test", "testvalue")
	if c.Get("test") != "testvalue" || c.Get("tEst") != "testvalue" {
		t.Fatal(c)
	}
}
