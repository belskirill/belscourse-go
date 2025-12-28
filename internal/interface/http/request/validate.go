package request

import (
	"belscourrsego/internal/interface/http/httperr"

	"github.com/go-playground/validator/v10"
)

func validationMessage(tag string) string {
	switch tag {
	case "required":
		return "field is required"
	case "email":
		return "invalid email format"
	case "xor":
		return "only one of username or email must be provided"
	default:
		return "invalid value"
	}
}

func RegisterValidations(v *validator.Validate) {
	v.RegisterStructValidation(
		LoginStructValidation,
		LoginRequest{},
	)
}

func ValidateStruct[T any](validate *validator.Validate, dst T) error {
	if err := validate.Struct(dst); err != nil {
		fields := map[string]string{}

		for _, fe := range err.(validator.ValidationErrors) {
			fields[fe.Field()] = validationMessage(fe.Tag())
		}

		return httperr.NewCodeInvalidInput(fields, err)
	}

	return nil
}

func LoginStructValidation(sl validator.StructLevel) {
	req := sl.Current().Interface().(LoginRequest)

	hasUsername := req.Username != ""
	hasEmail := req.Email != ""

	switch {
	case hasUsername && hasEmail:
		sl.ReportError(
			req.Username,
			"username",
			"Username",
			"xor",
			"",
		)
		sl.ReportError(
			req.Email,
			"email",
			"Email",
			"xor",
			"",
		)
	case !hasUsername && !hasEmail:
		sl.ReportError(
			req.Username,
			"username",
			"Username",
			"required",
			"",
		)
		sl.ReportError(
			req.Email,
			"email",
			"Email",
			"required",
			"",
		)
	}
}
