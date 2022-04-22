package errs

import "fmt"

type Err struct {
	code    int
	message string
}

func New(code int, message string) error {
	err := &Err{
		code:    code,
		message: message,
	}
	return err
}

func (e Err) Error() string {
	return fmt.Sprintf("code(%v): %v", e.code, e.message)
}

func (e Err) Code() int {
	return e.code
}

func (e Err) Message() string {
	return e.message
}
