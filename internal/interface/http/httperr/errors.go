package httperr

import "net/http"

type Code string

const (
	CodeInvalidInput       Code = "INVALID_INPUT"   // 400 — неверные данные
	CodeUnauthenticated    Code = "UNAUTHENTICATED" // 401 — неавторизован
	CodeForbidden          Code = "FORBIDDEN"       // 403 — нет прав
	CodeNotFound           Code = "NOT_FOUND"       // 404 — ресурс не найден
	MethodNotAllowed       Code = "METHOD_NOT_ALLOWED"
	CodeConflict           Code = "CONFLICT"            // 409 — конфликт данных
	CodeInternal           Code = "INTERNAL"            // 500 — внутренняя ошибка
	CodeServiceUnavailable Code = "SERVICE_UNAVAILABLE" // 503 — временная недоступность
	CodeTimeout            Code = "TIMEOUT"             // 504 — таймаут
)

const (
	MessageUnauthenticated string = "Unauthenticated"
)

var defaultMessages = map[Code]string{
	CodeInternal:        "Internal server error",
	CodeInvalidInput:    "Invalid input",
	MethodNotAllowed:    "Method not allowed",
	CodeUnauthenticated: "Unauthenticated",
}

type HTTPError struct {
	Code    Code
	Message string
	Fields  map[string]string `json:"fields,omitempty"`
	Err     error             `json:"-"`
}

func NewCodeConflict(Fields map[string]string, message string, err error) *HTTPError {
	return &HTTPError{
		Code:    CodeConflict,
		Message: message,
		Fields:  Fields,
		Err:     err,
	}
}

func NewCodeUnauthenticated(message string, err error) *HTTPError {
	return &HTTPError{
		Code:    CodeUnauthenticated,
		Message: message,
		Err:     err,
	}
}

func NewCodeInternal(err error) *HTTPError {
	return &HTTPError{
		Code:    CodeInternal,
		Message: defaultMessages[CodeInternal],
		Err:     err,
	}
}

func NewCodeInvalidInput(Fields map[string]string, err error) *HTTPError {
	return &HTTPError{
		Code:    CodeInvalidInput,
		Message: defaultMessages[CodeInvalidInput],
		Fields:  Fields,
		Err:     err,
	}
}

func NewMethodNotAllowed(r *http.Request, allowMethod string) *HTTPError {
	return &HTTPError{
		Code:    MethodNotAllowed,
		Message: defaultMessages[MethodNotAllowed],
		Fields: map[string]string{
			"method": r.Method,
			"allow":  allowMethod,
		},
	}
}

func (e *HTTPError) Error() string {
	return e.Message
}

func (e *HTTPError) Unwrap() error {
	return e.Err
}
