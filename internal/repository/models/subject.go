package models

type Subject struct {
	Id       int        `json:"id" db:"id"`
	Name     string     `json:"name" db:"name"`
	Semester int        `json:"semester" db:"semester"`
	Type     int        `json:"type" db:"type"`
	Variants *[]Variant `json:"-" db:"-"`
}
type SubjectType struct {
	Id   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}
