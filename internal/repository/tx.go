package psql

import (
	"context"
	"github.com/jackc/pgx/v5"
)

func (d Db) WithTransaction(ctx context.Context, f func(ctx context.Context) error) error {
	tx, err := d.client.Begin(ctx)
	defer tx.Commit(ctx)
	if err != nil {
		return err
	}
	ctx = context.WithValue(ctx, "tx", tx)
	err = f(ctx)
	if err != nil {
		tx.Rollback(ctx)
		return err
	}
	return err
}

func (d Db) queryRow(ctx context.Context, sql string, args ...any) pgx.Row {
	tx, ok := ctx.Value("tx").(pgx.Tx)
	if !ok {
		return d.client.QueryRow(ctx, sql, args...)
	}
	return tx.QueryRow(ctx, sql, args...)
}
func (d Db) query(ctx context.Context, sql string, args ...any) (pgx.Rows, error) {
	tx, ok := ctx.Value("tx").(pgx.Tx)
	if !ok {
		return d.client.Query(ctx, sql, args...)
	}
	return tx.Query(ctx, sql, args...)
}
