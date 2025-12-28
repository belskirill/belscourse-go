package user

import (
	"belscourrsego/internal/domain/user"
	"belscourrsego/internal/interface/http/request"
	"belscourrsego/internal/interface/http/response"
	"net/http"
	"time"
)

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) error {
	var req CreateUserRequest

	if err := request.DecodeJSON(r, &req); err != nil {
		return err
	}

	if err := request.ValidateStruct(h.validate, req); err != nil {
		return err
	}

	res, err := h.create.CreateUser(r.Context(), user.UserCreate{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	})

	if err != nil {
		return err
	}

	return response.ResponseJSON(w, http.StatusCreated, res)
}

func (h *UserHandler) CreateSession(w http.ResponseWriter, r *http.Request) error {
	var req request.LoginRequest

	if err := request.DecodeJSON(r, &req); err != nil {
		return err
	}

	if err := request.ValidateStruct(h.validate, req); err != nil {
		return err
	}

	token, err := h.login.CreateSession(r.Context(), user.UserWithPassword{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	})

	if err != nil {
		return err
	}

	response.SetTokenCookie(w, token, response.CookieConfig{
		Name:     "access_token",
		Path:     "/",
		HTTPOnly: true,
		SameSite: http.SameSiteStrictMode,
		MaxAge:   15 * time.Minute,
	})

	return response.ResponseJSON(w, http.StatusAccepted, nil)
}
