package user

import (
	"belscourrsego/internal/domain/user"
	"belscourrsego/internal/interface/http/response"
	"encoding/json"
	"net/http"
)

func (h *UserHandler) CreateSession(w http.ResponseWriter, r *http.Request) error {
	var req user.CreateUserRequest

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&req); err != nil {
		return err
	}
	err := h.create.CreateUser(r.Context(), req)
	if err != nil {
		return err
	}

	return response.ResponseJSON(w, http.StatusCreated, nil)
}
