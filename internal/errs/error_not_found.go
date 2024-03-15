package errs

import "fmt"

type ErrorNotFound struct {
	Msg string
}

func (e *ErrorNotFound) Error() string {
	return e.Msg
}

func NewErrorNotFound(username string) error {
	return &ErrorNotFound{fmt.Sprintf("user with username %s not found", username)}
}
