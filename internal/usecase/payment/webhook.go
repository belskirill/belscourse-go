package payment

import "context"

type Webhook struct {
	wb PaymentQueue
}

func NewWebhookInput(wb PaymentQueue) *Webhook {
	return &Webhook{
		wb: wb,
	}
}

func (wh *Webhook) WebhookAccessPayment(ctx context.Context) error {
	return nil
}
