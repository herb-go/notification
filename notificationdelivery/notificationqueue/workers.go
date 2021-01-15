package notificationqueue

const DirectiveNameCommon = "common"

type CommonConfig struct {
	Workers int
}

func (c *CommonConfig) AppylToPublisher(p *Publisher) error {
	p.Notifier.Workers = c.Workers
	return nil
}

var FactoryCommon = func(loader func(v interface{}) error) (Directive, error) {
	c := &CommonConfig{}
	err := loader(c)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func init() {
	Register(DirectiveNameCommon, FactoryCommon)
}
