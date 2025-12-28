package middleware

import (
	"belscourrsego/internal/domain/auth"
	"belscourrsego/internal/interface/http/httperr"
	"context"
	"net/http"

	"go.uber.org/zap"
)

func GetUserID(tkn auth.TokenService, logger *zap.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			cookie, err := r.Cookie("access_token")
			if err != nil {
				httperr.Write(w, logger, httperr.NewCodeUnauthenticated(httperr.MessageUnauthenticated, nil))
				return
			}

			token := cookie.Value
			if token == "" {
				httperr.Write(w, logger, httperr.NewCodeUnauthenticated(httperr.MessageUnauthenticated, nil))
				return
			}

			userID, err := tkn.ParseGetById(token)
			if err != nil {
				httperr.Write(w, logger, httperr.NewCodeUnauthenticated(httperr.MessageUnauthenticated, nil))
			}

			newCtx := context.WithValue(r.Context(), UserIDKey{}, userID)

			next.ServeHTTP(w, r.WithContext(newCtx))
		})
	}
}
