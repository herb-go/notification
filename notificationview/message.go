package notificationview

import (
	"encoding/json"
	"strings"
)

//Message Message type
//Message stores raw information.
type Message map[string]string

//Set set give value with given name to message
//Name will be converted to lower
func (m Message) Set(name string, value string) {
	m[strings.ToLower(name)] = value
}

//MustSetAsJSON set give value as json with given name to message
//Name will be converted to lower
//Panic if marshal fail
func (m Message) MustSetAsJSON(name string, v interface{}) {
	bs, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	m.Set(name, string(bs))

}

//Get get give value with given name from message
//Name will be converted to lower
func (m Message) Get(name string) string {
	return m[strings.ToLower(name)]
}

//Length return message length
func (m Message) Length() int {
	return len(m)
}

//Range range over message with given function.
//Stop range if function return false,
func (m Message) Range(f func(string, string) bool) {
	for k := range m {
		if !f(k, m[k]) {
			return
		}
	}
}

//NewMessage create new message.
func NewMessage() Message {
	return Message{}
}
