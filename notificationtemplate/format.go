package notificationtemplate

type Converter interface {
	Convert(string) (interface{}, error)
}

type Converters map[string]Converter

func (f *Converters) FormatModel(m Model) (Collection, error) {
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
