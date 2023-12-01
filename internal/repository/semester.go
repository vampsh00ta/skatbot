package psql

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"strconv"
)

type Semester interface {
	GetAllSemesters(ctx context.Context, asc bool) ([]int, error)
	GetAllSemestersBySubjectName(ctx context.Context, subjectname string, asc bool) ([]int, error)
	GetUniqueSemesters(ctx context.Context, subjectname string, instituteNum int, subjectType string, asc bool) ([]int, error)
}

func (d Db) GetUniqueSemesters(ctx context.Context, subjectname string, instituteNum int, subjectType string, asc bool) ([]int, error) {
	var err error
	//
	varCount := 1
	q := `
select semester_number from 
	( select *, row_number() over (partition by semester_number order by id) as num from active_subject 
		 `
	if instituteNum != 0 || subjectname != "" || subjectType != "" {
		q += " where "
	}
	input := []any{}
	if subjectname != "" {
		q += fmt.Sprintf(" name = $%d", varCount)
		varCount += 1
		input = append(input, subjectname)
	}

	if instituteNum != 0 {
		institute := strconv.Itoa(instituteNum)
		if varCount > 1 {
			q += " and "
		}
		q += fmt.Sprintf(" instistute_num = $%d", varCount)
		varCount += 1
		input = append(input, institute)

	}
	if subjectType != "" {
		if varCount > 1 {
			q += " and "
		}
		q += fmt.Sprintf(" type_name = $%d", varCount)
		varCount += 1
		input = append(input, subjectType)

	}
	q += `) active_subject where num = 1`

	q += " order by  semester_number"
	if !asc {
		q += " desc "
	}
	rows, err := d.client.Query(ctx, q, input...)
	if err != nil {
		return nil, err
	}
	semesters, err := pgx.CollectRows(rows, pgx.RowTo[int])
	if err != nil {
		return nil, err
	}

	return semesters, nil
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
