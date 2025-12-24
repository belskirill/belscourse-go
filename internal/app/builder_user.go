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

func buildUserHandlers(db postgres.Executor, sqldb *sql.DB) *userHandlers {
	repo := user.NewRepository(db)

	transaction := postgres.NewUnitOfWork(sqldb)
	CryptoPassword := crypto.NewHashCrypto()
	ServicePasssword := user4.NewServicePasswordPolicy(CryptoPassword)

	createSessionUC := user2.NewUseCaseCreateSession(repo, transaction, ServicePasssword)

	return &userHandlers{
		createSession: user3.NewHandler(createSessionUC),
	}
}
