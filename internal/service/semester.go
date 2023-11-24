package service

import (
	"context"
)

type Semester interface {
	GetAllSemesters(ctx context.Context, asc bool) ([]int, error)
	GetAllSemestersBySubjectName(ctx context.Context, subjectname string, asc bool) ([]int, error)
}

func (s service) GetAllSemesters(ctx context.Context, asc bool) ([]int, error) {
	sems, err := s.rep.GetAllSemesters(ctx, asc)
	if err != nil {
		return nil, err
	}
	return sems, nil
}
func (s service) GetAllSemestersBySubjectName(ctx context.Context, subjectname string, asc bool) ([]int, error) {
	sems, err := s.rep.GetAllSemestersBySubjectName(ctx, subjectname, asc)
	if err != nil {
		return nil, err
	}
	return sems, nil
}
