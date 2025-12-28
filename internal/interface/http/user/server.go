package user

import (
	"belscourrsego/internal/usecase/user"

	"github.com/go-playground/validator/v10"
)

type UserHandler struct {
	create   user.UseCreator
	login    user.CreateUserSession
	validate *validator.Validate
}

func NewHandler(usecase user.UseCreator, login user.CreateUserSession, validate *validator.Validate) *UserHandler {
	return &UserHandler{
		create:   usecase,
		login:    login,
		validate: validate,
	}
}
