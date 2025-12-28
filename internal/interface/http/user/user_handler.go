package user

import (
	"belscourrsego/internal/interface/http/middleware"
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
