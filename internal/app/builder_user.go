package app

import (
	user4 "belscourrsego/internal/domain/user"
	"belscourrsego/internal/infra/crypto"
	"belscourrsego/internal/infra/database/postgres"
	"belscourrsego/internal/infra/database/postgres/repositories/user"
	user3 "belscourrsego/internal/interface/http/user"
	user2 "belscourrsego/internal/usecase/user"
	"database/sql"
)

type userHandlers struct {
	createSession *user3.UserHandler
}

func buildUserHandlers(db *sql.DB) *userHandlers {
	repo := user.NewRepository(db)

	transaction := postgres.NewUnitOfWork(db)
	CryptoPassword := crypto.NewHashCrypto()
	ServicePassword := user4.NewServicePasswordPolicy(CryptoPassword)

	createSessionUC := user2.NewUseCaseCreateSession(repo, transaction, ServicePassword)

	return &userHandlers{
		createSession: user3.NewHandler(createSessionUC),
	}
}
