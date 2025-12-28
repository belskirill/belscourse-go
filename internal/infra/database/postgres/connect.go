package postgres

import (
	"context"
	"database/sql"
	"time"

	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

func Connect(ctx context.Context, dsn string, logger *zap.Logger) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		if err := db.Close(); err != nil {
			logger.Warn("Failed to close database", zap.Error(err))
			return nil, err
		}

		logger.Warn("Failed to ping database", zap.Error(err))
		return nil, err
	}

	logger.Info("Successfully connected to database")
	return db, nil
}
