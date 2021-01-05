package notification

import (
	"net/url"
)

var HeaderNameTarget = "target"
var HeaderNameBatch = "batch"
var HeaderNameMessage = "mesasge"
var HeaderNameTopic = "topic"

type Header map[string]string

func (h Header) Set(name string, value string) {
	h[name] = value
}
func (h Header) Get(name string) string {
	return h[name]
}
func (h Header) String() string {
	v := url.Values{}
	for k := range h {
		v.Set(k, h[k])
	}
	return v.Encode()
}

func NewHeader() Header {
	return Header{}
}
