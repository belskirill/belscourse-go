package user

import (
	"belscourrsego/internal/domain/user"
	"belscourrsego/internal/infra/database/postgres"
	"context"
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

func (r *Repository) InsertValue(ctx context.Context, req user.User) error {
	const query = `
		INSERT INTO users (username, email, password_hash)
		VALUES ($1, $2, $3)
`

	if _, err := r.db.ExecContext(ctx, query, req.Username, req.Email, req.PasswordHash); err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) {
			if pqErr.Code == "23505" {
				return user.New(
					user.ErrUserAlreadyExists,
					err,
					nil,
					"Такой пользователь уже существует!",
				)
			}
		}
	}
	return nil

	//if tx, ok := postgres.TxFromContext(ctx); ok {
	//	_, err := tx.ExecContext(ctx, query, "test1", "test2", "test3")
	//	if err != nil {
	//		return err
	//	}
	//	return nil
	//}

}

//func (r *Repository) SelectValue(ctx context.Context, value string) (*domain.User, error) {
//	const query = `SELECT * FROM users WHERE email = $1 or username = $1`
//	row := r.db.QueryRow(query, value)
//
//	var user domain.User
//	if err := row.Scan(&user); err != nil {
//		return nil, err
//	}
//
//	return &user, nil
//}
