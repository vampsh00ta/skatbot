package service

import (
	rep "skat_bot/internal/repository"
)

type Service interface {
	Subject
	Variant
	Semester
	Institute
	Skat
	Stats
}
type service struct {
	rep rep.Repository
}

func New(r rep.Repository) Service {
	return &service{
		rep: r,
	}
}
