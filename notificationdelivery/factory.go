package notificationdelivery

import (
	"errors"
	"fmt"
	"sort"
	"sync"
)

//Factory delivery driver factory
type Factory func(loader func(v interface{}) error) (DeliveryDriver, error)

var (
	factorysMu sync.RWMutex
	factories  = make(map[string]Factory)
)

// Register makes a driver creator available by the provided name.
// If Register is called twice with the same name or if driver is nil,
// it panics.
func Register(name string, f Factory) {
	factorysMu.Lock()
	defer factorysMu.Unlock()
	if f == nil {
		panic(errors.New("delivery: Register delivery factory is nil"))
	}
	if _, dup := factories[name]; dup {
		panic(errors.New("delivery: Register called twice for factory " + name))
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

//NewDriver create new driver with given name loader.
//Reutrn driver created and any error if raised.
func NewDriver(name string, loader func(v interface{}) error) (DeliveryDriver, error) {
	factorysMu.RLock()
	factoryi, ok := factories[name]
	factorysMu.RUnlock()
	if !ok {
		return nil, fmt.Errorf("delivery: unknown driver %q (forgotten import?)", name)
	}
	return factoryi(loader)
}
