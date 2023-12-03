package service

import (
	"context"
	"fmt"
	"skat_bot/internal/repository/models"
)

type Skat interface {
	AddSkat(ctx context.Context, subject models.Subject) (models.Subject, error)
}

func (s service) AddSkat(ctx context.Context, subject models.Subject) (models.Subject, error) {
	sub, err := s.AddOrGetSubject(ctx, subject)
	if err != nil {
		return models.Subject{}, err
	}
	sub.Variants[0].SubjectId = sub.Id
	if err := s.AddVariant(ctx, sub.Variants[0]); err != nil {
		return models.Subject{}, err
	}
	fmt.Println(sub.Id)
	return sub, nil
}
