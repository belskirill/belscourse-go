package app

import (
	"belscourrsego/internal/config"
	"belscourrsego/internal/infra/database/postgres"
	"belscourrsego/internal/interface/http/request"
	"context"
	"database/sql"
	"net/http"

	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type App struct {
	db       *sql.DB
	cfg      config.Config
	Logger   *zap.Logger
	validate *validator.Validate
}

func NewApp(ctx context.Context) (*App, error) {

	zapCfg := zap.NewProductionConfig()
	zapCfg.EncoderConfig.EncodeDuration = zapcore.StringDurationEncoder

	validate := validator.New()
	request.RegisterValidations(validate)

	logger, err := zapCfg.Build()
	if err != nil {
		return nil, err
	}

	cfg, err := config.Load(logger)
	if err != nil {
		return nil, err
	}

	db, err := postgres.Connect(ctx, cfg.DB.DSN(), logger)
	if err != nil {
		return nil, err
	}

	return &App{
		db:       db,
		cfg:      cfg,
		Logger:   logger,
		validate: validate,
	}, nil

}

func (a *App) BuilderHTTPServer() *http.Server {
	return NewHTTPServer("127.0.0.1:8080", a.db, a.Logger, a.validate, a.cfg)
}

func (a *App) Close() {
	if err := a.db.Close(); err != nil {
		a.Logger.Error("db close error", zap.Error(err))
	}

	if err := a.Logger.Sync(); err != nil {
		a.Logger.Error("logger sync error", zap.Error(err))
	}
}
