package mypostgres

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"

	"backend/internal/dto"
	"backend/internal/model"
	"backend/internal/storage"
	"backend/pkg/storage/postgres"
)

type CartStorage struct {
	*postgres.Postgres
}

func NewCartStorage(db *postgres.Postgres) storage.ICartStorage {
	return &CartStorage{db}
}

func (r *CartStorage) AddRacket(ctx context.Context, req *dto.AddRacketCartReq) error {
	query := r.Builder.
		Insert(cartRacketTable).
		Columns(
			cartIDField,
			racketIDField,
			quantityField).
		Values(
			req.UserID,
			req.RacketID,
			req.Quantity).
		Suffix("returning cart_id")

	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}

	row := r.Pool.QueryRow(ctx, sql, args...)
	err = row.Scan(
		&req.UserID,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *CartStorage) RemoveRacket(ctx context.Context, req *dto.RemoveRacketCartReq) error {
	query := r.Builder.
		Delete(cartRacketTable).
		Where(squirrel.And{
			squirrel.Eq{cartIDField: req.UserID},
			squirrel.Eq{racketIDField: req.RacketID},
		})

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

func (r *CartStorage) Create(ctx context.Context, cart *model.Cart) error {
	query := r.Builder.
		Insert(cartTable).
		Columns(
			userIDField,
			totalPriceField,
			quantityField).
		Values(
			cart.UserID,
			cart.TotalPrice,
			cart.Quantity).
		Suffix("returning user_id")

	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}

	row := r.Pool.QueryRow(ctx, sql, args...)
	err = row.Scan(
		&cart.UserID,
	)
	if err != nil {
		return err
	}

	for _, line := range cart.Lines {
		query := r.Builder.
			Insert(cartRacketTable).
			Columns(
				racketIDField,
				cartIDField,
				quantityField).
			Values(
				line.RacketID,
				cart.UserID,
				line.Quantity).
			Suffix("returning racket_id")

		sql, args, err := query.ToSql()
		if err != nil {
			return err
		}

		row := r.Pool.QueryRow(ctx, sql, args...)
		err = row.Scan(
			&line.RacketID,
		)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *CartStorage) Update(ctx context.Context, cart *model.Cart) error {
	for _, line := range cart.Lines {
		query := r.Builder.
			Update(cartRacketTable).
			Set(quantityField, line.Quantity).
			Where(squirrel.And{
				squirrel.Eq{cartIDField: cart.UserID},
				squirrel.Eq{racketIDField: line.RacketID},
			})

		sql, args, err := query.ToSql()
		if err != nil {
			return err
		}

		_, err = r.Pool.Exec(ctx, sql, args...)
		if err != nil {
			return err
		}
	}

	query := r.Builder.
		Update(cartTable).
		Set(quantityField, cart.Quantity).
		Set(totalPriceField, cart.TotalPrice).
		Where(squirrel.Eq{userIDField: cart.UserID})

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

func (r *CartStorage) GetCartByID(ctx context.Context, userID int) (*model.Cart, error) {
	query := r.Builder.
		Select("*").
		From(cartTable).
		Where(squirrel.Eq{userIDField: userID})

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	row := r.Pool.QueryRow(ctx, sql, args...)
	cart, err := r.rowToModel(row)
	if err != nil {
		return nil, err
	}

	cart.Quantity = 0
	cart.TotalPrice = 0

	query = r.Builder.
		Select(racketIDField, fmt.Sprintf("%s.%s", cartRacketTable, quantityField), priceField).
		From(cartRacketTable).
		Join(on(cartRacketTable, racketTable, racketIDField, idField)).
		Where(squirrel.Eq{cartIDField: userID})

	sql, args, err = query.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := r.Pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		line, err := r.rowToModelCartRacket(rows)
		if err != nil {
			return nil, err
		}

		cart.Quantity += line.Quantity
		cart.TotalPrice += float32(line.Quantity) * line.Price
		cart.Lines = append(cart.Lines, line)
	}

	return cart, nil
}

func (r *CartStorage) Remove(ctx context.Context, userID int) error {
	query := r.Builder.
		Delete(cartTable).
		Where(squirrel.Eq{userIDField: userID})

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

func (r *CartStorage) rowToModel(row pgx.Row) (*model.Cart, error) {
	var cart model.Cart

	err := row.Scan(
		&cart.UserID,
		&cart.Quantity,
		&cart.TotalPrice,
	)
	if err != nil {
		return nil, err
	}

	return &cart, nil
}

func (r *CartStorage) rowToModelCartRacket(row pgx.Row) (*model.CartLine, error) {
	var cartLine model.CartLine

	err := row.Scan(
		&cartLine.RacketID,
		&cartLine.Quantity,
		&cartLine.Price,
	)
	if err != nil {
		return nil, err
	}

	return &cartLine, nil
}
