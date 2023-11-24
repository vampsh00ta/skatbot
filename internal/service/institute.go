package service

import "context"

type Institute interface {
	GetUniqueInstitutes(ctx context.Context, subjectName string, sem int, subjectType string, asc bool) ([]int, error)
	GetAllInstitutes(ctx context.Context, asc bool) ([]int, error)
}

func (s service) GetUniqueInstitutes(ctx context.Context, subjectName string, sem int, subjectType string, asc bool) ([]int, error) {
	insts, err := s.rep.GetUniqueInstitutes(ctx, subjectName, sem, subjectType, asc)
	if err != nil {
		return nil, err
	}
	return insts, nil
}
func (s service) GetAllInstitutes(ctx context.Context, asc bool) ([]int, error) {
	insts, err := s.rep.GetAllInstitutes(ctx, asc)
	if err != nil {
		return nil, err
	}
	return insts, nil

}
