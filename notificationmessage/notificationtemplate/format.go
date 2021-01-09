package notificationtemplate

import (
	"github.com/herb-go/notification/notificationmessage"
)

type Format interface {
	Convert(string) (interface{}, error)
}

type Formatter map[string]Format

func (f *Formatter) FormatModel(m notificationmessage.Model) (Collection, error) {
	c := NewCollection()
	for k, v := range *f {
		data, err := v.Convert(m.Get(k))
		if err != nil {
			return nil, err
		}
		c.Set(k, data)
	}
	return c, nil
}
