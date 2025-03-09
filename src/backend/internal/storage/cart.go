package storage

import (
	"backend/internal/dto"
	"backend/internal/model"
	"context"
)

type ICartStorage interface {
	Create(ctx context.Context, cart *model.Cart) error
	Update(ctx context.Context, cart *model.Cart) error
	Remove(ctx context.Context, userID int) error
	AddRacket(ctx context.Context, req *dto.AddRacketCartReq) error
	RemoveRacket(ctx context.Context, req *dto.RemoveRacketCartReq) error
	GetCartByID(ctx context.Context, userID int) (*model.Cart, error)
}
