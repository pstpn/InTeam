package mypostgres

import (
	"backend/internal/dto"
	"backend/internal/model"
	"backend/internal/storage"
	"backend/pkg/storage/postgres"
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
)

type RacketStorage struct {
	*postgres.Postgres
}

func NewRacketStorage(db *postgres.Postgres) storage.IRacketStorage {
	return &RacketStorage{db}
}

func (r *RacketStorage) Create(ctx context.Context, racket *model.Racket) error {
	query := r.Builder.
		Insert(racketTable).
		Columns(
			brandField,
			weightField,
			balanceField,
			headSizeField,
			quantityField,
			priceField,
			availableField,
			imageField).
		Values(
			racket.Brand,
			racket.Weight,
			racket.Balance,
			racket.HeadSize,
			racket.Quantity,
			racket.Price,
			racket.Available,
			racket.Image).
		Suffix("returning id")

	sql, ars, err := query.ToSql()
	if err != nil {
		return err
	}

	row := r.Pool.QueryRow(ctx, sql, ars...)

	err = row.Scan(
		&racket.ID,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *RacketStorage) Update(ctx context.Context, racket *model.Racket) error {
	query := r.Builder.
		Update(racketTable).
		Set(quantityField, racket.Quantity).
		Set(availableField, racket.Available).
		Where(squirrel.Eq{idField: racket.ID})

	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}

	_, err = r.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *RacketStorage) Remove(ctx context.Context, id int) error {
	query := r.Builder.
		Delete(racketTable).
		Where(squirrel.Eq{idField: id})

	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}

	_, err = r.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *RacketStorage) GetRacketByID(ctx context.Context, id int) (*model.Racket, error) {
	query := r.Builder.
		Select("*").
		From(racketTable).
		Where(squirrel.Eq{idField: id})

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	row := r.Pool.QueryRow(ctx, sql, args...)

	return r.rowToModel(row)
}

func (r *RacketStorage) GetAllRackets(ctx context.Context, req *dto.ListRacketsReq) ([]*model.Racket, error) {
	query := r.Builder.
		Select("*").
		From(racketTable)

	query = req.Pagination.ToSQL(query)

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := r.Pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}

	var rackets []*model.Racket

	for rows.Next() {
		racket, err := r.rowToModel(rows)
		if err != nil {
			return nil, err
		}

		rackets = append(rackets, racket)
	}

	return rackets, nil
}

func (r *RacketStorage) rowToModel(row pgx.Row) (*model.Racket, error) {
	var racket model.Racket

	err := row.Scan(
		&racket.ID,
		&racket.Brand,
		&racket.Weight,
		&racket.Balance,
		&racket.HeadSize,
		&racket.Available,
		&racket.Price,
		&racket.Quantity,
		&racket.Image,
	)
	if err != nil {
		return nil, err
	}

	return &racket, nil
}
