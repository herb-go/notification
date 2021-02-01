package notificationview

type ViewConfig struct {
	//RenderType
	Type string
	//RenderConfig render config
	Config func(v interface{}) error `config:", lazyload"`
}

func (c *ViewConfig) CreateView() (View, error) {
	return NewView(c.Type, c.Config)
}

type NamedViewConfig struct {
	Name        string
	Description string
	ViewConfig
}

type Config struct {
	Views []*NamedViewConfig
}

func (c *Config) CreateViewCenter() (ViewCenter, error) {
	vc := NewViewCenter()
	for _, v := range c.Views {
		if v.Name == "" {
			return nil, ErrEmptyViewName
		}
		view, err := v.CreateView()
		if err != nil {
			return nil, err
		}
		vc.Set(v.Name, view)
	}
	return vc, nil
}
