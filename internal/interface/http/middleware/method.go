package middleware

import (
	"belscourrsego/internal/interface/http/httperr"
	"net/http"

	"go.uber.org/zap"
)

func CheckMethod(method string, logger *zap.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != method {
				w.Header().Set("Allow", method)
				httperr.Write(w, logger, httperr.NewMethodNotAllowed(r, method))
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
