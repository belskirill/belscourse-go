package app

import (
	"belscourrsego/internal/config"
	"belscourrsego/internal/infra/database/postgres"
	"context"
	"database/sql"
	"net/http"
)

type App struct {
	db  *sql.DB
	cfg config.Config
}

func NewApp(ctx context.Context) (*App, error) {

	cfg, err := config.Load()
	if err != nil {
		//
	}

	db, err := postgres.Connect(ctx, cfg.DB.DSN())

	return &App{
		db:  db,
		cfg: cfg,
	}, nil

}

func (a *App) BuilderHTTPServer() *http.Server {
	return NewHTTPServer(":8080", a.db)
}

func (a *App) Close() {
	a.db.Close()
}
