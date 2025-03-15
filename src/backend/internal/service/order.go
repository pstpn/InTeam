package service

import (
	"backend/internal/dto"
	"backend/internal/model"
	"backend/internal/storage"
	"backend/pkg/common"
	"backend/pkg/logger"
	"context"
	"fmt"
	"time"
)

type IOrderService interface {
	CreateOrder(ctx context.Context, req *dto.PlaceOrderReq) error
	GetMyOrders(ctx context.Context, req *dto.ListOrdersReq) ([]*model.Order, error)
	GetAllOrders(ctx context.Context, req *dto.ListOrdersReq) ([]*model.Order, error)
	GetOrderByID(ctx context.Context, orderID int) (*model.Order, error)
	UpdateOrderStatus(ctx context.Context, req *dto.UpdateOrderReq) (*model.Order, error)
}

type OrderService struct {
	logger     logger.Interface
	repo       storage.IOrderStorage
	repoCart   storage.ICartStorage
	repoRacket storage.IRacketStorage
}

func NewOrderService(logger logger.Interface, repo storage.IOrderStorage, repoCart storage.ICartStorage, repoRacket storage.IRacketStorage) *OrderService {
	return &OrderService{
		logger:     logger,
		repo:       repo,
		repoCart:   repoCart,
		repoRacket: repoRacket,
	}
}

func (s *OrderService) CreateOrder(ctx context.Context, req *dto.PlaceOrderReq) error {
	order := &model.Order{}
	common.Copy(&order, &req)

	order.CreationDate = time.Now()

	cart, err := s.repoCart.GetCartByID(ctx, order.UserID)
	if err != nil {
		s.logger.Errorf("get cart fail, error %s", err.Error())
		return fmt.Errorf("get cart fail, error %s", err)
	}

	for _, line := range cart.Lines {
		racket, err := s.repoRacket.GetRacketByID(ctx, line.RacketID)
		if err != nil {
			s.logger.Errorf("get racket by id fail, error %s", err.Error())
			return fmt.Errorf("get racket by id fail, error %s", err)
		}

		if line.Quantity <= racket.Quantity {
			racket.Quantity -= line.Quantity
		} else {
			s.logger.Errorf("not available amount or rackets, error")
			return fmt.Errorf("not available amount or rackets, error")
		}

		err = s.repoRacket.Update(ctx, racket)
		if err != nil {
			s.logger.Errorf("update racket after order creation fail, error %s", err.Error())
			return fmt.Errorf("update racket after order creation fail, error %s", err)
		}
	}

	order.Status = model.OrderStatusInProgress
	order.TotalPrice = cart.TotalPrice

	for _, line := range cart.Lines {
		order.Lines = append(order.Lines,
			&model.OrderLine{
				RacketID: line.RacketID,
				Quantity: line.Quantity,
			},
		)
	}

	err = s.repo.Create(ctx, order)
	if err != nil {
		s.logger.Errorf("create order fail, error %s", err.Error())
		return fmt.Errorf("create order fail, error %s", err)
	}

	err = s.repoCart.Remove(ctx, req.UserID)
	if err != nil {
		s.logger.Errorf("remove racket fail, error %s", err.Error())
		return fmt.Errorf("remove racket fail, error %s", err)
	}

	return nil
}

func (s *OrderService) GetMyOrders(ctx context.Context, req *dto.ListOrdersReq) ([]*model.Order, error) {
	orders, err := s.repo.GetMyOrders(ctx, req)
	if err != nil {
		s.logger.Errorf("get my orders fail, error %s", err.Error())
		return nil, fmt.Errorf("get my orders fail, error %s", err)
	}

	return orders, nil
}

func (s *OrderService) GetAllOrders(ctx context.Context, req *dto.ListOrdersReq) ([]*model.Order, error) {
	orders, err := s.repo.GetAllOrders(ctx, req)
	if err != nil {
		s.logger.Errorf("get all fail, error %s", err.Error())
		return nil, fmt.Errorf("get all fail, error %s", err)
	}

	return orders, nil
}

func (s *OrderService) GetOrderByID(ctx context.Context, orderID int) (*model.Order, error) {
	order, err := s.repo.GetOrderByID(ctx, orderID)
	if err != nil {
		s.logger.Errorf("get order by id fail, error %s", err.Error())
		return nil, fmt.Errorf("get order by id fail, error %s", err)
	}

	return order, nil
}

func (s *OrderService) UpdateOrderStatus(ctx context.Context, req *dto.UpdateOrderReq) (*model.Order, error) {
	order, err := s.repo.GetOrderByID(ctx, req.OrderID)
	if err != nil {
		s.logger.Errorf("get by id fail, error %s", err.Error())
		return nil, fmt.Errorf("get by id fail, error %s", err)
	}

	order.Status = req.Status

	err = s.repo.Update(ctx, order)
	if err != nil {
		s.logger.Errorf("update order fail, error %s", err.Error())
		return nil, fmt.Errorf("update order fail, error %s", err)
	}

	return order, nil
}
