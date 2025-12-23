package middleware

import (
	"net/http"
)

type AppHandler func(http.ResponseWriter, *http.Request) error

func Wrap(app AppHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := app(w, r); err != nil {
			//
		}

	}
}
