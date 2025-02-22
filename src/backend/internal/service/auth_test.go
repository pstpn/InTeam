package service_test

import (
	"backend/internal/dto"
	"backend/internal/model"
	"backend/internal/service"
	"backend/internal/storage"
	"backend/pkg/logger"
	"context"
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
	"time"

	"github.com/ovechkin-dm/mockio/mock"
)

func TestService_CertificatesGetCertificate(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	type args struct {
		req *dto.LoginReq
	}
	tests := []struct {
		name    string
		args    args
		setup   func(s *service.AuthService)
		want    string
		wantErr bool
	}{
		{
			name: "Failed to get by email",
			args: args{
				req: &dto.LoginReq{
					Email:    "test",
					Password: "test",
				},
			},
			setup: func(s *service.AuthService) {
				mock.WhenDouble(
					s.UserRepo.GetUserByEmail(mock.AnyContext(), mock.AnyString()),
				).ThenReturn(nil, fmt.Errorf("failed"))
			},
			want:    "",
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
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			mock.SetUp(t)
			s := service.NewAuthService(
				mock.Mock[logger.Interface](),
				mock.Mock[storage.IUserRepository](),
				"",
				time.Second,
			)
			tt.setup(s)
			got, err := s.Login(ctx, tt.args.req)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
			require.Equal(t, tt.want, got)
		})
	}
}
