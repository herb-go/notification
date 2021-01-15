package notificationqueue

//Config publisher config
type Config struct {
	Directives []*DirectiveConfig
}

//ApplyTo apply config to publisher
func (c *Config) ApplyTo(p *Publisher) error {
	for _, v := range c.Directives {
		d, err := NewDirective(v.Directive, v.DirectiveConfig)
		if err != nil {
			return err
		}
		err = d.AppylToPublisher(p)
		if err != nil {
			return err
		}
	}
	return nil
}

//DirectiveConfig Directive config
type DirectiveConfig struct {
	//Directive directive keyword
	Directive string
	//DirectiveConfig directive config
	DirectiveConfig func(v interface{}) error `config:", lazyload"`
}
