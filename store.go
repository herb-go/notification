package notification

//DefaultStoreListLimit default store list limit
const DefaultStoreListLimit = 10

//Store notitication store interface
type Store interface {
	//Open open store and return any error if raised
	Open() error
	//Close close store and return any error if raised
	Close() error
	//Save save given notificaiton to store.
	//Notification with same id will be overwritten.
	Save(notification *Notification) error
	//List list no more than count notifactions in store with given search conditions form start position .
	//Count should be greater than 0.
	//Found notifications and next list position iter will be returned.
	//Return largest id notification if asc is false.
	List(condition []*Condition, start string, asc bool, count int) (result []*Notification, iter string, err error)
	//Count count store with given search conditions
	Count(condition []*Condition) (int, error)
	//SupportedConditions return supported condition keyword list
	SupportedConditions() ([]string, error)
	//Remove remove notification by given id and return removed notification.
	Remove(id string) (*Notification, error)
}

//Searchable searchable interface
type Searchable interface {
	//Count count store with given search conditions
	Count(condition []*Condition) (int, error)
	//SupportedConditions return supported condition keyword list
	SupportedConditions() ([]string, error)
}

type NopStore struct{}

//Open open store and return any error if raised
func (n NopStore) Open() error {
	return nil
}

//Close close store and return any error if raised
func (n NopStore) Close() error {
	return nil
}

//Save save given notificaiton to store.
//Notification with same id will be overwritten.
func (n NopStore) Save(notification *Notification) error {
	return ErrStoreFeatureNotSupported
}

//List list no more than count notifactions in store with given search conditions form start position .
//Count should be greater than 0.
//Found notifications and next list position iter will be returned.
//Return largest id notification if asc is false.
func (n NopStore) List(condition []*Condition, start string, asc bool, count int) (result []*Notification, iter string, err error) {
	return nil, "", ErrStoreFeatureNotSupported
}

//Count count store with given search conditions
func (n NopStore) Count(condition []*Condition) (int, error) {
	return 0, ErrStoreFeatureNotSupported
}

//SupportedConditions return supported condition keyword list
func (n NopStore) SupportedConditions() ([]string, error) {
	return nil, ErrStoreFeatureNotSupported
}

//Remove remove notification by given id and return removed notification.
func (n NopStore) Remove(id string) (*Notification, error) {
	return nil, ErrStoreFeatureNotSupported
}
