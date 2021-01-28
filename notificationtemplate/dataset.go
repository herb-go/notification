package notificationtemplate

import "strings"

type Dataset map[string]interface{}

func (d Dataset) Set(name string, value interface{}) {
	d[strings.ToLower(name)] = value
}
func (d Dataset) Get(name string) interface{} {
	return d[strings.ToLower(name)]
}

func NewDataset() Dataset {
	return Dataset{}
}
