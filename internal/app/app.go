package app

import (
	"belscourrsego/internal/config"
	"belscourrsego/internal/infra/database/postgres"
	"context"
	"database/sql"
	"net/http"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type App struct {
	db     *sql.DB
	cfg    config.Config
	Logger *zap.Logger
}

func NewApp(ctx context.Context) (*App, error) {

	zapCfg := zap.NewProductionConfig()
	zapCfg.EncoderConfig.EncodeDuration = zapcore.StringDurationEncoder

	logger, err := zapCfg.Build()
	if err != nil {
		return nil, err
	}

	cfg, err := config.Load()
	if err != nil {
		return nil, err
	}

	db, err := postgres.Connect(ctx, cfg.DB.DSN())
	if err != nil {
		return nil, err
	}

	return &App{
		db:     db,
		cfg:    cfg,
		Logger: logger,
	}, nil

}

func (a *App) BuilderHTTPServer() *http.Server {
	return NewHTTPServer(":8080", a.db, a.Logger)
}

func (a *App) Close() {
	if err := a.db.Close(); err != nil {
		a.Logger.Error("db close error", zap.Error(err))
	}

	if err := a.Logger.Sync(); err != nil {
		a.Logger.Error("logger sync error", zap.Error(err))
	}
}
