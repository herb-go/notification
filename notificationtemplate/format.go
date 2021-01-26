package notificationtemplate

type Field interface {
	Convert(string) (interface{}, error)
}

type Topic struct {
	Fields map[string]Field
}

func (t *Topic) ConvertModel(m Model) (Collection, error) {
	c := NewCollection()
	for k, v := range t.Fields {
		data, err := v.Convert(m.Get(k))
		if err != nil {
			return nil, err
		}
		c.Set(k, data)
	}
	return c, nil
}
