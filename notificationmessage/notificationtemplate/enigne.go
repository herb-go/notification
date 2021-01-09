package notificationtemplate

type Engine interface {
	Parse(map[string]string, *Options) (View, error)
}

type Options struct {
}

func NewOptions() *Options {
	return &Options{}
}
