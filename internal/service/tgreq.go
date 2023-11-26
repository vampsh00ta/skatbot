package service

import (
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
)

type TgReq interface {
	getFilePath(fileId string) (string, error)
	getFile(variant models.Variant, filePath string) (string, string, error)
	//getUsernameById(id int) (string, error)
}
type ResultRequest struct {
	FilePath string `json:"file_path"`
}
type FileRequest struct {
	Result ResultRequest `json:"result"`
}

func (s service) getFilePath(fileId string) (string, error) {
	const urlFormPath = "https://api.telegram.org/bot%s/getFile?file_id=%s"

	urlPath := fmt.Sprintf(urlFormPath, token, fileId)
	fmt.Println(urlPath)
	resp, err := http.Get(urlPath)
	if err != nil {
		return "", err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	var pathReq FileRequest
	if err := json.Unmarshal(body, &pathReq); err != nil {
		return "", err
	}

	return pathReq.Result.FilePath, nil
}

func (s service) getFile(variant models.Variant, filePath string) (string, string, error) {
	const urlFormFile = "https://api.telegram.org/file/bot%s/%s"
	urlFile := fmt.Sprintf(urlFormFile, token, filePath)
	fmt.Println(urlFile)

	fileNameTg := strings.Split(filePath, "/")[1]
	fileTypeTg := strings.Split(fileNameTg, ".")[1]
	file := strings.Join(strings.Split(variant.Name, " "), "_") + "." + fileTypeTg
	fileName := strconv.Itoa(variant.Id)
	out, err := os.Create(fileName)
	if err != nil {
		return "", "", nil
	}
	defer out.Close()
	resp, err := http.Get(urlFile)
	defer resp.Body.Close()
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return "", "", nil
	}
	return file, fileName, nil
}

//func (s service) getUsernameById(id int) (string, error) {
//	const urlFormPath = "https://api.telegram.org/bot6352576956:AAH1icJcTeDiTI7ppqTuSrXZ8QdQ_sZAGYU/getUser?user_id=564764193"
//
//	urlPath := fmt.Sprintf(urlFormPath, token, fileId)
//
//	resp, err := http.Get(urlPath)
//	if err != nil {
//		return "", err
//	}
//	body, err := ioutil.ReadAll(resp.Body)
//	if err != nil {
//		log.Fatalln(err)
//	}
//	var pathReq FileRequest
//	if err := json.Unmarshal(body, &pathReq); err != nil {
//		return "", err
//	}
//
//	return pathReq.Result.FilePath, nil
//}
