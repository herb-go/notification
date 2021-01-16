package notification

//Checker notification checker interface
//Return if notification passed check
type Checker interface {
	Check(n *Notification) (bool, error)
}

type CheckerFunc func(n *Notification) (bool, error)

func (f CheckerFunc) Check(n *Notification) (bool, error) {
	return f(n)
}

type HasHeaderChecker string

func (c HasHeaderChecker) Check(n *Notification) (bool, error) {
	return n.Header.Get(string(c)) != "", nil
}
