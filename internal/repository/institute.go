package psql

import (
	"context"
	"github.com/jackc/pgx/v5"
)

type Institute interface {
	GetUniqueInstitutes(ctx context.Context, subjectName string, sem int, subjectType string, asc bool) ([]int, error)
	GetAllInstitutes(ctx context.Context, asc bool) ([]int, error)
}

func (d Db) GetUniqueInstitutes(ctx context.Context, subjectName string, sem int, subjectType string, asc bool) ([]int, error) {
	var err error
	//
	q := `select instistute_num from active_subject where name = $1 and semester_number = $2 and type_name = $3 order by type_name
		 `
	if !asc {
		q += " desc"
	}
	rows, err := d.client.Query(ctx, q, subjectName, sem, subjectType)
	if err != nil {
		return nil, err
	}
	insts, err := pgx.CollectRows(rows, pgx.RowTo[int])
	if err != nil {
		return nil, err
	}
	return insts, nil

}
func (d Db) GetAllInstitutes(ctx context.Context, asc bool) ([]int, error) {
	var err error
	//
	q := `select number from  instistute order by number
		 `
	if !asc {
		q += " desc"
	}
	rows, err := d.client.Query(ctx, q)
	if err != nil {
		return nil, err
	}
	insts, err := pgx.CollectRows(rows, pgx.RowTo[int])
	if err != nil {
		return nil, err
	}
	return insts, nil

}
