package handlers

import (
	"backend/internal/dto"
	api "backend/internal/router/ogen"
	"backend/internal/service"
	"backend/pkg/logger"
	"context"
	"errors"
)

var _ api.Handler = (*Handler)(nil)

type Handler struct {
	Logger      logger.Interface
	AuthService service.IAuthService
}

func NewHandler(l logger.Interface, authService service.IAuthService) *Handler {
	return &Handler{
		Logger:      l,
		AuthService: authService,
	}
}

func (h *Handler) AuthLogin(ctx context.Context, req *api.AuthLoginReq) (*api.LoginResponse, error) {
	token, err := h.AuthService.Login(ctx, &dto.LoginReq{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		h.Logger.Infof("failed to login user: %s", err.Error())
		return nil, ErrBadRequest
	}

	return &api.LoginResponse{AccessToken: token}, nil
}

func (h *Handler) NewError(ctx context.Context, err error) *api.ErrorResponseStatusCode {
	apiErr := api.ErrorResponseStatusCode{}
	if errors.Is(err, &apiErr) {
		return &apiErr
	}

	return ErrInternal
}
