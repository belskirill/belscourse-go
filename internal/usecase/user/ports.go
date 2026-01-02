package user

import (
	"belscourrsego/internal/domain/user"
	"context"
)

type UseCreator interface {
	CreateUser(ctx context.Context, req user.UserCreate) (user.UserBase, error)
}

type CreateUserSession interface {
	CreateSession(ctx context.Context, req user.UserWithPassword) (string, error)
}

type SendEmailerCode interface {
	SendEmailCode(ctx context.Context, userID int64) error
}

type WebHook interface {
}
