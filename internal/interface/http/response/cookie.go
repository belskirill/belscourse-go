package response

import (
	"net/http"
	"time"
)

type CookieConfig struct {
	Name     string
	Path     string
	Domain   string
	Secure   bool
	HTTPOnly bool
	SameSite http.SameSite
	MaxAge   time.Duration
}

func SetTokenCookie(w http.ResponseWriter, token string, cfg CookieConfig) {
	http.SetCookie(w, &http.Cookie{
		Name:     cfg.Name,
		Value:    token,
		Path:     cfg.Path,
		Domain:   cfg.Domain,
		HttpOnly: cfg.HTTPOnly,
		Secure:   cfg.Secure,
		SameSite: cfg.SameSite,
		MaxAge:   int(cfg.MaxAge.Seconds()),
	})
}
