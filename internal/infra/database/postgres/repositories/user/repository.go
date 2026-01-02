package user

import (
	"belscourrsego/internal/domain/user"
	"belscourrsego/internal/infra/database/postgres"
	"context"
	"database/sql"
	"errors"

	"github.com/lib/pq"
)

type Repository struct {
	db postgres.Executor
}

func NewRepository(db postgres.Executor) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) InsertValue(ctx context.Context, req user.UserCreate) (user.UserBase, error) {
	const query = `
		INSERT INTO users (username, email, password_hash)
		VALUES ($1, $2, $3)
		RETURNING id, username, email
`
	var response user.UserBase

	if err := r.db.QueryRowContext(ctx, query, req.Username, req.Email, req.PasswordHash).Scan(&response.ID, &response.Username, &response.Email); err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) {
			if pqErr.Code == "23505" {
				return user.UserBase{}, user.New(
					user.ErrUserAlreadyExists,
					err,
				)
			}
		}

		return user.UserBase{}, err
	}

	return response, nil
}

func (r *Repository) GetUserByEmailOrUsername(ctx context.Context, usr user.UserWithPassword) (user.UserWithPassword, error) {
	const query = `
		SELECT id, username, email, password_hash
		FROM users
		WHERE username = $1 or email = $2
`
	var userPass user.UserWithPassword

	if err := r.db.QueryRowContext(ctx, query, usr.Username, usr.Email).Scan(&userPass.ID, &userPass.Username, &userPass.Email, &userPass.HashPassword); err != nil {
		if err == sql.ErrNoRows {
			return user.UserWithPassword{}, user.New(user.ErrUserNotFound, err)
		}

		return user.UserWithPassword{}, err
	}

	return userPass, nil
}

func (r *Repository) GetUserByID(ctx context.Context, id int64) (user.UserBase, error) {
	const query = `
SELECT email
FROM users
WHERE id = $1`
	var response user.UserBase
	if err := r.db.QueryRowContext(ctx, query, id).Scan(&response.Email); err != nil {
		if err == sql.ErrNoRows {
			return user.UserBase{}, user.New(user.ErrUserNotFound, err)
		}

		return user.UserBase{}, err
	}

	return response, nil
}
