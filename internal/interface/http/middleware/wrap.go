package middleware

import (
	"belscourrsego/internal/interface/http/httperr"
	"errors"
	"net/http"

	"go.uber.org/zap"
)

type AppHandler func(http.ResponseWriter, *http.Request) error

func Wrap(app AppHandler, logger *zap.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := app(w, r); err != nil {
			var exc *httperr.HTTPError
			if errors.As(err, &exc) {
				httperr.Write(w, logger, exc)
				return
			}

			httpError := httperr.MapError(err)
			httperr.Write(w, logger, httpError)
		}

	}
}
