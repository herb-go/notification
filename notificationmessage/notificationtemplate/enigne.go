package notificationtemplate

type Engine interface {
	Parse(Template) (View, error)
}
