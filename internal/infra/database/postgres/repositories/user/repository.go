package user

import (
	"belscourrsego/internal/infra/database/postgres"
	"context"
	"database/sql"
	"errors"
)

type Repository struct {
	db postgres.Executor
}

func NewRepository(db postgres.Executor) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) InsertValue(ctx context.Context) error {
	const query = `
		INSERT INTO users (username, email, password_hash)
		VALUES ($1, $2, $3)
`

	if tx, ok := ctx.Value(postgres.SqlTxKey{}).(*sql.Tx); ok {
		_, err := tx.ExecContext(ctx, query, "test1", "test2", "test3")
		if err != nil {
			return err
		}
		return nil
	}
	return errors.New("transaction is required")
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
