package errors

import (
	gserrmodel "starshipCommsResolver/internal/pkg/entity"
)

type DependencyNotFoundErrors struct {
	*gserrmodel.DomainError
}

func NewDependencyNotFoundError(message string, original error) *DependencyNotFoundErrors {
	return &DependencyNotFoundErrors{&gserrmodel.DomainError{
		Err:     original,
		Code:    DependencyNotFoundError,
		Message: message,
		Details: original.Error(),
	}}
}
