package app

import (
	"belscourrsego/internal/interface/http/middleware"
	"database/sql"
	"net/http"

	"go.uber.org/zap"
)

func NewHTTPServer(addr string, db *sql.DB, logger *zap.Logger) *http.Server {
	mux := http.NewServeMux()

	user := buildUserHandlers(db)

	mux.HandleFunc("/create_session", middleware.Wrap(user.createSession.CreateSession, logger))

	handler := middleware.AccessLogging(logger)(mux)

	return &http.Server{
		Addr:    addr,
		Handler: handler,
	}
}
