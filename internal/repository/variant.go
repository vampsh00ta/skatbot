package psql

import (
	"context"
	"github.com/jackc/pgx/v5"
	"skat_bot/internal/repository/models"
)

type Variant interface {
	AddVariant(ctx context.Context, variant models.Variant) (models.Variant, error)
	DeleteVariantById(ctx context.Context, variantId int) error
	GetVariantsBySubjectId(ctx context.Context, subjectId int) ([]models.Variant, error)
	GetVariantTypes(ctx context.Context) ([]models.Variant, error)
}

func (d Db) AddVariant(ctx context.Context, variant models.Variant) (models.Variant, error) {
	var err error
	//
	q := `insert into variant (subject_id,name,num,grade,creation_time,type_name,file_id)
			values ($1,$2,$3,$4,$5,$6,$7) returning id 
		 `
	//loc, _ := time.LoadLocation("Europe/Moscow")
	//t := time.Now().In(loc)
	//variantType, err := d.GetVariantTypeByName(ctx, variant.TypeString)
	//if err != nil {
	//	return nil
	//}
	if err = d.client.QueryRow(ctx, q,
		variant.SubjectId,
		variant.Name,
		variant.Num,
		variant.Grade,
		variant.CreationTime,
		variant.TypeName,
		variant.FileId).Scan(&variant.Id); err != nil {

		return models.Variant{}, err
	}
	return variant, nil
}

func (d Db) DeleteVariantById(ctx context.Context, variantId int) error {
	//TODO implement me
	panic("implement me")
}

func (d Db) GetVariantsBySubjectId(ctx context.Context, subjectId int) ([]models.Variant, error) {
	var err error
	//
	q := `select * from  variant where subject_id = $1
		 `

	rows, err := d.client.Query(ctx, q, subjectId)
	if err != nil {
		return nil, err
	}
	variants, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[models.Variant])
	if err != nil {
		return nil, err
	}

	return variants, nil
}
func (d Db) GetVariantTypes(ctx context.Context) ([]models.Variant, error) {
	var err error
	//
	q := `select name as type_name from variant_type 
		 `
	rows, err := d.client.Query(ctx, q)
	if err != nil {
		return nil, err
	}
	variantTypes, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[models.Variant])
	if err != nil {
		return nil, err
	}

	return variantTypes, nil
}
