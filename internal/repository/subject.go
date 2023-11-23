package psql

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"skat_bot/internal/repository/models"
)

type Subject interface {
	AddSubject(ctx context.Context, subject models.Subject) (int, error)
	GetSubjectTypeByName(ctx context.Context, subjectT string) (models.SubjectType, error)
	GetAllSubjects(ctx context.Context) ([]models.Subject, error)
	GetAllSubjectsOrderByName(ctx context.Context, asc bool) ([]models.Subject, error)
}

func (d Db) AddSubject(ctx context.Context, subject models.Subject) (int, error) {
	var err error
	//
	q := `insert into subject (name,semester,type)
			values ($1,$2,$3) returning id 
		 `

	if err = d.client.QueryRow(ctx, q,
		subject.Name,
		subject.Semester,
		subject.Type).Scan(&subject.Id); err != nil {

		return 0, err
	}
	return subject.Id, nil

}
func (d Db) GetSubjectTypeByName(ctx context.Context, subjectT string) (models.SubjectType, error) {
	q := `select * from subject_type where name = $1 
		 `
	fmt.Println(q)
	var subjectType models.SubjectType
	if err := d.client.QueryRow(ctx, q, subjectT).Scan(&subjectType.Id, &subjectType.Name); err != nil {
		return models.SubjectType{}, err
	}
	return subjectType, nil
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
	subjects, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Subject])
	if err != nil {
		return nil, err
	}

	return subjects, nil
}

func (d Db) GetAllSubjectsOrderByName(ctx context.Context, asc bool) ([]models.Subject, error) {
	var err error
	//
	q := `select   * from  subject order by name
		 `
	if !asc {
		q += " desc"
	}

	rows, err := d.client.Query(ctx, q)
	if err != nil {
		return nil, err
	}
	subjects, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Subject])
	if err != nil {
		return nil, err
	}

	return subjects, nil
}
