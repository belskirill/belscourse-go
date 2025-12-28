package user

import (
	"belscourrsego/internal/domain/auth"
	"belscourrsego/internal/domain/user"
	"context"
)

type UseCaseCreateSession struct {
	repo          user.CreateSessionRepo
	domainService user.PasswordService
	auth          auth.TokenService
}

func NewUseCaseCreateSession(
	repo user.CreateSessionRepo,
	domainService user.PasswordService,
	auth auth.TokenService,
) *UseCaseCreateSession {
	return &UseCaseCreateSession{
		repo:          repo,
		domainService: domainService,
		auth:          auth,
	}
}

func (u *UseCaseCreateSession) CreateSession(ctx context.Context, req user.UserWithPassword) (string, error) {
	usr, err := u.repo.GetUserByEmailOrUsername(ctx, req)
	if err != nil {
		return "", err
	}

	if err := u.domainService.CompareService(usr.HashPassword, req.Password); err != nil {
		return "", user.New(user.ErrInvalidPassword, err)
	}

	token, err := u.auth.Generate(usr.ID)

	if err != nil {
		return "", err
	}

	return token, nil
}
