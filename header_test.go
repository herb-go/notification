package notification

import (
	"net/url"
	"testing"
)

func TestHeader(t *testing.T) {
	h := NewHeader()
	h.Set("Test1", "testvalue")
	if h.Get("test1") != "testvalue" || h.Get("TEST1") != "testvalue" {
		t.Fatal(h)
	}
	h.Set("tEst1", "testvalue")
	if h.Get("test1") != "testvalue" || h.Get("TEST1") != "testvalue" {
		t.Fatal(h)
	}
	h.Set("test2", "testvalue2")
	if h.Get("test1") != "testvalue" || h.Get("test2") != "testvalue2" {
		t.Fatal(h)
	}
	enc := h.String()
	data, err := url.ParseQuery(enc)
	if err != nil {
		panic(err)
	}
	h2 := NewContent()
	for k := range data {
		h2.Set(k, data.Get(k))
	}

	if h2.Get("test1") != "testvalue" || h2.Get("test2") != "testvalue2" {
		t.Fatal(h2)
	}
}
