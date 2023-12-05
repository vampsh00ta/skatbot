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
	GetVariantbyId(ctx context.Context, id int) (models.Variant, error)
	GetVariantbyTgid(ctx context.Context, id string) ([]models.Variant, error)
}

func (d Db) GetVariantbyTgid(ctx context.Context, id string) ([]models.Variant, error) {
	var err error
	q := `select * from variant  where tg_id = $1
		 `
	rows, err := d.query(ctx, q, id)

	variant, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[models.Variant])
	if err != nil {
		return nil, err
	}

	return variant, nil
}
func (d Db) GetVariantbyId(ctx context.Context, id int) (models.Variant, error) {
	var err error
	//
	q := `select * from variant  where id = $1
		 `
	rows, err := d.query(ctx, q, id)
	if err != nil {
		return models.Variant{}, err
	}
	variant, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[models.Variant])
	if err != nil {
		return models.Variant{}, err
	}
	var v models.Variant
	if len(variant) > 0 {
		v = variant[0]
	} else {
		v = models.Variant{}
	}
	return v, nil
}
func (d Db) AddVariant(ctx context.Context, variant models.Variant) (models.Variant, error) {
	var err error
	//

	q := `insert into variant (subject_id,name,num,grade,creation_time,type_name,file_id,file_path,tg_id,tg_username,file_type)
			values ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11) returning id 
		 `
	//loc, _ := time.LoadLocation("Europe/Moscow")
	//t := time.Now().In(loc)
	//variantType, err := d.GetVariantTypeByName(ctx, variant.TypeString)
	//if err != nil {
	//	return nil
	//}
	if err = d.queryRow(ctx, q,
		variant.SubjectId,
		variant.Name,
		variant.Num,
		variant.Grade,
		variant.CreationTime,
		variant.TypeName,
		variant.FileId,
		variant.FilePath,
		variant.TgId,
		variant.TgUsername,
		variant.FileType).Scan(&variant.Id); err != nil {

		return models.Variant{}, err
	}
	return variant, nil
}

func (d Db) DeleteVariantById(ctx context.Context, variantId int) error {
	var err error
	//

	q := `delete  from  variant where id = $1 returning id
			
		 `
	if err = d.queryRow(ctx, q, variantId).Scan(&variantId); err != nil {

		return err
	}
	return nil
}

func (d Db) GetVariantsBySubjectId(ctx context.Context, subjectId int) ([]models.Variant, error) {
	var err error
	//
	q := `select * from  variant where subject_id = $1
		 `

	rows, err := d.query(ctx, q, subjectId)
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
	rows, err := d.query(ctx, q)
	if err != nil {
		return nil, err
	}
	variantTypes, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[models.Variant])
	if err != nil {
		return nil, err
	}

	return variantTypes, nil
}
