package notification

import (
	"encoding/json"
	"strings"
)

//Content notification content type
//Content stores all messages in notification.
type Content map[string]string

//Set set give value with given name to content
//Name will be converted to lower
func (c Content) Set(name string, value string) {
	c[strings.ToLower(name)] = value
}

//Get get give value with given name from content
//Name will be converted to lower
func (c Content) Get(name string) string {
	return c[strings.ToLower(name)]
}

//Length return ccntent length
func (c Content) Length() int {
	return len(c)
}

//Range range over content with given function.
//Stop range if function return false,
func (c Content) Range(f func(string, string) bool) {
	for k := range c {
		if !f(k, c[k]) {
			return
		}
	}
}

//MustJSON must convent content to json format
//Panic if any error raised
func (c Content) MustJSON() string {
	bs, err := json.Marshal(c)
	if err != nil {
		panic(err)
	}
	return string(bs)
}

//NewContent create new content.
func NewContent() Content {
	return Content{}
}
