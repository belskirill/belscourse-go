package postgres

import (
	"context"
	"database/sql"
)

type SqlTxKey struct{}

func TxFromContext(ctx context.Context) (*sql.Tx, bool) {
	tx, ok := ctx.Value(SqlTxKey{}).(*sql.Tx)
	return tx, ok
}
