package user

import (
	"belscourrsego/internal/usecase/user"

	"github.com/go-playground/validator/v10"
)

type UserHandler struct {
	create    user.UseCreator
	login     user.CreateUserSession
	validate  *validator.Validate
	sendEmail user.SendEmailerCode
	webhook   user.WebHook
}

func NewHandler(usecase user.UseCreator, login user.CreateUserSession, validate *validator.Validate, sendEmail user.SendEmailerCode, wb user.WebHook) *UserHandler {
	return &UserHandler{
		create:    usecase,
		login:     login,
		validate:  validate,
		sendEmail: sendEmail,
		webhook:   wb,
	}
}
