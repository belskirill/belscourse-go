package app

import (
	"belscourrsego/internal/infra/database/postgres"
	"belscourrsego/internal/infra/database/postgres/repositories/user"
	user3 "belscourrsego/internal/interface/http/user"
	user2 "belscourrsego/internal/usecase/user"
)

type userHandlers struct {
	createSession *user3.UserHandler
}

func buildUserHandlers(db postgres.Executor) *userHandlers {
	repo := user.NewRepository(db)

	createSessionUC := user2.NewUseCaseCreateSession(repo)

	return &userHandlers{
		createSession: user3.NewHandler(createSessionUC),
	}
}
