package email

import (
	"belscourrsego/internal/usecase/notifications"
	"context"
)

type InMemoryEmailQueue struct {
	ch chan notifications.EmailTask
}

func NewInMemoryEmailQueue(buffer int) *InMemoryEmailQueue {
	return &InMemoryEmailQueue{
		ch: make(chan notifications.EmailTask, buffer),
	}
}

func (q *InMemoryEmailQueue) Enqueue(
	ctx context.Context,
	task notifications.EmailTask,
) error {

	select {
	case q.ch <- task:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// Channel — даёт доступ воркеру читать задачи
func (q *InMemoryEmailQueue) Channel() <-chan notifications.EmailTask {
	return q.ch
}
