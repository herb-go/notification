package notification

type Sender interface {
	Start() error
	Stop() error
	SendNotification(*NotificationInstance) error
	Name() string
}

type ChanSender struct {
	ID   string
	Size int
	C    chan *NotificationInstance
}

func (c *ChanSender) Start() error {
	c.C = make(chan *NotificationInstance, c.Size)
	return nil
}
func (c *ChanSender) Stop() error {
	close(c.C)
	return nil
}
func (c *ChanSender) SendNotification(n *NotificationInstance) error {
	c.C <- n
	return nil
}
func (c *ChanSender) Name() string {
	return c.ID
}

func NewChanSender() *ChanSender {
	return &ChanSender{}
}
