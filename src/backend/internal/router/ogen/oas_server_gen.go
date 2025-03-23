// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"
)

// Handler handles operations described by OpenAPI v3 specification.
type Handler interface {
	// AuthLogin implements Auth_login operation.
	//
	// Login user.
	//
	// POST /api/auth/login
	AuthLogin(ctx context.Context, req *AuthLoginReq) (*LoginResponse, error)
	// AuthRegister implements Auth_register operation.
	//
	// Register user.
	//
	// POST /api/auth/register
	AuthRegister(ctx context.Context, req *AuthRegisterReq) (*RegisterResponse, error)
	// CartAddRacket implements Cart_addRacket operation.
	//
	// Add racket in cart.
	//
	// POST /api/cart
	CartAddRacket(ctx context.Context, req *CartAddRacketReq) (*AddRacketResponse, error)
	// CartGetCart implements Cart_getCart operation.
	//
	// Get cart items.
	//
	// GET /api/cart
	CartGetCart(ctx context.Context) (*GetCartResponse, error)
	// FeedbacksCreateFeedback implements Feedbacks_createFeedback operation.
	//
	// Create feedback.
	//
	// POST /api/feedbacks
	FeedbacksCreateFeedback(ctx context.Context, req *FeedbacksCreateFeedbackReq) (*CreateFeedbackResponse, error)
	// FeedbacksDeleteFeedback implements Feedbacks_deleteFeedback operation.
	//
	// Delete feedback.
	//
	// DELETE /api/feedbacks/{racket_id}
	FeedbacksDeleteFeedback(ctx context.Context, params FeedbacksDeleteFeedbackParams) error
	// FeedbacksGetFeedbacks implements Feedbacks_getFeedbacks operation.
	//
	// Get user feedbacks.
	//
	// GET /api/feedbacks
	FeedbacksGetFeedbacks(ctx context.Context) (*GetFeedbacksResponse, error)
	// OrdersCreateOrder implements Orders_createOrder operation.
	//
	// Create order.
	//
	// POST /api/orders
	OrdersCreateOrder(ctx context.Context, req *OrdersCreateOrderReq) error
	// OrdersGetOrder implements Orders_getOrder operation.
	//
	// Get order.
	//
	// GET /api/orders/{order_id}
	OrdersGetOrder(ctx context.Context, params OrdersGetOrderParams) (*GetOrderResponse, error)
	// OrdersGetOrders implements Orders_getOrders operation.
	//
	// Get orders.
	//
	// GET /api/orders
	OrdersGetOrders(ctx context.Context, params OrdersGetOrdersParams) (*GetOrdersResponse, error)
	// OrdersUpdateOrderStatus implements Orders_updateOrderStatus operation.
	//
	// Update order status.
	//
	// PATCH /api/orders/{order_id}
	OrdersUpdateOrderStatus(ctx context.Context, req *OrdersUpdateOrderStatusReq, params OrdersUpdateOrderStatusParams) error
	// ProfileGetProfile implements Profile_getProfile operation.
	//
	// Get user profile.
	//
	// GET /api/profile
	ProfileGetProfile(ctx context.Context) (*GetProfileResponse, error)
	// RacketsCreateRacket implements Rackets_createRacket operation.
	//
	// Create racket in shop.
	//
	// POST /api/rackets
	RacketsCreateRacket(ctx context.Context, req *RacketsCreateRacketReq) (*CreateRacketResponse, error)
	// RacketsDeleteRacket implements Rackets_deleteRacket operation.
	//
	// Delete racket from cart.
	//
	// DELETE /api/cart/rackets/{racket_id}
	RacketsDeleteRacket(ctx context.Context, params RacketsDeleteRacketParams) (*DeleteRacketResponse, error)
	// RacketsGetRacket implements Rackets_getRacket operation.
	//
	// Get racket.
	//
	// GET /api/rackets/{racket_id}
	RacketsGetRacket(ctx context.Context, params RacketsGetRacketParams) (*GetRacketResponse, error)
	// RacketsGetRacketFeedbacks implements Rackets_getRacketFeedbacks operation.
	//
	// Get feedbacks for racket.
	//
	// GET /api/feedbacks/rackets/{racket_id}
	RacketsGetRacketFeedbacks(ctx context.Context, params RacketsGetRacketFeedbacksParams) (*GetRacketFeedbacksResponse, error)
	// RacketsGetRackets implements Rackets_getRackets operation.
	//
	// Get all rackets.
	//
	// GET /api/rackets
	RacketsGetRackets(ctx context.Context, params RacketsGetRacketsParams) (*GetRacketsResponse, error)
	// RacketsUpdateRacketQuantity implements Rackets_updateRacketQuantity operation.
	//
	// Update racket quantity.
	//
	// PATCH /api/rackets/{racket_id}
	RacketsUpdateRacketQuantity(ctx context.Context, req *RacketsUpdateRacketQuantityReq, params RacketsUpdateRacketQuantityParams) error
	// RacketsUpdateRacketsCount implements Rackets_updateRacketsCount operation.
	//
	// Update rackets count in cart.
	//
	// PUT /api/cart/rackets/{racket_id}
	RacketsUpdateRacketsCount(ctx context.Context, req *RacketsUpdateRacketsCountReq, params RacketsUpdateRacketsCountParams) (*UpdateRacketsCountResponse, error)
	// UsersChangeUserRole implements Users_changeUserRole operation.
	//
	// Change user role.
	//
	// PATCH /api/users/{user_id}
	UsersChangeUserRole(ctx context.Context, req *UsersChangeUserRoleReq, params UsersChangeUserRoleParams) error
	// UsersGetUser implements Users_getUser operation.
	//
	// Get user.
	//
	// GET /api/users/{user_id}
	UsersGetUser(ctx context.Context, params UsersGetUserParams) (*GetUserResponse, error)
	// UsersGetUsers implements Users_getUsers operation.
	//
	// Get all users.
	//
	// GET /api/users
	UsersGetUsers(ctx context.Context) (*GetUsersResponse, error)
	// NewError creates *ErrorResponseStatusCode from error returned by handler.
	//
	// Used for common default response.
	NewError(ctx context.Context, err error) *ErrorResponseStatusCode
}

// Server implements http server based on OpenAPI v3 specification and
// calls Handler to handle requests.
type Server struct {
	h   Handler
	sec SecurityHandler
	baseServer
}

// NewServer creates new Server.
func NewServer(h Handler, sec SecurityHandler, opts ...ServerOption) (*Server, error) {
	s, err := newServerConfig(opts...).baseServer()
	if err != nil {
		return nil, err
	}
	return &Server{
		h:          h,
		sec:        sec,
		baseServer: s,
	}, nil
}
