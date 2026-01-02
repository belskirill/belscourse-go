package app

import (
	"belscourrsego/internal/infra/webhooks/payment_webhook"
	"belscourrsego/internal/usecase/payment"
	"context"

	"go.uber.org/zap"
)

type WebhookSubsystem struct {
	Queue payment.PaymentQueue
}

// BuildEmailSubsystem
func BuildPaymentSubsystem(
	ctx context.Context,
	logger *zap.Logger,
) (*WebhookSubsystem, error) {

	queue := payment_webhook.NewInMemoryPaymentQueue(100)

	for i := 0; i < 5; i++ {
		worker := payment_webhook.NewPaymentWorker(
			queue,
			logger,
		)

		worker.Start(ctx)
	}

	logger.Info("email subsystem started")

	return &WebhookSubsystem{
		Queue: queue,
	}, nil
}
