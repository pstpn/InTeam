package handlers

import (
	"backend/internal/model"
	api "backend/internal/router/ogen"
	"backend/internal/service"
	"backend/pkg/common"
	"context"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
)

var _ api.SecurityHandler = (*SecurityHandler)(nil)

type SecurityHandler struct {
	userService service.IUserService
	signingKey  []byte
}

func NewSecurityHandler(u service.IUserService, signingKey []byte) *SecurityHandler {
	return &SecurityHandler{
		userService: u,
		signingKey:  signingKey,
	}
}

func (s *SecurityHandler) HandleBearerAuth(ctx context.Context, operationName api.OperationName, t api.BearerAuth) (context.Context, error) {
	id, err := s.parseToken(t.GetToken())
	if err != nil {
		return ctx, ErrUnauthorized
	}

	res, err := s.userService.GetUserByID(ctx, id)
	if err != nil {
		return ctx, ErrUnauthorized
	}

	if res.Role == model.RoleUser && isUserMethod(operationName) {
		return common.UserIDToCtx(ctx, id), nil
	} else if res.Role == model.RoleAdmin && isAdminMethod(operationName) {
		return common.AdminIDToCtx(ctx, id), nil
	}

	return ctx, ErrUnauthorized
}

type tokenClaims struct {
	jwt.RegisteredClaims
	UserID int
	Role   string
}

func (s *SecurityHandler) parseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (any, error) {
		_, err := token.Method.(*jwt.SigningMethodHMAC)
		if !err {
			return 0, fmt.Errorf("invalid signing method")
		}

		return s.signingKey, nil
	})
	if err != nil {
		return 0, err
	}

	claims := token.Claims.(*tokenClaims)
	if claims == nil {
		return 0, fmt.Errorf("token claims are not of type *tokenClaims")
	}

	return claims.UserID, nil
}

func isUserMethod(op api.OperationName) bool {
	_, ok := map[api.OperationName]struct{}{
		api.UserAddRacketOperation:          {},
		api.UserCreateFeedbackOperation:     {},
		api.UserCreateOrderOperation:        {},
		api.UserDeleteFeedbackOperation:     {},
		api.UserDeleteRacketOperation:       {},
		api.UserGetCartOperation:            {},
		api.UserGetFeedbacksOperation:       {},
		api.UserGetProfileOperation:         {},
		api.UserUpdateRacketsCountOperation: {},
	}[op]
	return ok
}

func isAdminMethod(op api.OperationName) bool {
	// TODO
	return false
}
