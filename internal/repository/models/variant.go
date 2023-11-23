package models

import "time"

type Variant struct {
	Id   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`

	SubjectId    int       `json:"subject_id" db:"subject_id"`
	Num          int       `json:"num_from" db:"num_from"`
	Grade        *int      `json:"grade" db:"grade"`
	CreationTime time.Time `json:"creation_time" db:"creation_time"`
	Type         int       `json:"type" db:"type"`
	TypeString   *string   `json:"type_string" db:"type_string"`
}

type VariantType struct {
	Id   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}
