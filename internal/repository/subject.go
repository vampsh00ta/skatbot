package psql

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"skat_bot/internal/repository/models"
	"strconv"
)

type Subject interface {
	AddSubject(ctx context.Context, subject models.Subject) (models.Subject, error)
	GetSubject(ctx context.Context, subject models.Subject) (models.Subject, error)
	GetAllSubjects(ctx context.Context) ([]models.Subject, error)
	GetUniqueSubjects(ctx context.Context, instituteNum, semester int, subjectType string, asc bool) ([]models.Subject, error)
	GetSubjectsByName(ctx context.Context, name string, asc bool) ([]models.Subject, error)
	GetAllSubjectNames(ctx context.Context, asc bool) ([]models.Subject, error)
	GetUniqueSubjectTypes(ctx context.Context, subjectName string, semester, instituteNum int, asc bool) ([]models.Subject, error)
	GetAllSubjectTypes(ctx context.Context, asc bool) ([]models.Subject, error)
	//AddOrGetSubject(ctx context.Context, subject models.Subject) ([]int, error)
}

func (d Db) AddSubject(ctx context.Context, subject models.Subject) (models.Subject, error) {
	var err error
	//

	q := `insert into active_subject (name,semester_number,instistute_num,type_name)
			values ($1,$2,$3,$4) returning id 
		 `

	if err = d.queryRow(ctx, q,
		subject.Name,
		subject.Semester,
		subject.InstistuteNum,

		subject.TypeName,
	).Scan(&subject.Id); err != nil {

		return models.Subject{}, err
	}
	return subject, nil

}

func (d Db) GetAllSubjects(ctx context.Context) ([]models.Subject, error) {
	var err error
	//
	q := `select   * from  subject
		 `

	rows, err := d.query(ctx, q)
	if err != nil {
		return nil, err
	}
	subjects, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[models.Subject])
	if err != nil {
		return nil, err
	}

	return subjects, nil
}

func (d Db) GetUniqueSubjects(ctx context.Context, instituteNum, semester int, subjectType string, asc bool) ([]models.Subject, error) {
	var err error
	//
	q := `select id,name,semester_number,instistute_num,type_name from 
		( select *, row_number() over (partition by name order by id) as num from active_subject 
		 `
	varCount := 1
	input := []any{}
	if instituteNum != 0 || semester != 0 || subjectType != "" {
		q += " where "
	}

	if instituteNum != 0 {
		institute := strconv.Itoa(instituteNum)
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
	if semester != 0 {
		institute := strconv.Itoa(semester)
		if varCount > 1 {
			q += " and "
		}
		q += fmt.Sprintf(" semester_number = $%d", varCount)
		varCount += 1
		input = append(input, institute)

	}
	q += `) active_subject where num = 1 order by name`
	if !asc {
		q += "desc "
	}
	rows, err := d.query(ctx, q, input...)
	if err != nil {
		return nil, err
	}
	subjects, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[models.Subject])
	if err != nil {
		return nil, err
	}

	return subjects, nil
}
func (d Db) GetSubjectsByName(ctx context.Context, name string, asc bool) ([]models.Subject, error) {

	var err error
	//
	q := `select   * from  subject where name = $1 
		 `
	if asc {
		q += " order by semester "
	}
	rows, err := d.query(ctx, q, name)
	if err != nil {
		return nil, err
	}
	subjects, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[models.Subject])
	if err != nil {
		return nil, err
	}

	return subjects, nil
}

func (d Db) GetAllSubjectTypes(ctx context.Context, asc bool) ([]models.Subject, error) {
	var err error
	//
	q := `select   name as type_name from  subject_type order by name
		 `
	if !asc {
		q += " desc "
	}
	rows, err := d.query(ctx, q)
	if err != nil {
		return nil, err
	}
	subjects, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[models.Subject])
	if err != nil {
		return nil, err
	}

	return subjects, nil
}
func (d Db) GetUniqueSubjectTypes(ctx context.Context, subjectName string, semester, instituteNum int, asc bool) ([]models.Subject, error) {
	var err error
	//
	varCount := 1
	q := `
select type_name from 
	( select *, row_number() over (partition by type_name order by id) as num from active_subject 
		 `
	if instituteNum != 0 || subjectName != "" || semester != 0 {
		q += " where "
	}
	input := []any{}
	if subjectName != "" {
		q += fmt.Sprintf(" name = $%d", varCount)
		varCount += 1
		input = append(input, subjectName)
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
	if semester != 0 {
		sem := strconv.Itoa(semester)

		if varCount > 1 {
			q += " and "
		}
		q += fmt.Sprintf(" semester_number = $%d", varCount)
		varCount += 1
		input = append(input, sem)

	}
	q += `) active_subject where num = 1`

	q += " order by  semester_number"
	if !asc {
		q += " desc "
	}
	rows, err := d.query(ctx, q, input...)
	types, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[models.Subject])
	if err != nil {
		return nil, err
	}

	return types, nil
}

func (d Db) GetSubject(ctx context.Context, subject models.Subject) (models.Subject, error) {
	//
	q := `select   id from  active_subject 
            where name = $1 and semester_number = $2 and type_name = $3 and instistute_num = $4
		 `
	if err := d.queryRow(ctx, q, subject.Name, subject.Semester, subject.TypeName, subject.InstistuteNum).
		Scan(&subject.Id); err != nil {
		return models.Subject{}, err
	}

	return subject, nil
}
func (d Db) GetAllSubjectNames(ctx context.Context, asc bool) ([]models.Subject, error) {
	var err error
	//
	q := `select   name   from  subject  order by name
		 `
	if !asc {
		q += "  desc "
	}
	rows, err := d.query(ctx, q)
	if err != nil {
		return nil, err
	}
	subjectsNames, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[models.Subject])
	if err != nil {
		return nil, err
	}
	return subjectsNames, nil
}
