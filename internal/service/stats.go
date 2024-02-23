package service

import (
	"context"
)

type Stats interface {
	AddDeleteLike(ctx context.Context, variantId int, username string) (int, error)
}

func (s service) AddDeleteLike(ctx context.Context, variantId int, username string) (int, error) {
	var res int
	err := s.rep.WithTransaction(ctx, func(ctx context.Context) error {
		var err error
		like, err := s.rep.CheckLike(ctx, variantId, username)
		if err != nil {
			return err
		}
		if len(like) == 0 {
			err = s.rep.AddLike(ctx, variantId, username)
			res = 1
		} else {
			err = s.rep.DeleteLike(ctx, variantId, username)
			res = -1
		}

		return err
	})
	if err != nil {
		return 0, err
	}

	return res, err
}
