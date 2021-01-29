package notificationrender

import "fmt"

//ErrRendererNotFound error raised if renderer not found
type ErrRendererNotFound struct {
	Renderer string
}

//Error return error message
func (e *ErrRendererNotFound) Error() string {
	return fmt.Sprintf("notification render: renderer not found [%s]", e.Renderer)
}

//NewErrRendererNotFound create new ErrDeliveryNotFound
func NewErrRendererNotFound(renderer string) *ErrRendererNotFound {
	return &ErrRendererNotFound{
		Renderer: renderer,
	}
}

//IsErrRendererNotFound check if given error is ErrDeliveryNotFound.
func IsErrRendererNotFound(err error) bool {
	if err == nil {
		return false
	}
	_, ok := err.(*ErrRendererNotFound)
	return ok
}
