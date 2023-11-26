package models

import "time"

type Variant struct {
	Id           int       `json:"id" db:"id"`
	Name         string    `json:"name" db:"name"`
	SubjectId    int       `json:"subject_id" db:"subject_id"`
	Num          *int      `json:"num" db:"num"`
	Grade        *int      `json:"grade" db:"grade"`
	CreationTime time.Time `json:"creation_time" db:"creation_time"`
	TypeName     string    `json:"type_name" db:"type_name"`
	FileId       string    `json:"file_id" db:"file_id"`
	FilePath     string    `json:"file_path" db:"file_path"`
	TgId         string    `json:"tg_id" db:"tg_id"`
	TgUsername   string    `json:"tg_username" db:"tg_username"`
}
