package models

type Subject struct {
	Id            int       `json:"id" db:"id"`
	Name          string    `json:"name" db:"name"`
	Semester      int       `json:"semester_number" db:"semester_number"`
	TypeName      string    `json:"type_name" db:"type_name"`
	InstistuteNum int       `json:"instistute_num" db:"instistute_num"`
	Variants      []Variant `json:"-" db:"-"`
}
type SubjectType struct {
	Id   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}
