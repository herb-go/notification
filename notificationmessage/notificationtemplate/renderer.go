package notificationtemplate

import (
	"github.com/herb-go/notification"
	"github.com/herb-go/notification/notificationmessage"
)

type Renderer struct {
	View
	Formatter
}

func (r *Renderer) Render(m notificationmessage.Model) (notification.Content, error) {
	collection, err := r.Formatter.FormatModel(m)
	if err != nil {
		return nil, err
	}
	return r.View.Render(collection)
}
