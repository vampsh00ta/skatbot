package service

import (
	"context"
	"skat_bot/internal/repository/models"
)

type Subject interface {
	GetAllSubjectsOrderByName(ctx context.Context, asc bool) ([]models.Subject, error)
}

func (s service) GetAllSubjectsOrderByName(ctx context.Context, asc bool) ([]models.Subject, error) {
	subjects, err := s.rep.GetAllSubjectsOrderByName(ctx, true)
	if err != nil {
		return nil, err
	}
	return subjects, nil

}
