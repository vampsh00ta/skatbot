package psql

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"skat_bot/internal/repository/models"
)

type Stats interface {
	AddLike(ctx context.Context, variantId int, username string) error
	DeleteLike(ctx context.Context, variantId int, username string) error
	CheckLike(ctx context.Context, variantId int, username string) ([]models.Stats, error)
}

func PgErrorCode(err error) string {

	pgerr, ok := err.(*pgconn.PgError)
	if !ok {
		return ""
	}
	return pgerr.Code
}

func (db Db) CheckLike(ctx context.Context, variantId int, username string) ([]models.Stats, error) {
	var err error
	q := `select * from  variant_likes where   variant_id = $1 and username = $2`

	rows, err := db.query(ctx, q, variantId, username)

	stat, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[models.Stats])
	if err != nil {
		return nil, err
	}
	return stat, nil
}
func (db Db) AddLike(ctx context.Context, variantId int, username string) error {
	var err error
	q := `insert into  variant_likes (variant_id,username,l) values($1,$2,1) returning variant_id;`
	if err = db.queryRow(ctx, q, variantId, username).Scan(&variantId); err != nil {
		return err

	}
	return nil
}
func (db Db) DeleteLike(ctx context.Context, variantId int, username string) error {
	var err error
	q := `delete from    variant_likes where variant_id = $1 and username  = $2 and l = 1 
		  returning variant_id`
	if err = db.queryRow(ctx, q, variantId, username).Scan(&variantId); err != nil {
		return err

	}
	return nil
}
