package notificationtemplate

import "strings"

type Collection map[string]interface{}

func (c Collection) Set(name string, value interface{}) {
	c[strings.ToLower(name)] = value
}
func (c Collection) Get(name string) interface{} {
	return c[strings.ToLower(name)]
}

func NewCollection() Collection {
	return Collection{}
}
