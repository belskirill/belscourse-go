package user

import (
	"context"
)

type CreateSessionRepo interface {
	InsertValue(ctx context.Context, req User) error
}
