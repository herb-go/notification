package notificationqueue

import "fmt"

type ErrConditionNotSupported string

func (e ErrConditionNotSupported) Error() string {
	return fmt.Sprintf("notificationqueue: condition [%s] not supported", string(e))
}

func IsErrConditionNotSupported(err error) bool {
	if err == nil {
		return false
	}
	_, ok := err.(ErrConditionNotSupported)
	return ok
}
