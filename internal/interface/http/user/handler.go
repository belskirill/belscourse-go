package user

import (
	"belscourrsego/internal/interface/http/response"
	"net/http"
)

func (h *UserHandler) CreateSession(w http.ResponseWriter, r *http.Request) error {
	err := h.create.CreateUser(r.Context())
	if err != nil {
		return err
	}

	return response.ResponseJSON(w, http.StatusCreated, nil)
}
