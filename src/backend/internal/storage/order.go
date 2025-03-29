package storage

import (
	"context"

	"backend/internal/dto"
	"backend/internal/model"
)

type IOrderStorage interface {
	Create(ctx context.Context, order *model.Order) error
	Update(ctx context.Context, order *model.Order) error
	Remove(ctx context.Context, orderID int) error
	GetAllOrders(ctx context.Context, req *dto.ListOrdersReq) ([]*model.Order, error)
	GetMyOrders(ctx context.Context, req *dto.ListOrdersReq) ([]*model.Order, error)
	GetOrderByID(ctx context.Context, orderID int) (*model.Order, error)
}
