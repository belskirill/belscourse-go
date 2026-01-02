package notifications

import (
	"belscourrsego/internal/domain/common/codegen"
	"context"
)

type Service struct {
	queue EmailQueue
}

func NewService(queue EmailQueue) *Service {
	return &Service{
		queue: queue,
	}
}

func (s *Service) SendChangePasswordCode(
	ctx context.Context,
	email string,
) error {

	code, err := codegen.Generate4DigitCode()
	if err != nil {
		return err
	}

	task := EmailTask{
		To:         email,
		Template:   "change_password",
		Code:       code,
		TTLMinutes: 15,
	}

	return s.queue.Enqueue(ctx, task)
}
