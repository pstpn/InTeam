// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"

	ht "github.com/ogen-go/ogen/http"
)

// UnimplementedHandler is no-op Handler which returns http.ErrNotImplemented.
type UnimplementedHandler struct{}

var _ Handler = UnimplementedHandler{}

// AuthLogin implements Auth_login operation.
//
// Login user.
//
// POST /api/auth/login
func (UnimplementedHandler) AuthLogin(ctx context.Context, req *AuthLoginReq) (r *LoginResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// AuthRegister implements Auth_register operation.
//
// Register user.
//
// POST /api/auth/register
func (UnimplementedHandler) AuthRegister(ctx context.Context, req *AuthRegisterReq) (r *RegisterResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// UserAddRacket implements User_addRacket operation.
//
// Add racket in cart.
//
// POST /api/cart
func (UnimplementedHandler) UserAddRacket(ctx context.Context, req *UserAddRacketReq) (r *AddRacketResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// UserCreateFeedback implements User_createFeedback operation.
//
// Create feedback.
//
// POST /api/feedbacks
func (UnimplementedHandler) UserCreateFeedback(ctx context.Context, req *UserCreateFeedbackReq) (r *CreateFeedbackResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// UserCreateOrder implements User_createOrder operation.
//
// Create order.
//
// POST /api/orders
func (UnimplementedHandler) UserCreateOrder(ctx context.Context, req *UserCreateOrderReq) error {
	return ht.ErrNotImplemented
}

// UserDeleteFeedback implements User_deleteFeedback operation.
//
// Delete feedback.
//
// DELETE /api/feedbacks/{racket_id}
func (UnimplementedHandler) UserDeleteFeedback(ctx context.Context, params UserDeleteFeedbackParams) error {
	return ht.ErrNotImplemented
}

// UserDeleteRacket implements User_deleteRacket operation.
//
// Delete racket from cart.
//
// DELETE /api/cart/rackets/{racket_id}
func (UnimplementedHandler) UserDeleteRacket(ctx context.Context, params UserDeleteRacketParams) (r *DeleteRacketResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// UserGetCart implements User_getCart operation.
//
// Get cart items.
//
// GET /api/cart
func (UnimplementedHandler) UserGetCart(ctx context.Context) (r *GetCartResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// UserGetFeedbacks implements User_getFeedbacks operation.
//
// Get user feedbacks.
//
// GET /api/feedbacks
func (UnimplementedHandler) UserGetFeedbacks(ctx context.Context) (r *GetFeedbacksResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// UserGetProfile implements User_getProfile operation.
//
// Get user profile.
//
// GET /api/profile
func (UnimplementedHandler) UserGetProfile(ctx context.Context) (r *GetProfileResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// UserUpdateRacketsCount implements User_updateRacketsCount operation.
//
// Update rackets count in cart.
//
// PUT /api/cart/rackets/{racket_id}
func (UnimplementedHandler) UserUpdateRacketsCount(ctx context.Context, req *UserUpdateRacketsCountReq, params UserUpdateRacketsCountParams) (r *UpdateRacketsCountResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// NewError creates *ErrorResponseStatusCode from error returned by handler.
//
// Used for common default response.
func (UnimplementedHandler) NewError(ctx context.Context, err error) (r *ErrorResponseStatusCode) {
	r = new(ErrorResponseStatusCode)
	return r
}
