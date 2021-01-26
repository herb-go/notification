package notificationtemplate

type Options struct {
	Helpers    []*NamedHelper
	Formatters []*NamedFormatter
}

func (o *Options) Reset() {
	o.Helpers = []*NamedHelper{}
	o.Formatters = []*NamedFormatter{}
}
func (o *Options) AddHelper(name string, helper Helper) {
	o.Helpers = append(o.Helpers, &NamedHelper{
		Name:   name,
		Helper: helper,
	})
}
func (o *Options) AddFormatter(name string, formatter Formatter) {
	o.Formatters = append(o.Formatters, &NamedFormatter{
		Name:      name,
		Formatter: formatter,
	})
}
func NewOptions() *Options {
	return &Options{
		Helpers:    []*NamedHelper{},
		Formatters: []*NamedFormatter{},
	}
}

var DefaultOptions = NewOptions()
