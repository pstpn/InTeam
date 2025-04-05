package handlers

import (
	"context"
	"errors"

	"github.com/golang-jwt/jwt/v5"

	"backend/internal/model"
	api "backend/internal/router/ogen"
	"backend/internal/service"
	"backend/pkg/common"
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
		return nil, ErrUnauthorized
	}

	res, err := s.userService.GetUserByID(ctx, id)
	if err != nil {
		return nil, ErrUnauthorized
	}
	ctx = UserToCtx(ctx, res)

	if isAllAuthorizedUserMethod(operationName) {
		return ctx, nil
	}

	if res.Role == model.RoleUser && isUserMethod(operationName) {
		return common.UserIDToCtx(ctx, id), nil
	} else if res.Role == model.RoleAdmin && isAdminMethod(operationName) {
		return common.AdminIDToCtx(ctx, id), nil
	}

	return nil, ErrUnauthorized
}

type tokenClaims struct {
	jwt.RegisteredClaims
	UserID int
	Role   string
}

func (s *SecurityHandler) parseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (any, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("invalid signing method")
		}

		return s.signingKey, nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserID, nil
}

func isUserMethod(op api.OperationName) bool {
	_, ok := map[api.OperationName]struct{}{
		api.CartGetCartOperation:               {},
		api.CartAddRacketOperation:             {},
		api.RacketsDeleteRacketOperation:       {},
		api.RacketsUpdateRacketsCountOperation: {},
		api.FeedbacksGetFeedbacksOperation:     {},
		api.FeedbacksCreateFeedbackOperation:   {},
		api.FeedbacksDeleteFeedbackOperation:   {},
		api.OrdersCreateOrderOperation:         {},
		api.ProfileGetProfileOperation:         {},
	}[op]
	return ok
}

func isAdminMethod(op api.OperationName) bool {
	_, ok := map[api.OperationName]struct{}{
		api.OrdersGetOrderOperation:              {},
		api.OrdersUpdateOrderStatusOperation:     {},
		api.RacketsCreateRacketOperation:         {},
		api.RacketsUpdateRacketQuantityOperation: {},
		api.UsersGetUsersOperation:               {},
		api.UsersGetUserOperation:                {},
		api.UsersChangeUserRoleOperation:         {},
	}[op]
	return ok
}

func isAllAuthorizedUserMethod(op api.OperationName) bool {
	_, ok := map[api.OperationName]struct{}{
		api.OrdersGetOrdersOperation: {},
	}[op]
	return ok
}
