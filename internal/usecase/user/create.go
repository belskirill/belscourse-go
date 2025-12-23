package user

import "belscourrsego/internal/domain/user"

type UseCaseCreateSession struct {
	repo user.CreateSessionRepo
}

func NewUseCaseCreateSession(repo user.CreateSessionRepo) *UseCaseCreateSession {
	return &UseCaseCreateSession{repo: repo}
}

func (uc *UseCaseCreateSession) CreateUser() error {
	return uc.repo.InsertValue()
}
