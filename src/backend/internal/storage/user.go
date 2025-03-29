package storage

import (
	"context"

	"backend/internal/model"
)

type IUserStorage interface {
	Create(ctx context.Context, user *model.User) error
	UpdateRole(ctx context.Context, user *model.User) error
	GetUserByID(ctx context.Context, id int) (*model.User, error)
	GetUserByEmail(ctx context.Context, email string) (*model.User, error)
	GetAllUsers(ctx context.Context) ([]*model.User, error)
	Remove(ctx context.Context, email string) error
}
