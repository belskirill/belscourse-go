package app

import (
	"belscourrsego/internal/config"
	"belscourrsego/internal/interface/http/middleware"
	"database/sql"
	"net/http"

	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

func NewHTTPServer(addr string, db *sql.DB, logger *zap.Logger, validate *validator.Validate, cfg config.Config) *http.Server {
	root := http.NewServeMux()

	router := buildUserHandlers(db, validate, cfg)
	userV1 := http.NewServeMux()
	authV1 := http.NewServeMux()

	authV1.Handle("/create", middleware.CheckMethod(http.MethodPost, logger)(middleware.Wrap(router.user.CreateUser, logger)))
	authV1.Handle("/create_session", middleware.CheckMethod(http.MethodPost, logger)(middleware.Wrap(router.user.CreateSession, logger)))

	userV1.Handle("/edit", middleware.CheckMethod(http.MethodPost, logger)(middleware.Wrap(router.user.EditProfile, logger)))

	root.Handle("/user/v1/auth/", http.StripPrefix("/user/v1/auth", authV1))
	root.Handle("/user/v1/", middleware.GetUserID(router.jwtService, logger)(http.StripPrefix("/user/v1", userV1)))

	handler := middleware.AccessLogging(logger)(root)
	handler = middleware.Recovery(logger)(handler)

	return &http.Server{
		Addr:    addr,
		Handler: handler,
	}
}
