package user

import (
	"belscourrsego/internal/domain/user"
	"belscourrsego/internal/usecase/notifications"
	"context"
)

type EmailSender struct {
	repo user.GetUserByIDs
	send notifications.ChangePasswordNotifier
}

func NewEmailSender(
	repo user.GetUserByIDs,
	send notifications.ChangePasswordNotifier,
) EmailSender {
	return EmailSender{
		repo: repo,
		send: send,
	}
}

func (s EmailSender) SendEmailCode(ctx context.Context, userID int64) error {
	u, err := s.repo.GetUserByID(ctx, userID)
	if err != nil {
		return err
	}

	return s.send.SendChangePasswordCode(ctx, u.Email)
}
