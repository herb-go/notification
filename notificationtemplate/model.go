package notificationtemplate

import "strings"

type Model map[string]string

func (m Model) Set(name string, value string) {
	m[strings.ToLower(name)] = value
}
func (m Model) Get(name string) string {
	return m[strings.ToLower(name)]
}
