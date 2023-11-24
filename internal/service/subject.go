package service

import (
	"context"
	"skat_bot/internal/repository/models"
)

type Subject interface {
	GetAllSubjectsOrderByName(ctx context.Context, asc bool) ([]models.Subject, error)
	GetSubjectsByName(ctx context.Context, name string) ([]models.Subject, error)
	GetAllSubjectTypes(ctx context.Context, asc bool) ([]models.SubjectType, error)
	GetUniqueSubjects(ctx context.Context, asc bool) ([]models.Subject, error)
	GetAllSubjectNames(ctx context.Context, asc bool) ([]string, error)
	GetSubject(ctx context.Context, subject models.Subject) (models.Subject, error)
	GetUniqueSubjectTypes(ctx context.Context, subjectName string, sem int, asc bool) ([]string, error)
	GetUniqueInstitutes(ctx context.Context, subjectName string, sem int, subjectType string, asc bool) ([]int, error)

	AddOrGetSubject(ctx context.Context, subject models.Subject) (models.Subject, error)
}

func (s service) GetAllSubjectsOrderByName(ctx context.Context, asc bool) ([]models.Subject, error) {
	return nil, nil
}
func (s service) GetAllSubjectNames(ctx context.Context, asc bool) ([]string, error) {
	subjects, err := s.rep.GetAllSubjectNames(ctx, asc)
	if err != nil {
		return nil, err
	}
	return subjects, nil
}
func (s service) GetUniqueSubjects(ctx context.Context, asc bool) ([]models.Subject, error) {
	subjects, err := s.rep.GetUniqueSubjects(ctx, true)
	if err != nil {
		return nil, err
	}
	return subjects, nil

}
func (s service) GetSubjectsByName(ctx context.Context, name string) ([]models.Subject, error) {
	subjects, err := s.rep.GetSubjectsByName(ctx, name, true)
	if err != nil {
		return nil, err
	}
	return subjects, nil

}

func (s service) GetAllSubjectTypes(ctx context.Context, asc bool) ([]models.SubjectType, error) {
	subjectTypes, err := s.rep.GetAllSubjectTypes(ctx, asc)
	if err != nil {
		return nil, err
	}
	return subjectTypes, nil
}
func (s service) GetSubject(ctx context.Context, subject models.Subject) (models.Subject, error) {
	subject, err := s.rep.GetSubject(ctx, subject)
	if err != nil {
		return models.Subject{}, err
	}
	return subject, nil
}
func (s service) GetUniqueSubjectTypes(ctx context.Context, subjectName string, sem int, asc bool) ([]string, error) {
	types, err := s.rep.GetUniqueSubjectTypes(ctx, subjectName, sem, asc)
	if err != nil {
		return nil, err
	}
	return types, nil
}

func (s service) AddOrGetSubject(ctx context.Context, subject models.Subject) (models.Subject, error) {
	subj, err := s.rep.GetSubject(ctx, subject)
	if err != nil && err.Error() != "no rows in result set" {
		return models.Subject{}, err
	}
	id, err := s.rep.AddSubject(ctx, subject)
	if err != nil {
		return models.Subject{}, err
	}
	subj.Id = id
	return subj, nil
}
