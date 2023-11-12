package entity

import "fmt"

type DomainError struct {
	Err     error
	Code    string
	Message string
	Details string
}

func (e *DomainError) GetCode() string {
	return e.Code
}

func (e *DomainError) GetDetails() string {
	return e.Details
}

func (e *DomainError) GetMessage() string {
	return e.Message
}

func (e *DomainError) Error() string {
	return e.Err.Error()
}

func (e *DomainError) FullMessage() string {
	return fmt.Sprintf("[%s] %s :: %s", e.GetCode(), e.GetMessage(), e.GetDetails())
}
