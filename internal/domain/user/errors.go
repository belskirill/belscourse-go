package user

type DomainErrorCode string

const (
	ErrUserAlreadyExists DomainErrorCode = "USER_ALREADY_EXISTS"
	ErrUserNotFound      DomainErrorCode = "USER_NOT_FOUND"
	ErrInvalidPassword   DomainErrorCode = "INVALID_PASSWORD"
)

type DomainError struct {
	Code    DomainErrorCode
	Fields  map[string]string
	Message string
	Err     error
}

func New(code DomainErrorCode, err error, fields map[string]string, message string) *DomainError {
	return &DomainError{
		Code:    code,
		Err:     err,
		Fields:  fields,
		Message: message,
	}
}

func (e *DomainError) Error() string {
	return e.Message
}

func (e *DomainError) Unwrap() error {
	return e.Err
}
