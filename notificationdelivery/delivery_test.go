package notificationdelivery

import (
	"testing"
)

func TestDelivery(t *testing.T) {
	d := NewDeliveryServer()
	if d.Delivery != "" || d.Description != "" || d.DeliveryDriver != nil {
		t.Fatal(d)
	}
}
