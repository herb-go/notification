package notificationtemplate

type Template struct {
	Delivery              string
	TTLInSeconds          int
	Topic                 string
	HeaderTemplate        map[string]string
	ContentHeaderTemplate map[string]string
}

func NewTemplate() *Template {
	return &Template{
		HeaderTemplate:        map[string]string{},
		ContentHeaderTemplate: map[string]string{},
	}
}
