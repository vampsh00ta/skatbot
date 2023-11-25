package service

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"skat_bot/internal/repository/models"
	"strconv"
	"strings"
	"time"
)

const (
	token = "6352576956:AAH1icJcTeDiTI7ppqTuSrXZ8QdQ_sZAGYU"
)

type ResultRequest struct {
	FilePath string `json:"file_path"`
}
type FileRequest struct {
	Result ResultRequest `json:"result"`
}
type Variant interface {
	AddVariant(ctx context.Context, variant models.Variant) error
	GetVariantbyId(ctx context.Context, id int) (models.Variant, error)
	GetVariantsBySubject(ctx context.Context, subject models.Subject) ([]models.Variant, error)
	GetVariantsBySubjectId(ctx context.Context, id int) ([]models.Variant, error)

	GetVariantTypes(ctx context.Context) ([]models.Variant, error)
	DownloadVariant(ctx context.Context, variant models.Variant) (string, *[]byte, error)
	DownloadVariantById(ctx context.Context, id int) (string, *[]byte, error)
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
	const urlFormFile = "https://api.telegram.org/file/bot%s/%s"
	const urlFormPath = "https://api.telegram.org/bot%s/getFile?file_id=%s"

	urlPath := fmt.Sprintf(urlFormPath, token, variant.FileId)

	resp, err := http.Get(urlPath)
	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	var pathReq FileRequest
	if err := json.Unmarshal(body, &pathReq); err != nil {
		return "", nil, err
	}
	urlFile := fmt.Sprintf(urlFormFile, token, pathReq.Result.FilePath)
	fileNameTg := strings.Split(pathReq.Result.FilePath, "/")[1]
	fileTypeTg := strings.Split(fileNameTg, ".")[1]
	file := strings.Join(strings.Split(variant.Name, " "), "_") + "." + fileTypeTg
	fileName := strconv.Itoa(variant.Id)
	out, err := os.Create(fileName)
	if err != nil {
		return "", nil, nil
	}
	defer out.Close()
	resp, err = http.Get(urlFile)
	defer resp.Body.Close()
	_, err = io.Copy(out, resp.Body)
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
