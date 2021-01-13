package notificationdelivery

import "github.com/herb-go/notification"

type Config struct {
	DeliveryType   string
	DeliveryConfig func(v interface{}) error `config:", lazyload"`
}

func (c *Config) CreateDriver() (notification.DeliveryDriver, error) {
	return NewDriver(c.DeliveryType, c.DeliveryConfig)
}

type DeliverServerConfig struct {
	Delivery    string
	Description string
	Config
}

func (c *DeliverServerConfig) CreateDeliverServer() (*notification.DeliveryServer, error) {
	d, err := c.Config.CreateDriver()
	if err != nil {
		return nil, err
	}
	return &notification.DeliveryServer{
		Delivery:       c.Delivery,
		Description:    c.Description,
		DeliveryDriver: d,
	}, nil
}

type DeliveryCenterConfig []*DeliverServerConfig

func (c *DeliveryCenterConfig) CreateDeliveryCenter() (DeliveryCenter, error) {
	p := NewPlainDeliveryCenter()
	for _, v := range *c {
		server, err := v.CreateDeliverServer()
		if err != nil {
			return nil, err
		}
		p.Insert(server)
	}
	return p, nil
}
