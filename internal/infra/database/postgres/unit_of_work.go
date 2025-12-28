package postgres

import (
	"context"
	"database/sql"
)

type UnitOfWork struct {
	db *sql.DB
}

func NewUnitOfWork(db *sql.DB) *UnitOfWork {
	return &UnitOfWork{db: db}
}

func Do[T any](
	ctx context.Context,
	uow *UnitOfWork,
	fn func(ctx context.Context) (T, error),
) (result T, err error) {
	tx, err := uow.db.BeginTx(ctx, nil)
	if err != nil {
		return result, err
	}

	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback()
			panic(p)
		}

		if err != nil {
			_ = tx.Rollback()
		}
	}()

	ctxWithTx := context.WithValue(ctx, SqlTxKey{}, tx)

	result, err = fn(ctxWithTx)

	if err != nil {
		return result, err
	}

	if err = tx.Commit(); err != nil {
		return result, err
	}

	return result, nil
}
