package user

import (
	"belscourrsego/internal/domain/common"
	"belscourrsego/internal/domain/user"
	"context"
)

type UseCaseCreateSession struct {
	repo          user.CreateSessionRepo
	transaction   common.UnitOfWorker
	domainService user.PasswordService
}

func NewUseCaseCreateSession(
	repo user.CreateSessionRepo,
	transaction common.UnitOfWorker,
	domainService user.PasswordService,
) *UseCaseCreateSession {
	return &UseCaseCreateSession{
		repo:          repo,
		transaction:   transaction,
		domainService: domainService,
	}
}

func (uc *UseCaseCreateSession) CreateUser(ctx context.Context, req user.CreateUserRequest) error {
	if err := uc.transaction.Do(ctx, func(ctx context.Context) error {

		hash, err := uc.domainService.HashService(req.Password, 12)
		if err != nil {
			return err
		}

		UserWithHash := user.User{
			Username:     req.Username,
			Email:        req.Email,
			PasswordHash: hash,
		}

		return uc.repo.InsertValue(ctx, UserWithHash)
	}); err != nil {
		return err
	}

	return nil
}
