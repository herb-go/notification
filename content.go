package notification

type Content map[string]string

func (c Content) Set(name string, value string) {
	c[name] = value
}
func (c Content) Get(name string) string {
	return c[name]
}

func NewContent() Content {
	return Content{}
}
