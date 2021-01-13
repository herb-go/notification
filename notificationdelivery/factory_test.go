package notificationdelivery

import "testing"

func TestFactory(t *testing.T) {
	defer func() {
		UnregisterAll()
		RegisterNullFactory()
	}()
	UnregisterAll()
	fs := Factories()
	if len(fs) != 0 {
		t.Fatal(fs)
	}
	RegisterNullFactory()
	fs = Factories()

	if len(fs) != 1 {
		t.Fatal(fs)
	}
}

func TestNilFactory(t *testing.T) {
	defer func() {
		UnregisterAll()
		RegisterNullFactory()
	}()
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
		RegisterNullFactory()
	}()
	defer func() {
		r := recover()
		if r == nil {
			t.Fatal(r)
		}
	}()
	RegisterNullFactory()
}

func TestUnknownFactory(t *testing.T) {
	d, err := NewDriver("unknown", nil)
	if d != nil || err == nil {
		t.Fatal(d, err)
	}
}
