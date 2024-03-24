package errs

import "fmt"

type ErrorUsernameAlreadyExists struct {
	Msg string
}

func (e *ErrorUsernameAlreadyExists) Error() string {
	return e.Msg
}

func NewErrorUsernameAlreadyExists(username string) error {
	return &ErrorUsernameAlreadyExists{fmt.Sprintf("user with username %s already exists", username)}
}
