package notification

import (
	"fmt"
	"strings"
)

//Error notofication error struct
type Error struct {
	//Notofication error no
	Errno int32
	//Err real error
	Err error
}

//Error return error string
func (e *Error) Error() string {
	return e.Err.Error()
}

//NewError create error
func NewError() *Error {
	return &Error{}
}

//NewRequiredContentError create required content error with given fields.
func NewRequiredContentError(fields []string) *Error {
	e := NewError()
	e.Errno = ErrnoInvalidContent
	e.Err = fmt.Errorf("notification: content [%s] required", strings.Join(fields, " , "))
	return e
}

//IsInvalidContentError check if given error is invalid content error.
func IsInvalidContentError(err error) bool {
	if err == nil {
		return false
	}
	e, ok := err.(*Error)
	if !ok {
		return false
	}
	return e.Errno == ErrnoInvalidContent
}

//CheckRequiredContentError check if fields in content.
//If give fields is not in content,InvalidContentError will be returned.
//Otherwise nil will be returned.
func CheckRequiredContentError(c Content, fields []string) error {
	var missed = []string{}
	for k := range fields {
		if c.Get(fields[k]) == "" {
			missed = append(missed, fields[k])
		}
	}
	if len(missed) != 0 {
		return NewRequiredContentError(missed)
	}
	return nil
}

//ErrNotificationIDNotFound error raised if given notification not found
type ErrNotificationIDNotFound string

//Error return error message
func (e ErrNotificationIDNotFound) Error() string {
	return fmt.Sprintf("notification: id not found [%s]", string(e))
}

//IsErrNotificationIDNotFound check if given error is ErrNotificationIDNotFound.
func IsErrNotificationIDNotFound(err error) bool {
	if err == nil {
		return false
	}
	_, ok := err.(ErrNotificationIDNotFound)
	return ok
}

//ErrDeliveryNotFound error raised if given delivery not found
type ErrDeliveryNotFound string

//Error return error message
func (e ErrDeliveryNotFound) Error() string {
	return fmt.Sprintf("notification: delivery not found [%s]", string(e))
}

//IsErrDeliveryNotFound check if given error is ErrDeliveryNotFound.
func IsErrDeliveryNotFound(err error) bool {
	if err == nil {
		return false
	}
	_, ok := err.(ErrDeliveryNotFound)
	return ok
}
