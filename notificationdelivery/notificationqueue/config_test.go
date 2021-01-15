package notificationqueue

import "testing"

func TestConfig(t *testing.T) {
	defer func() {
		UnregisterAll()
		registerTestFactory()
	}()
	registerTestFactory()
	p := NewPublisher()
	c := &Config{
		Directives: []*DirectiveConfig{
			&DirectiveConfig{
				Directive:       "testfactory",
				DirectiveConfig: nil,
			},
		},
	}
	err := c.ApplyTo(p)
	if err != nil {
		panic(err)
	}
}
