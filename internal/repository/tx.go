package psql

import (
	"context"
	"fmt"
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
		fmt.Println(5, tx)
		tx.Rollback(ctx)
		return err
	}
	return err
}

func (d Db) getTransaction(ctx context.Context) pgx.Tx {
	tx := ctx.Value("tx").(pgx.Tx)
	return tx
}
