package notification

import (
	"strings"
	"testing"
)

func TestNofication(t *testing.T) {
	n := New()
	n.ID = "nid"
	n.Header.Set("header1", "headervalue1")
	n.Header.Set("Header2", "headervalue2")
	n.Content.Set("content1", "contentvalue1")
	msg := n.String()
	if !strings.Contains(msg, "nid") || !strings.Contains(msg, "headervalue1") || !strings.Contains(msg, "headervalue2") ||
		strings.Contains(msg, "contentvalue1") || strings.Contains(msg, "content1") {
		t.Fatal(msg)
	}
}
