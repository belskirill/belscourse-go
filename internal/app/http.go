package app

import (
	"belscourrsego/internal/interface/http/middleware"
	"database/sql"
	"net/http"
)

func NewHTTPServer(addr string, db *sql.DB) *http.Server {
	mux := http.NewServeMux()

	user := buildUserHandlers(db)

	mux.HandleFunc("/create_session", middleware.Wrap(user.createSession.CreateSession))

	return &http.Server{
		Addr:    addr,
		Handler: mux,
	}
}
