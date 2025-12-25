package middleware

import (
	"belscourrsego/internal/interface/http/httperr"
	"net/http"

	"go.uber.org/zap"
)

type AppHandler func(http.ResponseWriter, *http.Request) error

func Wrap(app AppHandler, logger *zap.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := app(w, r); err != nil {
			httpError := httperr.MapError(err)
			httperr.Write(w, logger, httpError)
		}

	}
}
