package notification

import (
	"encoding/json"
	"strings"
)

type Content map[string]string

func (c Content) Set(name string, value string) {
	c[strings.ToLower(name)] = value
}
func (c Content) Get(name string) string {
	return c[strings.ToLower(name)]
}

func (c Content) MustJSON() string {
	bs, err := json.Marshal(c)
	if err != nil {
		panic(err)
	}
	return string(bs)
}

func NewContent() Content {
	return Content{}
}
