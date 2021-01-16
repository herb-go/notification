package notificationqueue

import (
	"errors"
)

//ErrQueueDriverRequired queue driver required error.
var ErrQueueDriverRequired = errors.New("queue driver required")

//ErrIDGeneratorRequired id generator required
var ErrIDGeneratorRequired = errors.New("id generator required")
