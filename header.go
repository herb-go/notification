package notification

import (
	"net/url"
	"strings"
)

const (
	//HeaderNameTarget header name for message targe
	HeaderNameTarget = "target"
	//HeaderNameBatch header name for notification batch id
	HeaderNameBatch = "batch"
	//HeaderNameMessage header name for mesage id
	HeaderNameMessage = "mesasge"
	//HeaderNameTopic header name for message topic
	HeaderNameTopic = "topic"
	//HeaderNameDraftMode header name for draft-mode
	HeaderNameDraftMode = "draftmode"
)

//Header notification header struct
type Header map[string]string

//Set set give value with given name to header
//Name will be converted to lower
func (h Header) Set(name string, value string) {
	h[strings.ToLower(name)] = value
}

//Get get give value with given name from header
//Name will be converted to lower
func (h Header) Get(name string) string {
	return h[strings.ToLower(name)]
}

//String convert header to urlencoded format string
func (h Header) String() string {
	v := url.Values{}
	for k := range h {
		v.Set(k, h[k])
	}
	return v.Encode()
}

//NewHeader create new header
func NewHeader() Header {
	return Header{}
}
