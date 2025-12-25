package user

import (
	"belscourrsego/internal/domain/user"
	"context"
)

type UseCreatorSession interface {
	CreateUser(ctx context.Context, req user.CreateUserRequest) error
}
