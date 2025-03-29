package service_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/ovechkin-dm/mockio/mock"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"

	"backend/internal/dto"
	"backend/internal/model"
	"backend/internal/service"
	"backend/internal/storage"
	"backend/pkg/logger"
)

func TestService_Login(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	type args struct {
		req *dto.LoginReq
	}
	tests := []struct {
		name    string
		args    args
		setup   func(s *service.AuthService)
		wantErr bool
	}{
		{
			name: "Failed to get user by email",
			args: args{
				req: &dto.LoginReq{
					Email:    "test",
					Password: "test",
				},
			},
			setup: func(s *service.AuthService) {
				mock.WhenDouble(
					s.UserRepo.GetUserByEmail(mock.AnyContext(), mock.AnyString()),
				).ThenReturn(nil, errors.New("failed"))
			},
			wantErr: true,
		},
		{
			name: "Incorrect password",
			args: args{
				req: &dto.LoginReq{
					Email:    "test",
					Password: "test",
				},
			},
			setup: func(s *service.AuthService) {
				mock.WhenDouble(
					s.UserRepo.GetUserByEmail(mock.AnyContext(), mock.AnyString()),
				).ThenReturn(&model.User{
					ID:       1,
					Name:     "test",
					Surname:  "test",
					Email:    "test",
					Password: "incorrect",
					Role:     "test",
				}, nil)
			},
			wantErr: true,
		},
		{
			name: "Success",
			args: args{
				req: &dto.LoginReq{
					Email:    "test",
					Password: "test",
				},
			},
			setup: func(s *service.AuthService) {
				password, _ := bcrypt.GenerateFromPassword([]byte("test"), bcrypt.MinCost)
				mock.WhenDouble(
					s.UserRepo.GetUserByEmail(mock.AnyContext(), mock.AnyString()),
				).ThenReturn(&model.User{
					ID:       1,
					Name:     "test",
					Surname:  "test",
					Email:    "test",
					Password: string(password),
					Role:     "test",
				}, nil)
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			mock.SetUp(t)
			s := service.NewAuthService(
				mock.Mock[logger.Interface](),
				mock.Mock[storage.IUserStorage](),
				"test",
				time.Second,
			)
			tt.setup(s)
			got, err := s.Login(ctx, tt.args.req)
			if tt.wantErr {
				require.Error(t, err)
				require.Empty(t, got)
			} else {
				require.NoError(t, err)
				require.NotEmpty(t, got)
			}
		})
	}
}

func TestService_Register(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	type args struct {
		req *dto.RegisterReq
	}
	tests := []struct {
		name    string
		args    args
		setup   func(s *service.AuthService)
		wantErr bool
	}{
		{
			name: "Failed to get user by email",
			args: args{
				req: &dto.RegisterReq{
					Name:     "test",
					Surname:  "test",
					Email:    "test",
					Password: "test",
				},
			},
			setup: func(s *service.AuthService) {
				mock.WhenDouble(
					s.UserRepo.GetUserByEmail(mock.AnyContext(), mock.AnyString()),
				).ThenReturn(nil, errors.New("failed"))
			},
			wantErr: true,
		},
		{
			name: "Failed to create user",
			args: args{
				req: &dto.RegisterReq{
					Name:     "test",
					Surname:  "test",
					Email:    "test",
					Password: "test",
				},
			},
			setup: func(s *service.AuthService) {
				mock.WhenDouble(
					s.UserRepo.GetUserByEmail(mock.AnyContext(), mock.AnyString()),
				).ThenReturn(&model.User{
					ID:       1,
					Name:     "test",
					Surname:  "test",
					Email:    "test",
					Password: "test",
					Role:     "test",
				}, storage.ErrNotFound)
				mock.WhenSingle(
					s.UserRepo.Create(mock.AnyContext(), mock.Any[*model.User]()),
				).ThenReturn(errors.New("failed"))
			},
			wantErr: true,
		},
		{
			name: "Success",
			args: args{
				req: &dto.RegisterReq{
					Name:     "test",
					Surname:  "test",
					Email:    "test",
					Password: "test",
				},
			},
			setup: func(s *service.AuthService) {
				mock.WhenDouble(
					s.UserRepo.GetUserByEmail(mock.AnyContext(), mock.AnyString()),
				).ThenReturn(&model.User{
					ID:       1,
					Name:     "test",
					Surname:  "test",
					Email:    "test",
					Password: "test",
					Role:     "test",
				}, storage.ErrNotFound)
				mock.WhenSingle(
					s.UserRepo.Create(mock.AnyContext(), mock.Any[*model.User]()),
				).ThenReturn(nil)
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			mock.SetUp(t)
			s := service.NewAuthService(
				mock.Mock[logger.Interface](),
				mock.Mock[storage.IUserStorage](),
				"test",
				time.Second,
			)
			tt.setup(s)
			got, err := s.Register(ctx, tt.args.req)
			if tt.wantErr {
				require.Error(t, err)
				require.Empty(t, got)
			} else {
				require.NoError(t, err)
				require.NotEmpty(t, got)
			}
		})
	}
}
