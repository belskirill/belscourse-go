package user

type DomainErrorCode string

const (
	ErrUserAlreadyExists DomainErrorCode = "USER_ALREADY_EXISTS"
	ErrUserNotFound      DomainErrorCode = "USER_NOT_FOUND"
	ErrInvalidPassword   DomainErrorCode = "INVALID_PASSWORD"
)

var defaultMessages = map[DomainErrorCode]string{
	ErrUserAlreadyExists: "user already exists",
	ErrUserNotFound:      "user not found",
	ErrInvalidPassword:   "invalid credentials",
}

type DomainError struct {
	Code    DomainErrorCode
	Message string
	Err     error
}

func New(code DomainErrorCode, err error) *DomainError {
	return &DomainError{
		Code:    code,
		Err:     err,
		Message: defaultMessages[code],
	}
}

func (e *DomainError) Error() string {
	return e.Message
}

func (e *DomainError) Unwrap() error {
	return e.Err
}
