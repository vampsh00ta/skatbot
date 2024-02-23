package psql

import (
	"context"
	"skat_bot/pkg/client"
)

type Repository interface {
	Subject
	Semester
	Variant
	Institute
	Stats
	WithTransaction(ctx context.Context, f func(ctx context.Context) error) error
}

type Db struct {
	client client.Client
}

func New(client client.Client) Repository {

	return &Db{
		client: client,
	}
}
