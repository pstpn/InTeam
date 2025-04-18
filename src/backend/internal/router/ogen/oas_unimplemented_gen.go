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

// CartAddRacket implements Cart_addRacket operation.
//
// Add racket in cart.
//
// POST /api/cart
func (UnimplementedHandler) CartAddRacket(ctx context.Context, req *CartAddRacketReq) (r *AddRacketResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// CartGetCart implements Cart_getCart operation.
//
// Get cart items.
//
// GET /api/cart
func (UnimplementedHandler) CartGetCart(ctx context.Context) (r *GetCartResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// FeedbacksCreateFeedback implements Feedbacks_createFeedback operation.
//
// Create feedback.
//
// POST /api/feedbacks
func (UnimplementedHandler) FeedbacksCreateFeedback(ctx context.Context, req *FeedbacksCreateFeedbackReq) (r *CreateFeedbackResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// FeedbacksDeleteFeedback implements Feedbacks_deleteFeedback operation.
//
// Delete feedback.
//
// DELETE /api/feedbacks/{racket_id}
func (UnimplementedHandler) FeedbacksDeleteFeedback(ctx context.Context, params FeedbacksDeleteFeedbackParams) error {
	return ht.ErrNotImplemented
}

// FeedbacksGetFeedbacks implements Feedbacks_getFeedbacks operation.
//
// Get user feedbacks.
//
// GET /api/feedbacks
func (UnimplementedHandler) FeedbacksGetFeedbacks(ctx context.Context) (r *GetFeedbacksResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// OrdersCreateOrder implements Orders_createOrder operation.
//
// Create order.
//
// POST /api/orders
func (UnimplementedHandler) OrdersCreateOrder(ctx context.Context, req *OrdersCreateOrderReq) error {
	return ht.ErrNotImplemented
}

// OrdersGetOrder implements Orders_getOrder operation.
//
// Get order.
//
// GET /api/orders/{order_id}
func (UnimplementedHandler) OrdersGetOrder(ctx context.Context, params OrdersGetOrderParams) (r *GetOrderResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// OrdersGetOrders implements Orders_getOrders operation.
//
// Get orders.
//
// GET /api/orders
func (UnimplementedHandler) OrdersGetOrders(ctx context.Context, params OrdersGetOrdersParams) (r *GetOrdersResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// OrdersUpdateOrderStatus implements Orders_updateOrderStatus operation.
//
// Update order status.
//
// PATCH /api/orders/{order_id}
func (UnimplementedHandler) OrdersUpdateOrderStatus(ctx context.Context, req *OrdersUpdateOrderStatusReq, params OrdersUpdateOrderStatusParams) error {
	return ht.ErrNotImplemented
}

// ProfileGetProfile implements Profile_getProfile operation.
//
// Get profile.
//
// GET /api/profile
func (UnimplementedHandler) ProfileGetProfile(ctx context.Context) (r *GetProfileResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// RacketsCreateRacket implements Rackets_createRacket operation.
//
// Create racket in shop.
//
// POST /api/rackets
func (UnimplementedHandler) RacketsCreateRacket(ctx context.Context, req *RacketsCreateRacketReq) (r *CreateRacketResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// RacketsDeleteRacket implements Rackets_deleteRacket operation.
//
// Delete racket from cart.
//
// DELETE /api/cart/rackets/{racket_id}
func (UnimplementedHandler) RacketsDeleteRacket(ctx context.Context, params RacketsDeleteRacketParams) (r *DeleteRacketResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// RacketsGetRacket implements Rackets_getRacket operation.
//
// Get racket.
//
// GET /api/rackets/{racket_id}
func (UnimplementedHandler) RacketsGetRacket(ctx context.Context, params RacketsGetRacketParams) (r *GetRacketResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// RacketsGetRacketFeedbacks implements Rackets_getRacketFeedbacks operation.
//
// Get feedbacks for racket.
//
// GET /api/feedbacks/rackets/{racket_id}
func (UnimplementedHandler) RacketsGetRacketFeedbacks(ctx context.Context, params RacketsGetRacketFeedbacksParams) (r *GetRacketFeedbacksResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// RacketsGetRackets implements Rackets_getRackets operation.
//
// Get all rackets.
//
// GET /api/rackets
func (UnimplementedHandler) RacketsGetRackets(ctx context.Context, params RacketsGetRacketsParams) (r *GetRacketsResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// RacketsUpdateRacketQuantity implements Rackets_updateRacketQuantity operation.
//
// Update racket quantity.
//
// PATCH /api/rackets/{racket_id}
func (UnimplementedHandler) RacketsUpdateRacketQuantity(ctx context.Context, req *RacketsUpdateRacketQuantityReq, params RacketsUpdateRacketQuantityParams) error {
	return ht.ErrNotImplemented
}

// RacketsUpdateRacketsCount implements Rackets_updateRacketsCount operation.
//
// Update rackets count in cart.
//
// PUT /api/cart/rackets/{racket_id}
func (UnimplementedHandler) RacketsUpdateRacketsCount(ctx context.Context, req *RacketsUpdateRacketsCountReq, params RacketsUpdateRacketsCountParams) (r *UpdateRacketsCountResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// UsersChangeUserRole implements Users_changeUserRole operation.
//
// Change user role.
//
// PATCH /api/users/{user_id}
func (UnimplementedHandler) UsersChangeUserRole(ctx context.Context, req *UsersChangeUserRoleReq, params UsersChangeUserRoleParams) error {
	return ht.ErrNotImplemented
}

// UsersGetUser implements Users_getUser operation.
//
// Get user.
//
// GET /api/users/{user_id}
func (UnimplementedHandler) UsersGetUser(ctx context.Context, params UsersGetUserParams) (r *GetUserResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// UsersGetUsers implements Users_getUsers operation.
//
// Get all users.
//
// GET /api/users
func (UnimplementedHandler) UsersGetUsers(ctx context.Context) (r *GetUsersResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// NewError creates *ErrorResponseStatusCode from error returned by handler.
//
// Used for common default response.
func (UnimplementedHandler) NewError(ctx context.Context, err error) (r *ErrorResponseStatusCode) {
	r = new(ErrorResponseStatusCode)
	return r
}
