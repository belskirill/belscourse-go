package app

import (
	"belscourrsego/internal/config"
	"belscourrsego/internal/infra/email"
	"belscourrsego/internal/usecase/notifications"
	"context"

	"go.uber.org/zap"
)

type EmailSubsystem struct {
	Queue notifications.EmailQueue
}

// BuildEmailSubsystem
func BuildEmailSubsystem(
	ctx context.Context,
	cfg config.Config,
	logger *zap.Logger,
) (*EmailSubsystem, error) {

	// 1️⃣ SMTP sender (infra)
	smtpSender := email.NewSMTPSender(
		cfg.Email.Host,
		cfg.Email.Port,
		cfg.Email.User,
		cfg.Email.Pass,
		cfg.Email.From,
	)

	// 2️⃣ Queue (infra)
	queue := email.NewInMemoryEmailQueue(100)

	for i := 0; i < 5; i++ {
		worker := email.NewEmailWorker(
			queue,      // EmailQueue
			smtpSender, // EmailSender
			logger,
		)

		worker.Start(ctx)
	}

	//worker := email.NewEmailWorker(
	//	queue,      // EmailQueue
	//	smtpSender, // EmailSender
	//	logger,
	//)
	//
	//worker.Start(ctx)

	logger.Info("email subsystem started")

	return &EmailSubsystem{
		Queue: queue,
	}, nil
}
