package service

import (
	"context"
	"fmt"
	"skat_bot/internal/repository/models"
)

type Variant interface {
	AddVariant(ctx context.Context, variant models.Variant) error
	GetVariantsBySubject(ctx context.Context, subject models.Subject) ([]models.Variant, error)
	GetVariantTypes(ctx context.Context) ([]models.Variant, error)
}

func (s service) AddVariant(ctx context.Context, variant models.Variant) error {
	if _, err := s.rep.AddVariant(ctx, variant); err != nil {
		return err
	}
	return nil
}
func (s service) GetVariantsBySubject(ctx context.Context, subject models.Subject) ([]models.Variant, error) {
	fmt.Println(subject)
	subject, err := s.rep.GetSubject(ctx, subject)
	if err != nil {
		return nil, err
	}
	fmt.Println("slattttt", subject.Id)
	variants, err := s.rep.GetVariantsBySubjectId(ctx, subject.Id)
	if err != nil {
		return nil, err
	}
	fmt.Println("xyyyyu", variants)

	return variants, nil
}

func (s service) GetVariantTypes(ctx context.Context) ([]models.Variant, error) {
	variantTypes, err := s.rep.GetVariantTypes(ctx)
	if err != nil {
		return nil, err
	}
	return variantTypes, nil
}
