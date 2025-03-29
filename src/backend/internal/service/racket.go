package service

import (
	"context"
	"errors"
	"fmt"

	"backend/internal/dto"
	"backend/internal/model"
	"backend/internal/storage"
	"backend/pkg/common"
	"backend/pkg/logger"
)

type IRacketService interface {
	CreateRacket(ctx context.Context, req *dto.CreateRacketReq) (*model.Racket, error)
	UpdateRacket(ctx context.Context, req *dto.UpdateRacketReq) error
	GetRacketByID(ctx context.Context, id int) (*model.Racket, error)
	GetAllRackets(ctx context.Context, req *dto.ListRacketsReq) ([]*model.Racket, error)
}

type RacketService struct {
	l    logger.Interface
	repo storage.IRacketStorage
}

func NewRacketService(l logger.Interface, repo storage.IRacketStorage) *RacketService {
	return &RacketService{
		l:    l,
		repo: repo,
	}
}

func (s *RacketService) CreateRacket(ctx context.Context, req *dto.CreateRacketReq) (*model.Racket, error) {
	var racket model.Racket

	common.Copy(&racket, req)

	switch {
	case racket.Quantity > 0:
		racket.Available = true
	case racket.Quantity == 0:
		racket.Available = false
	default:
		s.l.Errorf("unavailable amount of rackets, error")
		return nil, errors.New("unavailable amount of rackets, error")
	}

	err := s.repo.Create(ctx, &racket)
	if err != nil {
		s.l.Errorf("create fail, error %s", err.Error())
		return nil, fmt.Errorf("create fail, error %w", err)
	}

	return &racket, nil
}

func (s *RacketService) UpdateRacket(ctx context.Context, req *dto.UpdateRacketReq) error {
	racket, err := s.repo.GetRacketByID(ctx, req.ID)
	if racket == nil {
		s.l.Errorf("get racket by id, error %s", err.Error())
		return fmt.Errorf("get racket by id, error %w", err)
	}

	if req.Quantity < 0 {
		s.l.Errorf("unavailbale amount of rackets")
		return errors.New("unavailable amount of rackets")
	}

	racket.Quantity = req.Quantity

	err = s.repo.Update(ctx, racket)
	if err != nil {
		s.l.Errorf("update racket fail, error %s", err.Error())
		return fmt.Errorf("update racket fail, error %w", err)
	}

	return nil
}

func (s *RacketService) GetRacketByID(ctx context.Context, id int) (*model.Racket, error) {
	racket, err := s.repo.GetRacketByID(ctx, id)
	if err != nil {
		s.l.Errorf("get racket by id fail, error %s", err.Error())
		return nil, fmt.Errorf("get racket by id fail, error %w", err)
	}

	return racket, nil
}

func (s *RacketService) GetAllRackets(ctx context.Context, req *dto.ListRacketsReq) ([]*model.Racket, error) {
	rackets, err := s.repo.GetAllRackets(ctx, req)
	if err != nil {
		s.l.Errorf("get all fail, error %s", err.Error())
		return nil, fmt.Errorf("get all fail, error %w", err)
	}

	return rackets, nil
}
