package psql

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"strconv"
)

type Institute interface {
	GetUniqueInstitutes(ctx context.Context, subjectName string, semester int, subjectType string, asc bool) ([]int, error)
	GetAllInstitutes(ctx context.Context, asc bool) ([]int, error)
}

func (d Db) GetUniqueInstitutes(ctx context.Context, subjectName string, semester int, subjectType string, asc bool) ([]int, error) {
	var err error
	//
	q := `select instistute_num from 
		( select *, row_number() over (partition by instistute_num order by id) as num from active_subject 
		 `
	varCount := 1
	input := []any{}
	if subjectName != "" || semester != 0 || subjectType != "" {
		q += " where "
	}

	if subjectName != "" {
		q += fmt.Sprintf(" name = $%d", varCount)
		varCount += 1
		input = append(input, subjectName)

	}
	if subjectType != "" {
		if varCount > 1 {
			q += " and "
		}
		q += fmt.Sprintf(" type_name = $%d", varCount)
		varCount += 1
		input = append(input, subjectType)

	}
	if semester != 0 {
		institute := strconv.Itoa(semester)
		if varCount > 1 {
			q += " and "
		}
		q += fmt.Sprintf(" semester_number = $%d", varCount)
		varCount += 1
		input = append(input, institute)

	}
	q += `) active_subject where num = 1 order by instistute_num`
	if !asc {
		q += "  desc"
	}
	rows, err := d.query(ctx, q, input...)
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
	rows, err := d.query(ctx, q)
	if err != nil {
		return nil, err
	}
	insts, err := pgx.CollectRows(rows, pgx.RowTo[int])
	if err != nil {
		return nil, err
	}
	return insts, nil

}
