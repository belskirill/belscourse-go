package middleware

import (
	"belscourrsego/internal/interface/http/httperr"
	"net/http"

	"go.uber.org/zap"
)

func Recovery(logger *zap.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if err := recover(); err != nil {
					logger.Error("panic recovered",
						zap.Any("panic", err),
						zap.Stack("stack"),
						zap.String("method", r.Method),
						zap.String("path", r.URL.Path),
					)

					httperr.Write(w, logger, httperr.NewCodeInternal(nil))
				}
			}()

			next.ServeHTTP(w, r)
		})
	}
}
