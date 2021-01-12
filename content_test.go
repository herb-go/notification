package notification

import (
	"encoding/json"
	"testing"
)

func TestContent(t *testing.T) {
	c := NewContent()
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
	js := c.MustJSON()
	c2 := NewContent()
	err := json.Unmarshal([]byte(js), &c2)
	if err != nil {
		panic(err)
	}
	if c2.Get("test1") != "testvalue" || c2.Get("test2") != "testvalue2" {
		t.Fatal(c)
	}
}
