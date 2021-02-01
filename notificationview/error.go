package notificationview

import (
	"errors"
	"fmt"
)

//ErrViewNotFound error raised if view not found
type ErrViewNotFound struct {
	View string
}

//Error return error message
func (e *ErrViewNotFound) Error() string {
	return fmt.Sprintf("notification view: view not found [%s]", e.View)
}

//NewErrViewNotFound create new ErrDeliveryNotFound
func NewErrViewNotFound(view string) *ErrViewNotFound {
	return &ErrViewNotFound{
		View: view,
	}
}

//IsErrViewNotFound check if given error is ErrDeliveryNotFound.
func IsErrViewNotFound(err error) bool {
	if err == nil {
		return false
	}
	_, ok := err.(*ErrViewNotFound)
	return ok
}

var ErrEmptyViewName = errors.New("notification view:empty view name")
