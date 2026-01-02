package payment_webhook

import (
	"belscourrsego/internal/usecase/payment"
	"context"
)

type QueueWebhookPayment struct {
	ch chan payment.AccessPayment
}

func NewInMemoryPaymentQueue(buffer int) *QueueWebhookPayment {
	return &QueueWebhookPayment{
		ch: make(chan payment.AccessPayment, buffer),
	}
}

func (qw *QueueWebhookPayment) Enqueue(ctx context.Context, data payment.AccessPayment) error {
	select {
	case qw.ch <- data:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func (qw *QueueWebhookPayment) Channel() <-chan payment.AccessPayment {
	return qw.ch
}
