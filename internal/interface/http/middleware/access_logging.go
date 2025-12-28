package middleware

import (
	"net/http"
	"time"

	"go.uber.org/zap"
)

type WrapStatusCode struct {
	http.ResponseWriter
	StatusCode int
}

func (w *WrapStatusCode) WriteHeader(statusCode int) {
	w.StatusCode = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}

func AccessLogging(logger *zap.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			rw := &WrapStatusCode{
				ResponseWriter: w,
				StatusCode:     http.StatusOK,
			}

			next.ServeHTTP(rw, r)

			duration := time.Since(start)

			logger.Info("user_service",
				zap.String("method", r.Method),
				zap.Int("status_code", rw.StatusCode),
				zap.String("url", r.URL.String()),
				zap.Duration("duration", duration),
			)
		})
	}
}
