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

func (uc *UseCaseCreateSession) CreateUser(ctx context.Context) error {
	if err := uc.transaction.Do(ctx, func(ctx context.Context) error {
		return uc.repo.InsertValue(ctx)
	}); err != nil {
		return err
	}

	return nil

	//if err := uc.transaction.Do(ctx, func(ctx context.Context) error {
	//	hash, err := uc.domainService.Hash()
	//	if err != nil {
	//		return err
	//	}
	//
	//
	//})
}
