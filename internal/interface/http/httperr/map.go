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
			return NewCodeConflict(nil, de.Message, de.Err)
		case user.ErrInvalidPassword:
			return NewCodeUnauthenticated(de.Message, de.Err)
		}
	}

	return NewCodeInternal(err)
}
