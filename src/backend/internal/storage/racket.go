package storage

import (
	"context"

	"backend/internal/dto"
	"backend/internal/model"
)

type IRacketStorage interface {
	Create(ctx context.Context, racket *model.Racket) error
	Update(ctx context.Context, racket *model.Racket) error
	Remove(ctx context.Context, id int) error
	GetRacketByID(ctx context.Context, id int) (*model.Racket, error)
	GetAllRackets(ctx context.Context, req *dto.ListRacketsReq) ([]*model.Racket, error)
}
