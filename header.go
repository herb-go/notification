package notification

import (
	"net/url"
	"strings"
)

var HeaderNameNotifactionID = "nid"
var HeaderNameTarget = "target"
var HeaderNameBatch = "batch"
var HeaderNameMessage = "mesasge"
var HeaderNameTopic = "topic"
var HeaderNameDraftMode = "draftmode"

type Header map[string]string

func (h Header) Set(name string, value string) {
	h[strings.ToLower(name)] = value
}
func (h Header) Get(name string) string {
	return h[strings.ToLower(name)]
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
