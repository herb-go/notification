package notificationtemplate

type Template struct {
	Delivery string
	TTL      int
	Topic    string
	Header   map[string]string
	Content  map[string]string
}

func NewTemplate() *Template {
	return &Template{
		Header:  map[string]string{},
		Content: map[string]string{},
	}
}
