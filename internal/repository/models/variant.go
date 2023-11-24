package models

import "time"

type Variant struct {
	Id           int       `json:"id" db:"id"`
	Name         string    `json:"name" db:"name"`
	SubjectId    int       `json:"subject_id" db:"subject_id"`
	Num          int       `json:"num" db:"num"`
	Grade        *int      `json:"grade" db:"grade"`
	CreationTime time.Time `json:"creation_time" db:"creation_time"`
	TypeName     string    `json:"type_name" db:"type_name"`
}

type VariantType struct {
	Id   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}
