package user

import (
	"belscourrsego/internal/domain/user"
	"belscourrsego/internal/infra/database/postgres"
	"context"
)

type UseCaseCreate struct {
	repo          user.CreateRepo
	uow           *postgres.UnitOfWork
	domainService user.PasswordService
}

func NewUseCaseCreate(
	repo user.CreateRepo,
	uow *postgres.UnitOfWork,
	domainService user.PasswordService,
) *UseCaseCreate {
	return &UseCaseCreate{
		repo:          repo,
		uow:           uow,
		domainService: domainService,
	}
}

func (uc *UseCaseCreate) CreateUser(ctx context.Context, req user.UserCreate) (user.UserBase, error) {
	res, err := postgres.Do(ctx, uc.uow, func(ctx context.Context) (user.UserBase, error) {
		hash, err := uc.domainService.HashService(req.Password, 12)
		if err != nil {
			return user.UserBase{}, err
		}

		req.PasswordHash = hash

		res, err := uc.repo.InsertValue(ctx, req)
		if err != nil {
			return user.UserBase{}, err
		}

		return res, nil
	})

	if err != nil {
		return user.UserBase{}, err
	}

	return res, nil
}
