package notifications

import "context"

type EmailSender interface {
	Send(to, subject, body string) error
}

type ChangePasswordNotifier interface {
	SendChangePasswordCode(
		ctx context.Context,
		email string,
	) error
}

type EmailQueue interface {
	Enqueue(ctx context.Context, task EmailTask) error
	Channel() <-chan EmailTask
}
