package notificationdelivery

//Config delivery driver config
type Config struct {
	//DeliveryType delivery driver type
	DeliveryType string
	//DeliveryConfig delivery dirver config
	DeliveryConfig func(v interface{}) error `config:", lazyload"`
}

//CreateDriver create delivery driver.
func (c *Config) CreateDriver() (DeliveryDriver, error) {
	return NewDriver(c.DeliveryType, c.DeliveryConfig)
}

//DeliveryServerConfig delivery config struct
type DeliveryServerConfig struct {
	//Delivery delivery keyword
	Delivery string
	//Disabled if delivery disabled
	Disabled bool
	//Description delivery description
	Description string
	Config
}

//CreateDeliverServer create delivery server.
func (c *DeliveryServerConfig) CreateDeliverServer() (*DeliveryServer, error) {
	d, err := c.Config.CreateDriver()
	if err != nil {
		return nil, err
	}
	return &DeliveryServer{
		Delivery:       c.Delivery,
		Description:    c.Description,
		Disabled:       c.Disabled,
		DeliveryDriver: d,
	}, nil
}

//DeliveryCenterConfig plain delivery center config
type DeliveryCenterConfig struct {
	Server []*DeliveryServerConfig
}

//CreateDeliveryCenter create delivery center
func (c *DeliveryCenterConfig) CreateDeliveryCenter() (DeliveryCenter, error) {
	p := NewPlainDeliveryCenter()
	for _, v := range c.Server {
		server, err := v.CreateDeliverServer()
		if err != nil {
			return nil, err
		}
		p.Insert(server)
	}
	return p, nil
}
