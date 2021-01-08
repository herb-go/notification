package notification

import (
	"fmt"
	"strings"
)

type Error struct {
	Errno int32
	Err   error
}

func (e *Error) Error() string {
	return e.Err.Error()
}

func NewError() *Error {
	return &Error{}
}

func NewRequiredContentError(fields []string) *Error {
	e := NewError()
	e.Errno = ErrnoInvalidContent
	e.Err = fmt.Errorf("notification: content [%s] required", strings.Join(fields, " , "))
	return e
}

func CheckRequiredContentError(c Content, fields []string) error {
	var missed = []string{}
	for k := range c {
		if c.Get(k) == "" {
			missed = append(missed, k)
		}
	}
	if len(missed) != 0 {
		return NewRequiredContentError(missed)
	}
	return nil
}

type ErrNofitactionIDNotFound string

func (e ErrNofitactionIDNotFound) Error() string {
	return fmt.Sprintf("notification: id not found [%s]", e)
}

type ErrorDeliveryNotFound string

func (e ErrorDeliveryNotFound) Error() string {
	return fmt.Sprintf("notification: delivery not found [%s]", e)

}
