package email

import (
	"belscourrsego/internal/usecase/notifications"
	"context"

	"go.uber.org/zap"
)

type EmailWorker struct {
	queue  notifications.EmailQueue
	sender notifications.EmailSender
	logger *zap.Logger
}

func NewEmailWorker(
	queue notifications.EmailQueue,
	sender notifications.EmailSender,
	logger *zap.Logger,
) *EmailWorker {
	return &EmailWorker{
		queue:  queue,
		sender: sender,
		logger: logger,
	}
}

func (w *EmailWorker) Start(ctx context.Context) {
	go func() {
		for {
			select {
			case task := <-w.queue.Channel():
				w.handle(task)

			case <-ctx.Done():
				w.logger.Info("email worker stopped")
				return
			}
		}
	}()
}

func (w *EmailWorker) handle(task notifications.EmailTask) {
	switch task.Template {

	case "change_password":
		subject := "Change password"
		body := "Your code: " + task.Code

		if err := w.sender.Send(task.To, subject, body); err != nil {
			w.logger.Error("email send failed",
				zap.String("to", task.To),
				zap.Error(err),
			)
		}
	}
}
