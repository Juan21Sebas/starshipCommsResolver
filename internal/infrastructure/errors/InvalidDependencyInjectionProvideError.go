package errors

import (
	gserrmodel "starshipCommsResolver/internal/pkg/entity"
)

type InvalidDependencyInjectionProvideError struct {
	*gserrmodel.DomainError
}

func NewInvalidDependencyInjectionProvideError(message string, original error) *InvalidDependencyInjectionProvideError {
	return &InvalidDependencyInjectionProvideError{&gserrmodel.DomainError{
		Err:     original,
		Code:    InvalidDiProvideError,
		Message: message,
		Details: original.Error(),
	}}
}
