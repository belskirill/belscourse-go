package payment_webhook

import (
	"belscourrsego/internal/usecase/payment"
	"context"

	"go.uber.org/zap"
)

type PaymentWorker struct {
	queue  payment.PaymentQueue
	logger *zap.Logger
}

func NewPaymentWorker(queue payment.PaymentQueue, logger *zap.Logger) *PaymentWorker {
	return &PaymentWorker{
		queue:  queue,
		logger: logger,
	}
}

func (w *PaymentWorker) Start(ctx context.Context) {
	go func() {
		for {
			select {
			case task := <-w.queue.Channel():
				w.handle(task)
			case <-ctx.Done():
				w.logger.Warn("payment worker stopped")
				return
			}
		}
	}()
}

func (w *PaymentWorker) handle(task payment.AccessPayment) {
	// dsdsdsdsd
}
