package notificationdelivery

import (
	"testing"
	"time"

	"github.com/herb-go/notification"
)

func TestConfig(t *testing.T) {
	c := &DeliveryCenterConfig{
		&DeliveryServerConfig{
			Delivery:    "test",
			Description: "test desc",
			Config: Config{
				DeliveryType: "null",
			},
		},
		&DeliveryServerConfig{
			Delivery:    "disabled",
			Description: "disabled desc",
			Disabled:    true,
			Config: Config{
				DeliveryType: "null",
			},
		},
	}
	dc, err := c.CreateDeliveryCenter()
	if dc == nil || err != nil {
		t.Fatal(dc, err)
	}
	n := notification.New()
	n.Delivery = "test"
	s, r, err := DeliverNotification(dc, n)
	if s != DeliveryStatusSuccess || r != "" || err != nil {
		t.Fatal(s, r, err)
	}
	n = notification.New()
	n.Delivery = "disabled"
	s, r, err = DeliverNotification(dc, n)
	if s != DeliveryStatusDisabled || r != "" || err != nil {
		t.Fatal(s, r, err)
	}
	n = notification.New()
	n.Delivery = "test"
	n.ExpiredTime = time.Now().Add(-time.Hour).Unix()
	s, r, err = DeliverNotification(dc, n)
	if s != DeliveryStatusExpired || r != "" || err != nil {
		t.Fatal(s, r, err)
	}
}
