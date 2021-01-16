package notification

import (
	"errors"
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

//CheckRequiredContent check if fields in content.
//If give fields is not in missed fields will be returned.
func CheckRequiredContent(c Content, fields []string) []string {
	var missed = []string{}
	for k := range fields {
		if c.Get(fields[k]) == "" {
			missed = append(missed, fields[k])
		}
	}
	return missed
}

//CheckRequiredContentError check if fields in content.
//If give fields is not in content,InvalidContentError will be returned.
//Otherwise nil will be returned.
func CheckRequiredContentError(c Content, fields []string) error {
	missed := CheckRequiredContent(c, fields)
	if len(missed) != 0 {
		return NewRequiredContentError(missed)
	}
	return nil
}

//ErrNotificationIDNotFound error raised if given notification not found
type ErrNotificationIDNotFound struct {
	NID string
}

//Error return error message
func (e *ErrNotificationIDNotFound) Error() string {
	return fmt.Sprintf("notification: id not found [%s]", e.NID)
}

//NewErrNotificationIDNotFound create new ErrNotificationIDNotFound
func NewErrNotificationIDNotFound(nid string) *ErrNotificationIDNotFound {
	return &ErrNotificationIDNotFound{
		NID: nid,
	}
}

//IsErrNotificationIDNotFound check if given error is ErrNotificationIDNotFound.
func IsErrNotificationIDNotFound(err error) bool {
	if err == nil {
		return false
	}
	_, ok := err.(*ErrNotificationIDNotFound)
	return ok
}

//ErrConditionNotSupported error rasied when condition not supproted
type ErrConditionNotSupported struct {
	Condition string
}

//Error return error message
func (e *ErrConditionNotSupported) Error() string {
	return fmt.Sprintf("notificationqueue: condition [%s] not supported", e.Condition)
}

//NewErrConditionNotSupported create new ErrConditionNotSupported
func NewErrConditionNotSupported(condition string) *ErrConditionNotSupported {
	return &ErrConditionNotSupported{
		Condition: condition,
	}
}

//IsErrConditionNotSupported check if given error is ErrConditionNotSupported
func IsErrConditionNotSupported(err error) bool {
	if err == nil {
		return false
	}
	_, ok := err.(*ErrConditionNotSupported)
	return ok
}

//ErrInvalidConditionValue error raised when condition value invalid
type ErrInvalidConditionValue struct {
	Condition string
}

//Error return error message
func (e *ErrInvalidConditionValue) Error() string {
	return fmt.Sprintf("notificationqueue: condition [%s] value invalid", e.Condition)
}

//NewErrInvalidConditionValue create new ErrInvalidConditionValue
func NewErrInvalidConditionValue(condition string) *ErrInvalidConditionValue {
	return &ErrInvalidConditionValue{
		Condition: condition,
	}
}

//IsErrInvalidConditionValue check if given error is ErrInvalidConditionValue
func IsErrInvalidConditionValue(err error) bool {
	if err == nil {
		return false
	}
	_, ok := err.(*ErrInvalidConditionValue)
	return ok
}

//ErrStoreFeatureNotSupported error raised when store feature not supported
var ErrStoreFeatureNotSupported = errors.New("store feature not supported")
