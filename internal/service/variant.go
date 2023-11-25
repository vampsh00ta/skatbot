package service

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"skat_bot/internal/repository/models"
	"strings"
)

const (
	token = "6352576956:AAH1icJcTeDiTI7ppqTuSrXZ8QdQ_sZAGYU"
)

type FileRequest struct {
	FileId   string `json:"fileId"`
	FilePath string `json:"filePath"`
}
type Variant interface {
	AddVariant(ctx context.Context, variant models.Variant) error
	GetVariantsBySubject(ctx context.Context, subject models.Subject) ([]models.Variant, error)
	GetVariantTypes(ctx context.Context) ([]models.Variant, error)
	DownloadVariant(ctx context.Context, variant models.Variant) error
}

func (s service) DownloadVariant(ctx context.Context, variant models.Variant) error {
	const urlForm = "https://api.telegram.org/file/bot%s/%s"
	url := fmt.Sprintf(urlForm, token, variant.FilePath)
	fileNameTg := strings.Split(variant.FilePath, "/")[1]

	fileTypeTg := strings.Split(fileNameTg, ".")[1]
	out, err := os.Create(variant.Name + "." + fileTypeTg)
	if err != nil {
		return err
	}
	defer out.Close()
	resp, err := http.Get(url)
	defer resp.Body.Close()
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}
	return nil
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
	variants, err := s.rep.GetVariantsBySubjectId(ctx, subject.Id)
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
