package request

type LoginRequest struct {
	Username string `json:"username" validate:"omitempty,min=3"`
	Email    string `json:"email" validate:"omitempty,email"`
	Password string `json:"password" validate:"required,min=8"`
}
