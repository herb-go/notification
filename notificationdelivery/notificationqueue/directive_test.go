package notificationqueue

import "testing"

var testDirective = DirectiveFunc(func(*Publisher) error {
	return nil
})

var testFactory = func(loader func(v interface{}) error) (Directive, error) {
	return testDirective, nil
}

func registerTestFactory() {
	Register("testfactory", testFactory)
}
func TestFactory(t *testing.T) {
	defer func() {
		UnregisterAll()
	}()
	UnregisterAll()
	fs := Factories()
	if len(fs) != 0 {
		t.Fatal(fs)
	}
	registerTestFactory()
	fs = Factories()

	if len(fs) != 1 {
		t.Fatal(fs)
	}
}

func TestNilFactory(t *testing.T) {
	defer func() {
		UnregisterAll()
	}()
	registerTestFactory()
	defer func() {
		r := recover()
		if r == nil {
			t.Fatal(r)
		}
	}()
	Register("test", nil)
}

func TestDupFactory(t *testing.T) {
	defer func() {
		UnregisterAll()
	}()
	registerTestFactory()
	defer func() {
		r := recover()
		if r == nil {
			t.Fatal(r)
		}
	}()
	registerTestFactory()
}

func TestUnknownFactory(t *testing.T) {
	d, err := NewDirective("unknown", nil)
	if d != nil || err == nil {
		t.Fatal(d, err)
	}
}
