package service

import (
	"context"
	"fmt"

	"backend/internal/dto"
	"backend/internal/model"
	"backend/internal/storage"
	"backend/pkg/logger"
)

type IUserService interface {
	GetUserByID(ctx context.Context, id int) (*model.User, error)
	GetAllUsers(ctx context.Context) ([]*model.User, error)
	UpdateRole(ctx context.Context, req *dto.UpdateRoleReq) (*model.User, error)
}

type UserService struct {
	l    logger.Interface
	repo storage.IUserStorage
}

func NewUserService(
	l logger.Interface,
	repo storage.IUserStorage) *UserService {
	return &UserService{
		l:    l,
		repo: repo,
	}
}

func (s *UserService) UpdateRole(ctx context.Context, req *dto.UpdateRoleReq) (*model.User, error) {
	user, err := s.repo.GetUserByID(ctx, req.ID)
	if err != nil {
		s.l.Errorf("get user by id fail, error %s", err.Error())
		return nil, fmt.Errorf("get user by id fail, error %w", err)
	}

	user.Role = req.Role

	err = s.repo.UpdateRole(ctx, user)
	if err != nil {
		s.l.Errorf("update fail, error %s", err.Error())
		return nil, fmt.Errorf("update fail, error %w", err)
	}

	return user, nil
}

func (s *UserService) GetAllUsers(ctx context.Context) ([]*model.User, error) {
	users, err := s.repo.GetAllUsers(ctx)
	if err != nil {
		s.l.Errorf("get all users fail, error %s", err.Error())
		return nil, fmt.Errorf("get all users fail, error %w", err)
	}

	return users, nil
}

func (s *UserService) GetUserByID(ctx context.Context, id int) (*model.User, error) {
	user, err := s.repo.GetUserByID(ctx, id)
	if err != nil {
		s.l.Errorf("get user by id fail, error %s", err.Error())
		return nil, fmt.Errorf("get user by id fail, error %w", err)
	}

	return user, nil
}
