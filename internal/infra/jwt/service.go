package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Service struct {
	secret []byte
	ttl    time.Duration
}

func NewTokenService(secret string, ttl time.Duration) *Service {
	return &Service{
		secret: []byte(secret),
		ttl:    ttl,
	}
}

type claims struct {
	UserID int64 `json:"user_id"`
	jwt.RegisteredClaims
}

func (s *Service) Generate(userID int64) (string, error) {
	now := time.Now()

	c := claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(s.ttl)),
			IssuedAt:  jwt.NewNumericDate(now),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)

	return token.SignedString(s.secret)
}

func (s *Service) Parse(tokenStr string) (*claims, error) {
	token, err := jwt.ParseWithClaims(
		tokenStr,
		&claims{},
		func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("unexpected signing method")
			}

			return s.secret, nil
		},
	)

	if err != nil || !token.Valid {
		return nil, errors.New("invalid token")
	}

	c, ok := token.Claims.(*claims)
	if !ok {
		return nil, errors.New("invalid claims")
	}

	return c, nil
}

func (s *Service) ParseGetById(token string) (int64, error) {
	claims, err := s.Parse(token)

	if err != nil {
		return 0, err
	}

	return claims.UserID, nil
}
