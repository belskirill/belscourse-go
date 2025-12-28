package config

import (
	"fmt"
	"time"

	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

type Config struct {
	DB  DBConfig
	JWT JwtConfig
	APP AppConfig
}

type AppConfig struct {
	MODE string `env:"MODE,required"`
}

type JwtConfig struct {
	JWTSecret string        `env:"JWT_SECRET,required"`
	JWTExpire time.Duration `env:"JWT_ACCESS_TTL,required"`
}

type DBConfig struct {
	Host     string `env:"DB_HOST,required"`
	Port     string `env:"DB_PORT,required"`
	Name     string `env:"DB_NAME,required"`
	User     string `env:"DB_USER,required"`
	Password string `env:"DB_PASSWORD,required"`
	SSLMode  string `env:"DB_SSLMODE" envDefault:"disable"`
}

func (c DBConfig) DSN() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		c.User, c.Password, c.Host, c.Port, c.Name, c.SSLMode)
}

func Load(logger *zap.Logger) (Config, error) {
	_ = godotenv.Load()

	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		logger.Warn("failed to parse config", zap.Error(err))
		return Config{}, err
	}

	logger.Info("loaded config")
	return cfg, nil
}
