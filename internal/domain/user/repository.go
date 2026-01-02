package user

import (
	"context"
)

type CreateRepo interface {
	InsertValue(ctx context.Context, req UserCreate) (UserBase, error)
}
type CreateSessionRepo interface {
	GetUserByEmailOrUsername(ctx context.Context, usr UserWithPassword) (UserWithPassword, error)
}

type GetUserByIDs interface {
	GetUserByID(ctx context.Context, id int64) (UserBase, error)
}
