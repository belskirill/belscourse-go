package common

import "context"

type UnitOfWorker interface {
	Do(ctx context.Context, fn func(ctx context.Context) error) error
}
