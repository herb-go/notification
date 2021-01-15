package notificationqueue

import (
	"errors"
	"fmt"
	"sort"
	"sync"
)

//Directive publisher directive
type Directive interface {
	//AppylToPublisher applu directive to publisher
	AppylToPublisher(*Publisher) error
}

//DirectiveFunc directive func
type DirectiveFunc func(*Publisher) error

//AppylToPublisher applu directive to publisher
func (f DirectiveFunc) AppylToPublisher(p *Publisher) error {
	return f(p)
}

//Factory directive factory
type Factory func(loader func(v interface{}) error) (Directive, error)

var (
	factorysMu sync.RWMutex
	factories  = make(map[string]Factory)
)

// Register makes a directive factory available by the provided name.
// If Register is called twice with the same name or if driver is nil,
// it panics.
func Register(name string, f Factory) {
	factorysMu.Lock()
	defer factorysMu.Unlock()
	if f == nil {
		panic(errors.New("notificationqueue: Register directive factory is nil"))
	}
	if _, dup := factories[name]; dup {
		panic(errors.New("notificationqueue: Register called twice for factory " + name))
	}
	factories[name] = f
}

//UnregisterAll unregister all driver
func UnregisterAll() {
	factorysMu.Lock()
	defer factorysMu.Unlock()
	// For tests.
	factories = make(map[string]Factory)
}

//Factories returns a sorted list of the names of the registered factories.
func Factories() []string {
	factorysMu.RLock()
	defer factorysMu.RUnlock()
	var list []string
	for name := range factories {
		list = append(list, name)
	}
	sort.Strings(list)
	return list
}

//NewDirective create new directive with given name loader.
//Reutrn directive created and any error if raised.
func NewDirective(name string, loader func(v interface{}) error) (Directive, error) {
	factorysMu.RLock()
	factoryi, ok := factories[name]
	factorysMu.RUnlock()
	if !ok {
		return nil, fmt.Errorf("notificationqueue: unknown directive %q (forgotten import?)", name)
	}
	return factoryi(loader)
}
