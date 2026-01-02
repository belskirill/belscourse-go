package app

import (
	"belscourrsego/internal/config"
	user4 "belscourrsego/internal/domain/user"
	"belscourrsego/internal/infra/crypto"
	"belscourrsego/internal/infra/database/postgres"
	"belscourrsego/internal/infra/database/postgres/repositories/user"
	"belscourrsego/internal/infra/jwt"
	user3 "belscourrsego/internal/interface/http/user"
	"belscourrsego/internal/usecase/notifications"
	"belscourrsego/internal/usecase/payment"
	user2 "belscourrsego/internal/usecase/user"
	"context"
	"database/sql"

	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

type userHandlers struct {
	user       *user3.UserHandler
	jwtService *jwt.Service
}

func buildUserHandlers(db *sql.DB, validate *validator.Validate, cfg config.Config, logger *zap.Logger) *userHandlers {
	repo := user.NewRepository(db)

	emailSubsystem, err := BuildEmailSubsystem(context.Background(), cfg, logger)
	if err != nil {
		panic(err)
	}

	paymentSybSystem, err := BuildPaymentSubsystem(context.Background(), logger)
	if err != nil {
		panic(err)
	}

	emailService := notifications.NewService(emailSubsystem.Queue)

	sendEml := user2.NewEmailSender(repo, emailService)

	webhook := payment.NewWebhookInput(paymentSybSystem.Queue)

	transaction := postgres.NewUnitOfWork(db)
	CryptoPassword := crypto.NewHashCrypto()
	ServicePassword := user4.NewServicePasswordPolicy(CryptoPassword)

	authService := jwt.NewTokenService(cfg.JWT.JWTSecret, cfg.JWT.JWTExpire)

	createSessionUC := user2.NewUseCaseCreate(repo, transaction, ServicePassword)

	loginSession := user2.NewUseCaseCreateSession(repo, ServicePassword, authService)

	return &userHandlers{
		user:       user3.NewHandler(createSessionUC, loginSession, validate, sendEml, webhook),
		jwtService: authService,
	}
}
