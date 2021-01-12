package notificationqueue

import "fmt"

//ErrConditionNotSupported error rasied when condition not supproted
type ErrConditionNotSupported string

//Error return error message
func (e ErrConditionNotSupported) Error() string {
	return fmt.Sprintf("notificationqueue: condition [%s] not supported", string(e))
}

//IsErrConditionNotSupported check if given error is ErrConditionNotSupported
func IsErrConditionNotSupported(err error) bool {
	if err == nil {
		return false
	}
	_, ok := err.(ErrConditionNotSupported)
	return ok
}
