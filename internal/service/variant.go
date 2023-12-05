package service

import (
	"context"
	"os"
	"skat_bot/internal/repository/models"
	"time"
)

const (
	token = "6352576956:AAH1icJcTeDiTI7ppqTuSrXZ8QdQ_sZAGYU"
)

type Variant interface {
	AddVariant(ctx context.Context, variant models.Variant) error
	DeleteVariantById(ctx context.Context, id int) error
	GetVariantbyId(ctx context.Context, id int) (models.Variant, error)
	GetVariantsBySubject(ctx context.Context, subject models.Subject) ([]models.Variant, error)
	GetVariantsBySubjectId(ctx context.Context, id int) ([]models.Variant, error)
	GetVariantbyTgid(ctx context.Context, id string) ([]models.Variant, error)

	GetVariantTypes(ctx context.Context) ([]models.Variant, error)
	DownloadVariant(ctx context.Context, variant models.Variant) (string, *[]byte, error)
	DownloadVariantById(ctx context.Context, id int) (string, *[]byte, error)
}

func (s service) DeleteVariantById(ctx context.Context, id int) error {
	err := s.rep.WithTransaction(ctx, func(ctx context.Context) error {
		variant, err := s.rep.GetVariantbyId(ctx, id)
		if err != nil {
			return err
		}
		subjectVariats, err := s.rep.GetVariantsBySubjectId(ctx, variant.SubjectId)
		if err != nil {
			return err
		}
		if len(subjectVariats) == 1 {
			err = s.rep.DeleteSubjectById(ctx, variant.SubjectId)

		} else {
			err = s.rep.DeleteVariantById(ctx, id)
		}

		return err
	})
	return err
}

func (s service) GetVariantbyTgid(ctx context.Context, id string) ([]models.Variant, error) {
	variant, err := s.rep.GetVariantbyTgid(ctx, id)
	if err != nil {
		return nil, err
	}
	return variant, nil
}

func (s service) GetVariantbyId(ctx context.Context, id int) (models.Variant, error) {
	variant, err := s.rep.GetVariantbyId(ctx, id)
	if err != nil {
		return models.Variant{}, err
	}
	return variant, nil
}
func (s service) DownloadVariantById(ctx context.Context, id int) (string, *[]byte, error) {
	variant, err := s.GetVariantbyId(ctx, id)
	if err != nil {
		return "", nil, nil
	}
	fileName, b, err := s.DownloadVariant(ctx, variant)
	if err != nil {
		return "", nil, nil
	}
	return fileName, b, nil

}
func (s service) DownloadVariant(ctx context.Context, variant models.Variant) (string, *[]byte, error) {
	filePath, err := s.getFilePath(variant.FileId)

	if err != nil {
		return "", nil, nil
	}
	file, fileName, err := s.getFile(variant, filePath)
	if err != nil {
		return "", nil, nil
	}

	fileData, errReadFile := os.ReadFile(fileName)
	if errReadFile != nil {
		return "", nil, nil
	}
	go func() {
		time.Sleep(time.Second * 60)
		os.Remove(fileName)
	}()
	return file, &fileData, nil
}
func (s service) AddVariant(ctx context.Context, variant models.Variant) error {
	if _, err := s.rep.AddVariant(ctx, variant); err != nil {
		return err
	}
	return nil
}
func (s service) GetVariantsBySubjectId(ctx context.Context, id int) ([]models.Variant, error) {
	variants, err := s.rep.GetVariantsBySubjectId(ctx, id)
	if err != nil {
		return nil, err
	}

	return variants, nil
}
func (s service) GetVariantsBySubject(ctx context.Context, subject models.Subject) ([]models.Variant, error) {
	var variants []models.Variant
	err := s.rep.WithTransaction(ctx, func(ctx context.Context) error {
		subject, err := s.rep.GetSubject(ctx, subject)
		if err != nil {
			return err
		}
		variants, err = s.rep.GetVariantsBySubjectId(ctx, subject.Id)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return variants, nil
}

func (s service) GetVariantTypes(ctx context.Context) ([]models.Variant, error) {
	variantTypes, err := s.rep.GetVariantTypes(ctx)
	if err != nil {
		return nil, err
	}
	return variantTypes, nil
}
