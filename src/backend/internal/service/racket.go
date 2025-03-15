package service

import (
	"context"
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
	logger logger.Interface
	repo   storage.IRacketStorage
}

func NewRacketService(logger logger.Interface, repo storage.IRacketStorage) *RacketService {
	return &RacketService{
		logger: logger,
		repo:   repo,
	}
}

func (s *RacketService) CreateRacket(ctx context.Context, req *dto.CreateRacketReq) (*model.Racket, error) {
	var racket model.Racket

	common.Copy(&racket, req)

	if racket.Quantity > 0 {
		racket.Available = true
	} else if racket.Quantity == 0 {
		racket.Available = false
	} else {
		s.logger.Errorf("unavailable amount of rackets, error")
		return nil, fmt.Errorf("unavailable amount of rackets, error")
	}

	err := s.repo.Create(ctx, &racket)
	if err != nil {
		s.logger.Errorf("create fail, error %s", err.Error())
		return nil, fmt.Errorf("create fail, error %s", err)
	}

	return &racket, nil

}

func (s *RacketService) UpdateRacket(ctx context.Context, req *dto.UpdateRacketReq) error {
	racket, err := s.repo.GetRacketByID(ctx, req.ID)
	if racket == nil {
		s.logger.Errorf("get racket by id, error %s", err.Error())
		return fmt.Errorf("get racket by id, error %s", err)
	}

	if req.Quantity < 0 {
		s.logger.Errorf("unavailbale amount of rackets")
		return fmt.Errorf("unavailbale amount of rackets")
	}

	racket.Quantity = req.Quantity

	err = s.repo.Update(ctx, racket)
	if err != nil {
		s.logger.Errorf("update racket fail, error %s", err.Error())
		return fmt.Errorf("update racket fail, error %s", err)
	}

	return nil
}

func (s *RacketService) GetRacketByID(ctx context.Context, id int) (*model.Racket, error) {
	racket, err := s.repo.GetRacketByID(ctx, id)
	if err != nil {
		s.logger.Errorf("get racket by id fail, error %s", err.Error())
		return nil, fmt.Errorf("get racket by id fail, error %s", err)
	}

	return racket, nil
}

func (s *RacketService) GetAllRackets(ctx context.Context, req *dto.ListRacketsReq) ([]*model.Racket, error) {
	rackets, err := s.repo.GetAllRackets(ctx, req)
	if err != nil {
		s.logger.Errorf("get all fail, error %s", err.Error())
		return nil, fmt.Errorf("get all fail, error %s", err)
	}

	return rackets, nil
}
