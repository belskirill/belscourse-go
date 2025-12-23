package user

import "belscourrsego/internal/usecase/user"

type UserHandler struct {
	create user.UseCreatorSession
}

func NewHandler(usecase user.UseCreatorSession) *UserHandler {
	return &UserHandler{
		create: usecase,
	}
}
