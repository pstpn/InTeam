package service

import (
	"context"
	"errors"
	"fmt"

	"backend/internal/dto"
	"backend/internal/model"
	"backend/internal/storage"
	"backend/pkg/logger"
)

type ICartService interface {
	AddRacket(ctx context.Context, req *dto.AddRacketCartReq) (*model.Cart, error)
	RemoveRacket(ctx context.Context, req *dto.RemoveRacketCartReq) (*model.Cart, error)
	UpdateRacket(ctx context.Context, req *dto.UpdateRacketCartReq) (*model.Cart, error)
	GetCartByID(ctx context.Context, userID int) (*model.Cart, error)
}

type CartService struct {
	l          logger.Interface
	repo       storage.ICartStorage
	repoRacket storage.IRacketStorage
}

func NewCartService(l logger.Interface, repo storage.ICartStorage, repoRacket storage.IRacketStorage) *CartService {
	return &CartService{
		l:          l,
		repo:       repo,
		repoRacket: repoRacket,
	}
}

func (s *CartService) UpdateRacket(ctx context.Context, req *dto.UpdateRacketCartReq) (*model.Cart, error) {
	cart, _ := s.repo.GetCartByID(ctx, req.UserID)

	for i, lines := range cart.Lines {
		if lines.RacketID != req.RacketID {
			continue
		}

		racket, _ := s.repoRacket.GetRacketByID(ctx, req.RacketID)

		curQuantity := lines.Quantity + req.Quantity
		switch {
		case curQuantity <= -1:
			return cart, nil
		case curQuantity >= racket.Quantity:
			curQuantity = racket.Quantity
		case curQuantity == 0:
			cart.Lines = append(cart.Lines[:i], cart.Lines[i+1:]...)
		}

		lines.Quantity = curQuantity
		cart.Quantity += req.Quantity
		cart.TotalPrice += float32(req.Quantity) * lines.Price

		err := s.repo.Update(ctx, cart)
		if err != nil {
			s.l.Errorf("UpdateRacket fail, error %s", err.Error())
			return nil, fmt.Errorf("UpdateRacket fail, error %w", err)
		}

		return cart, nil
	}

	return cart, nil
}

func (s *CartService) AddRacket(ctx context.Context, req *dto.AddRacketCartReq) (*model.Cart, error) {
	cart, err := s.repo.GetCartByID(ctx, req.UserID)
	if err != nil {
		racket, err := s.repoRacket.GetRacketByID(ctx, req.RacketID)
		if err != nil {
			s.l.Errorf("getRacketByID fail, error %s", err.Error())
			return nil, fmt.Errorf("getRacketByID fail, error %w", err)
		}

		if !racket.Available {
			s.l.Errorf("get racket fail, error racket is not available")
			return nil, errors.New("get racket fail, error racket is not available")
		}

		switch {
		case req.Quantity <= 0 || racket.Quantity <= 0:
			s.l.Errorf("addRacketByID fail, error request quantity <= 0")
			return nil, fmt.Errorf("getRacketByID fail, error %w", err)
		case req.Quantity >= racket.Quantity:
			req.Quantity = racket.Quantity
			racket.Quantity = 0
		default:
			racket.Quantity -= req.Quantity
		}

		cart = &model.Cart{
			UserID:     req.UserID,
			TotalPrice: racket.Price * float32(req.Quantity),
			Quantity:   req.Quantity,
			Lines: []*model.CartLine{{
				RacketID: req.RacketID,
				Quantity: req.Quantity,
				Price:    racket.Price,
			}},
		}

		err = s.repo.Create(ctx, cart)
		if err != nil {
			s.l.Errorf("create cart fail, error %s", err.Error())
			return nil, fmt.Errorf("create cart fail, error %w", err)
		}

		return cart, nil
	}

	for _, line := range cart.Lines {
		if line.RacketID == req.RacketID {
			return cart, nil
		}
	}

	racket, err := s.repoRacket.GetRacketByID(ctx, req.RacketID)
	if err != nil {
		s.l.Errorf("get racket fail, error %s", err.Error())
		return nil, fmt.Errorf("get racket fail, error %w", err)
	}

	if !racket.Available {
		s.l.Errorf("get racket fail, error racket is not available")
		return nil, errors.New("get racket fail, error racket is not available")
	}

	switch {
	case req.Quantity <= 0:
		s.l.Errorf("addRacketByID fail, error request quantity <= 0")
		return nil, fmt.Errorf("getRacketByID fail, error %w", err)
	case req.Quantity >= racket.Quantity:
		req.Quantity = racket.Quantity
		racket.Quantity = 0
	default:
		racket.Quantity -= req.Quantity
	}

	err = s.repo.AddRacket(ctx, req)

	if err != nil {
		s.l.Errorf("add racket fail, error %s", err.Error())
		return nil, fmt.Errorf("add racket fail, error %w", err)
	}

	cart.Lines = append(cart.Lines,
		&model.CartLine{
			RacketID: req.RacketID,
			Quantity: req.Quantity,
			Price:    racket.Price,
		})

	cart.Quantity += req.Quantity
	cart.TotalPrice += racket.Price * float32(req.Quantity)

	err = s.repo.Update(ctx, cart)
	if err != nil {
		s.l.Errorf("update cart fail, error %s", err.Error())
		return nil, fmt.Errorf("update cart fail, error %w", err)
	}

	return cart, nil
}

func (s *CartService) RemoveRacket(ctx context.Context, req *dto.RemoveRacketCartReq) (*model.Cart, error) {
	cart, err := s.repo.GetCartByID(ctx, req.UserID)
	if err != nil {
		cart = &model.Cart{UserID: req.UserID}

		err = s.repo.Create(ctx, cart)
		if err != nil {
			s.l.Errorf("create cart fail, error %s", err.Error())
			return nil, fmt.Errorf("create cart fail, error %w", err)
		}

		return cart, nil
	}

	for i := 0; i < len(cart.Lines); i++ {
		if cart.Lines[i].RacketID == req.RacketID {
			continue
		}

		racket, err := s.repoRacket.GetRacketByID(ctx, req.RacketID)
		if err != nil {
			s.l.Errorf("add racket fail, error %s", err.Error())
			return nil, fmt.Errorf("add racket fail, error %w", err)
		}

		cart.Quantity -= cart.Lines[i].Quantity
		cart.TotalPrice -= racket.Price * float32(cart.Lines[i].Quantity)

		cart.Lines = append(cart.Lines[:i], cart.Lines[i+1:]...)
		break
	}

	err = s.repo.RemoveRacket(ctx, req)
	if err != nil {
		s.l.Errorf("remove racket fail, error %s", err.Error())
		return nil, fmt.Errorf("remove racket fail, error %w", err)
	}

	err = s.repo.Update(ctx, cart)
	if err != nil {
		s.l.Errorf("update racket fail, error %s", err.Error())
		return nil, fmt.Errorf("update racket fail, error %w", err)
	}

	return cart, nil
}

func (s *CartService) GetCartByID(ctx context.Context, userID int) (*model.Cart, error) {
	cart, err := s.repo.GetCartByID(ctx, userID)
	if err != nil {
		cart := &model.Cart{UserID: userID}

		err = s.repo.Create(ctx, cart)
		if err != nil {
			s.l.Errorf("create cart fail, error %s", err.Error())
			return nil, fmt.Errorf("create cart fail, error %w", err)
		}

		return cart, nil
	}

	err = s.repo.Update(ctx, cart)
	if err != nil {
		s.l.Errorf("update cart fail, error %s", err.Error())
		return nil, fmt.Errorf("update cart fail, error %w", err)
	}

	return cart, nil
}
