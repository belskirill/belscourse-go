package user

import "context"

type UseCreatorSession interface {
	CreateUser(ctx context.Context) error
}
