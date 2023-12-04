package service

import (
	"context"
	"skat_bot/internal/repository/models"
)

type Skat interface {
	AddSkat(ctx context.Context, subject models.Subject) (models.Subject, error)
}

func (s service) AddSkat(ctx context.Context, subject models.Subject) (models.Subject, error) {
	var sub models.Subject
	err := s.rep.WithTransaction(ctx, func(ctx context.Context) error {

		sub, err := s.rep.AddOrGetSubject(ctx, subject)
		if err != nil {
			return err
		}
		sub.Variants[0].SubjectId = sub.Id
		if _, err := s.rep.AddVariant(ctx, sub.Variants[0]); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return models.Subject{}, err
	}
	return sub, nil
}
