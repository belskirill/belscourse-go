package httperr

type Code string

const (
	CodeInvalidInput       Code = "INVALID_INPUT"       // 400 — неверные данные
	CodeUnauthenticated    Code = "UNAUTHENTICATED"     // 401 — неавторизован
	CodeForbidden          Code = "FORBIDDEN"           // 403 — нет прав
	CodeNotFound           Code = "NOT_FOUND"           // 404 — ресурс не найден
	CodeConflict           Code = "CONFLICT"            // 409 — конфликт данных
	CodeInternal           Code = "INTERNAL"            // 500 — внутренняя ошибка
	CodeServiceUnavailable Code = "SERVICE_UNAVAILABLE" // 503 — временная недоступность
	CodeTimeout            Code = "TIMEOUT"             // 504 — таймаут
)

type HTTPError struct {
	Code    Code
	Message string
	Fields  map[string]string `json:"fields,omitempty"`
	Err     error             `json:"-"`
}

func NewCodeConflict(Message string, Fields map[string]string, err error) *HTTPError {
	return &HTTPError{
		Code:    CodeConflict,
		Message: Message,
		Fields:  Fields,
		Err:     err,
	}
}

func NewCodeInternal(err error) *HTTPError {
	return &HTTPError{
		Code:    CodeInternal,
		Message: "Internal server error",
		Err:     err,
	}
}

func (e *HTTPError) Error() string {
	return e.Message
}

func (e *HTTPError) Unwrap() error {
	return e.Err
}
