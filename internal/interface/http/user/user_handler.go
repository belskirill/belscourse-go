package user

import (
	"belscourrsego/internal/interface/http/middleware"
	"belscourrsego/internal/interface/http/response"
	"fmt"
	"net/http"
)

func (h *UserHandler) EditProfile(w http.ResponseWriter, r *http.Request) error {
	userID, ok := middleware.FromUserIDContext(r.Context())
	if !ok {
		return nil
	}

	fmt.Println(userID)
	return nil
}

func (h *UserHandler) SendEmail(w http.ResponseWriter, r *http.Request) error {
	userID, ok := middleware.FromUserIDContext(r.Context())
	if !ok {
		return nil
	}

	if err := h.sendEmail.SendEmailCode(r.Context(), userID); err != nil {
		return err
	}

	return response.ResponseJSON(w, http.StatusOK, nil)
}
