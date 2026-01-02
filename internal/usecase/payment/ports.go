package payment

import "context"

type PaymentQueue interface {
	Enqueue(ctx context.Context, data AccessPayment) error
	Channel() <-chan AccessPayment
}
