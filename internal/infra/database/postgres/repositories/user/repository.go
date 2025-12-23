package user

import "belscourrsego/internal/infra/database/postgres"

type Repository struct {
	db postgres.Executor
}

func NewRepository(db postgres.Executor) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) InsertValue() error {
	return nil
}
