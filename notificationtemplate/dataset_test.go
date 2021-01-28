package notificationtemplate

import "testing"

func TestDataset(t *testing.T) {
	c := NewDataset()
	c.Set("Test", "testvalue")
	if c.Get("test") != "testvalue" || c.Get("tEst") != "testvalue" {
		t.Fatal(c)
	}
}
