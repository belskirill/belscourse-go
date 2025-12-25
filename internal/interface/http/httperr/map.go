package httperr

import (
	"belscourrsego/internal/domain/user"
	"errors"
)

func MapError(err error) *HTTPError {
	var de *user.DomainError
	if errors.As(err, &de) {
		switch de.Code {
		case user.ErrUserAlreadyExists:
			return NewCodeConflict(de.Message, de.Fields, de.Err)
		}
	}

	return NewCodeInternal(err)
}
