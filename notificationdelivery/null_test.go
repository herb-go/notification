package notificationdelivery

import (
	"testing"

	"github.com/herb-go/notification"
)

func TestNull(t *testing.T) {
	null, err := NewDriver(DeliveryNull, nil)
	if err != nil || null == nil {
		t.Fatal(null, err)
	}
	if null.DeliveryType() != DeliveryNull {
		t.Fatal(null)
	}

	i, err := null.CheckInvalidContent(notification.NewContent())
	if len(i) != 0 || err != nil {
		t.Fatal(i, err)
	}
	s, r, err := null.Deliver(notification.NewContent())
	if s != DeliveryStatusSuccess || r != "" || err != nil {
		t.Fatal(s, r, err)
	}
}
