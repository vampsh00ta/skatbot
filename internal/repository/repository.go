package psql

import (
	"skat_bot/pkg/client"
)

type Repository interface {
	Subject
	Semester
	Variant
	Institute
}

type Db struct {
	client client.Client
}

func New(client client.Client) Repository {

	return &Db{
		client: client,
	}
}
