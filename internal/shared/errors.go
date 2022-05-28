package shared

import "fmt"

type ErrorAtPosition struct {
	Position int
	Err      error
}

func NewErrorAtPosition(err error, pos int) *ErrorAtPosition {
	return &ErrorAtPosition{
		Position: pos,
		Err:      err,
	}
}

func (eap *ErrorAtPosition) Error() string {
	return fmt.Sprintf("%s at %d", eap.Err, eap.Position)
}
