package notification

import "testing"

func TestDelivery(t *testing.T) {
	d := NewDeliveryServer()
	if d.Delivery != "" || d.Description != "" || d.Driver != nil {
		t.Fatal(d)
	}
}
