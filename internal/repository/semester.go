package psql

import (
	"context"
	"github.com/jackc/pgx/v5"
)

type Semester interface {
	GetAllSemesters(ctx context.Context, asc bool) ([]int, error)
	GetAllSemestersBySubjectName(ctx context.Context, subjectname string, asc bool) ([]int, error)
}

func (d Db) GetAllSemestersBySubjectName(ctx context.Context, subjectname string, asc bool) ([]int, error) {
	var err error
	//
	q := `select   distinct semester_number from  active_subject where name = $1 order by  semester_number
		 `
	if !asc {
		q += " desc "
	}
	rows, err := d.client.Query(ctx, q, subjectname)
	if err != nil {
		return nil, err
	}
	semesters, err := pgx.CollectRows(rows, pgx.RowTo[int])
	if err != nil {
		return nil, err
	}

	return semesters, nil
}
func (d Db) GetAllSemesters(ctx context.Context, asc bool) ([]int, error) {
	var err error
	//
	q := `select   number from  semester order by number
		 `
	if !asc {
		q += " desc "
	}
	rows, err := d.client.Query(ctx, q)
	if err != nil {
		return nil, err
	}
	semesters, err := pgx.CollectRows(rows, pgx.RowTo[int])
	if err != nil {
		return nil, err
	}

	return semesters, nil
}
