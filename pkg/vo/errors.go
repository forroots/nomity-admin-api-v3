package vo

import "fmt"

type VOError struct {
	Message string `json:"message"`
}

func (e *VOError) Error() string {
	return fmt.Sprintf("%s", e.Message)
}
func NewVOErrorf(format string, args ...any) *VOError {
	return &VOError{
		Message: fmt.Sprintf(format, args...),
	}
}
