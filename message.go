package notification

import (
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
