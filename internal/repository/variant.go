package psql

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"skat_bot/internal/repository/models"
)

type Variant interface {
	AddVariant(ctx context.Context, variant models.Variant) (int, error)
	DeleteVariantById(ctx context.Context, variantId int) error
	GetVariantsBySubjectId(ctx context.Context, subjectId int) ([]models.Variant, error)
	GetVariantTypes(ctx context.Context) ([]models.VariantType, error)
}

func (d Db) GetVariantTypeByName(ctx context.Context, variantT string) (models.VariantType, error) {
	q := `select * from variant_type where name = $1
		 `
	var variantType models.VariantType
	if err := d.client.QueryRow(ctx, q, variantT).Scan(&variantType.Id, &variantType.Name); err != nil {
		return models.VariantType{}, err
	}
	return variantType, nil
}

func (d Db) AddVariant(ctx context.Context, variant models.Variant) (int, error) {
	var err error
	//
	fmt.Println(variant)
	q := `insert into variant (subject_id,name,num,grade,creation_time,type_name)
			values ($1,$2,$3,$4,$5,$6) returning id 
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
		variant.TypeName).Scan(&variant.Id); err != nil {

		return 0, err
	}
	return variant.Id, nil
}

func (d Db) DeleteVariantById(ctx context.Context, variantId int) error {
	//TODO implement me
	panic("implement me")
}

func (d Db) GetVariantsBySubjectId(ctx context.Context, subjectId int) ([]models.Variant, error) {
	var err error
	//
	fmt.Println(subjectId)
	q := `select * from  variant where subject_id = $1
		 `

	rows, err := d.client.Query(ctx, q, subjectId)
	if err != nil {
		return nil, err
	}
	variants, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Variant])
	if err != nil {
		return nil, err
	}

	return variants, nil
}
func (d Db) GetVariantTypes(ctx context.Context) ([]models.VariantType, error) {
	var err error
	//
	q := `select * from variant_type 
		 `
	rows, err := d.client.Query(ctx, q)
	if err != nil {
		return nil, err
	}
	variantTypes, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.VariantType])
	if err != nil {
		return nil, err
	}

	return variantTypes, nil
}
