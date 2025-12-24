package user

import (
	domain "belscourrsego/internal/domain/user"
	"belscourrsego/internal/infra/database/postgres"
)

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

func (r *Repository) SelectValue(value string) (*domain.User, error) {
	const query = `SELECT * FROM users WHERE email = $1 or username = $1`
	row := r.db.QueryRow(query, value)

	var user domain.User
	if err := row.Scan(&user); err != nil {
		return nil, err
	}

	return &user, nil
}
