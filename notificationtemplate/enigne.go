package notificationtemplate

type Engine interface {
	Parse(map[string]string, *Options) (View, error)
	ApplyOptions(*Options) error
}
