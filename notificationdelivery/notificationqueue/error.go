package notificationqueue

import (
	"errors"
	"fmt"
)

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

//ErrQueueDriverRequired queue driver required error.
var ErrQueueDriverRequired = errors.New("queue driver required")

//ErrDraftBoxRequired draft box required
var ErrDraftBoxRequired = errors.New("draft box required")

//ErrIDGeneratorRequired id generator required
var ErrIDGeneratorRequired = errors.New("id generator required")
