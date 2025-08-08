package domain

type DomainError struct {
	Message string
}

func (e *DomainError) Error() string {
	return e.Message
}

func NewDomainError(message string) *DomainError {
	return &DomainError{
		Message: message,
	}
}

func IsDomainError(err error) bool {
	_, ok := err.(*DomainError)
	return ok
}
