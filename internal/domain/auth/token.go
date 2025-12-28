package auth

import (
	"time"
)

type TokenService interface {
	Generate(userID int64) (string, error)
	ParseGetById(token string) (int64, error)
}

type Claims struct {
	UserID int64
	Exp    time.Time
}
