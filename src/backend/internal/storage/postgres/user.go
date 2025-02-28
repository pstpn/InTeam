package mypostgres

import (
	"backend/internal/model"
	"backend/internal/storage"
	"backend/pkg/storage/postgres"
	"context"
	"github.com/go-faster/errors"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
)

type UserStorage struct {
	*postgres.Postgres
}

func NewUserStorage(db *postgres.Postgres) storage.IUserRepository {
	return &UserStorage{db}
}

func (r *UserStorage) Create(ctx context.Context, user *model.User) error {
	query := r.Builder.
		Insert(userTable).
		Columns(
			nameField,
			surnameField,
			emailField,
			passwordField,
			roleField).
		Values(user.Name,
			user.Surname,
			user.Email,
			user.Password,
			user.Role).
		Suffix("returning id")

	sql, ars, err := query.ToSql()
	if err != nil {
		return err
	}

	row := r.Pool.QueryRow(ctx, sql, ars...)

	return row.Scan(
		&user.ID,
	)
}

func (r *UserStorage) UpdateRole(ctx context.Context, user *model.User) error {
	query := r.Builder.
		Update(userTable).
		Set(roleField, user.Role).
		Where(squirrel.Eq{emailField: user.Email})

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

func (r *UserStorage) GetAllUsers(ctx context.Context) ([]*model.User, error) {
	query := r.Builder.
		Select("*").
		From(userTable)

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := r.Pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}

	var users []*model.User
	for rows.Next() {
		user, err := r.rowToModel(rows)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (r *UserStorage) GetUserByID(ctx context.Context, id int) (*model.User, error) {
	query := r.Builder.
		Select("*").
		From(userTable).
		Where(squirrel.Eq{idField: id})

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	row := r.Pool.QueryRow(ctx, sql, args...)

	return r.rowToModel(row)
}

func (r *UserStorage) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	query := r.Builder.
		Select("*").
		From(userTable).
		Where(squirrel.Eq{emailField: email})

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	row := r.Pool.QueryRow(ctx, sql, args...)

	return r.rowToModel(row)
}

func (r *UserStorage) Remove(ctx context.Context, email string) error {
	query := r.Builder.
		Delete(userTable).
		Where(squirrel.Eq{emailField: email})

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

func (r *UserStorage) rowToModel(row pgx.Row) (*model.User, error) {
	var user model.User

	err := row.Scan(
		&user.ID,
		&user.Name,
		&user.Surname,
		&user.Email,
		&user.Password,
		&user.Role,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, storage.ErrNotFound
		}
		return nil, err
	}

	return &user, nil
}
