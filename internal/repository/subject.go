package psql

import (
	"context"
	"github.com/jackc/pgx/v5"
	"skat_bot/internal/repository/models"
)

type Subject interface {
	AddSubject(ctx context.Context, subject models.Subject) (models.Subject, error)
	GetSubject(ctx context.Context, subject models.Subject) (models.Subject, error)
	GetAllSubjects(ctx context.Context) ([]models.Subject, error)
	GetUniqueSubjects(ctx context.Context, asc bool) ([]models.Subject, error)
	GetSubjectsByName(ctx context.Context, name string, asc bool) ([]models.Subject, error)
	GetAllSubjectNames(ctx context.Context, asc bool) ([]models.Subject, error)
	GetUniqueSubjectTypes(ctx context.Context, subjectName string, sem int, asc bool) ([]models.Subject, error)
	GetAllSubjectTypes(ctx context.Context, asc bool) ([]models.Subject, error)

	//AddOrGetSubject(ctx context.Context, subject models.Subject) ([]int, error)
}

func (d Db) AddSubject(ctx context.Context, subject models.Subject) (models.Subject, error) {
	var err error
	//
	q := `insert into active_subject (name,semester_number,instistute_num,type_name)
			values ($1,$2,$3,$4) returning id 
		 `

	if err = d.client.QueryRow(ctx, q,
		subject.Name,
		subject.Semester,
		subject.InstistuteNum,

		subject.TypeName).Scan(&subject.Id); err != nil {

		return models.Subject{}, err
	}
	return subject, nil

}

func (d Db) GetAllSubjects(ctx context.Context) ([]models.Subject, error) {
	var err error
	//
	q := `select   * from  subject
		 `

	rows, err := d.client.Query(ctx, q)
	if err != nil {
		return nil, err
	}
	subjects, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[models.Subject])
	if err != nil {
		return nil, err
	}

	return subjects, nil
}

func (d Db) GetUniqueSubjects(ctx context.Context, asc bool) ([]models.Subject, error) {
	var err error
	//
	q := `
select id,name,semester_number,instistute_num,type_name from 
	( select *, row_number() over (partition by name order by id) as num from active_subject ) active_subject 
                                            where num = 1
		 `

	if asc {
		q += " order by name"
	}

	rows, err := d.client.Query(ctx, q)
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
	rows, err := d.client.Query(ctx, q, name)
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
	rows, err := d.client.Query(ctx, q)
	if err != nil {
		return nil, err
	}
	subjects, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[models.Subject])
	if err != nil {
		return nil, err
	}

	return subjects, nil
}
func (d Db) GetUniqueSubjectTypes(ctx context.Context, subjectName string, sem int, asc bool) ([]models.Subject, error) {
	var err error
	//
	q := `select type_name from active_subject where name = $1 and semester_number = $2 order by type_name
		 `
	if !asc {
		q += " desc"
	}
	rows, err := d.client.Query(ctx, q, subjectName, sem)
	if err != nil {
		return nil, err
	}
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
	if err := d.client.QueryRow(ctx, q, subject.Name, subject.Semester, subject.TypeName, subject.InstistuteNum).
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
	rows, err := d.client.Query(ctx, q)
	if err != nil {
		return nil, err
	}
	subjectsNames, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[models.Subject])
	if err != nil {
		return nil, err
	}
	return subjectsNames, nil
}
