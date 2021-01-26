package notificationtemplate

type Engine interface {
	Parse(*Template, *Options) (View, error)
	ApplyOptions(*Options) error
}
