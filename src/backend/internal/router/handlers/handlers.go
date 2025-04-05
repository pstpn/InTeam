package handlers

import (
	"context"
	"errors"

	"backend/internal/dto"
	"backend/internal/model"
	api "backend/internal/router/ogen"
	"backend/internal/service"
	"backend/pkg/common"
	"backend/pkg/logger"
	"backend/pkg/storage/postgres"
)

var _ api.Handler = (*Handler)(nil)

type Handler struct {
	Logger          logger.Interface
	AuthService     service.IAuthService
	CartService     service.ICartService
	UserService     service.IUserService
	FeedbackService service.IFeedbackService
	OrderService    service.IOrderService
	RacketService   service.IRacketService
}

func NewHandler(
	l logger.Interface,
	authService service.IAuthService,
	cartService service.ICartService,
	userService service.IUserService,
	feedbackService service.IFeedbackService,
	orderService service.IOrderService,
	racketService service.IRacketService,
) *Handler {
	return &Handler{
		Logger:          l,
		AuthService:     authService,
		CartService:     cartService,
		UserService:     userService,
		FeedbackService: feedbackService,
		OrderService:    orderService,
		RacketService:   racketService,
	}
}

func (h *Handler) AuthRegister(ctx context.Context, req *api.AuthRegisterReq) (*api.RegisterResponse, error) {
	token, err := h.AuthService.Register(ctx, &dto.RegisterReq{
		Name:     req.Name,
		Surname:  req.Surname,
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		h.Logger.Infof("failed to register user: %s", err.Error())
		return nil, ErrBadRequest
	}

	return &api.RegisterResponse{AccessToken: token}, nil
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

func (h *Handler) CartAddRacket(ctx context.Context, req *api.CartAddRacketReq) (*api.AddRacketResponse, error) {
	userID := common.MustUserIDFromCtx(ctx)

	cart, err := h.CartService.AddRacket(ctx, &dto.AddRacketCartReq{
		UserID:   userID,
		RacketID: req.RacketID,
		Quantity: req.Quantity,
	})
	if err != nil {
		return nil, ErrBadRequest
	}

	return &api.AddRacketResponse{
		Cart: api.Cart{
			Lines:      modelToAPICartLines(cart.Lines),
			Quantity:   cart.Quantity,
			TotalPrice: cart.TotalPrice,
		},
	}, nil
}

func (h *Handler) FeedbacksCreateFeedback(ctx context.Context, req *api.FeedbacksCreateFeedbackReq) (*api.CreateFeedbackResponse, error) {
	userID := common.MustUserIDFromCtx(ctx)

	feedback, err := h.FeedbackService.CreateFeedback(ctx, &dto.CreateFeedbackReq{
		RacketID: req.RacketID,
		UserID:   userID,
		Feedback: req.Feedback,
		Rating:   req.Rating,
	})
	if err != nil {
		return nil, ErrBadRequest
	}

	return &api.CreateFeedbackResponse{
		Feedback: api.Feedback{
			Date:     feedback.Date,
			Feedback: feedback.Feedback,
			RacketID: feedback.RacketID,
			Rating:   feedback.Rating,
		},
	}, nil
}

func (h *Handler) OrdersCreateOrder(ctx context.Context, req *api.OrdersCreateOrderReq) error {
	userID := common.MustUserIDFromCtx(ctx)

	if err := h.OrderService.CreateOrder(ctx, &dto.PlaceOrderReq{
		UserID:        userID,
		DeliveryDate:  req.DeliveryDate,
		Address:       req.Address,
		RecepientName: req.RecipientName,
	}); err != nil {
		return ErrBadRequest
	}

	return nil
}

func (h *Handler) FeedbacksDeleteFeedback(ctx context.Context, params api.FeedbacksDeleteFeedbackParams) error {
	userID := common.MustUserIDFromCtx(ctx)

	if err := h.FeedbackService.RemoveFeedback(ctx, &dto.RemoveFeedbackReq{
		RacketID: params.RacketID,
		UserID:   userID,
	}); err != nil {
		return ErrBadRequest
	}

	return nil
}

func (h *Handler) RacketsDeleteRacket(ctx context.Context, params api.RacketsDeleteRacketParams) (*api.DeleteRacketResponse, error) {
	userID := common.MustUserIDFromCtx(ctx)

	cart, err := h.CartService.RemoveRacket(ctx, &dto.RemoveRacketCartReq{
		UserID:   userID,
		RacketID: params.RacketID,
	})
	if err != nil {
		return nil, ErrBadRequest
	}

	return &api.DeleteRacketResponse{
		Cart: api.Cart{
			Lines:      modelToAPICartLines(cart.Lines),
			Quantity:   cart.Quantity,
			TotalPrice: cart.TotalPrice,
		},
	}, nil
}

func (h *Handler) CartGetCart(ctx context.Context) (*api.GetCartResponse, error) {
	userID := common.MustUserIDFromCtx(ctx)

	cart, err := h.CartService.GetCartByID(ctx, userID)
	if err != nil {
		return nil, ErrBadRequest
	}

	return &api.GetCartResponse{
		Cart: api.Cart{
			Lines:      modelToAPICartLines(cart.Lines),
			Quantity:   cart.Quantity,
			TotalPrice: cart.TotalPrice,
		},
	}, nil
}

func (h *Handler) FeedbacksGetFeedbacks(ctx context.Context) (*api.GetFeedbacksResponse, error) {
	userID := common.MustUserIDFromCtx(ctx)

	feedbacks, err := h.FeedbackService.GetFeedbacksByUserID(ctx, userID)
	if err != nil {
		return nil, ErrBadRequest
	}

	return &api.GetFeedbacksResponse{Feedbacks: modelToAPIFeedbacks(feedbacks)}, nil
}

func (h *Handler) ProfileGetProfile(ctx context.Context) (*api.GetProfileResponse, error) {
	return &api.GetProfileResponse{User: modelToAPIUser(MustUserFromCtx(ctx))}, nil
}

func (h *Handler) RacketsUpdateRacketsCount(
	ctx context.Context,
	req *api.RacketsUpdateRacketsCountReq,
	params api.RacketsUpdateRacketsCountParams,
) (*api.UpdateRacketsCountResponse, error) {
	userID := common.MustUserIDFromCtx(ctx)

	cart, err := h.CartService.UpdateRacket(ctx, &dto.UpdateRacketCartReq{
		UserID:   userID,
		RacketID: params.RacketID,
		Quantity: req.Quantity,
	})
	if err != nil {
		return nil, ErrBadRequest
	}

	return &api.UpdateRacketsCountResponse{
		Cart: api.Cart{
			Lines:      modelToAPICartLines(cart.Lines),
			Quantity:   cart.Quantity,
			TotalPrice: cart.TotalPrice,
		},
	}, nil
}

func (h *Handler) RacketsCreateRacket(ctx context.Context, req *api.RacketsCreateRacketReq) (*api.CreateRacketResponse, error) {
	imageBytes := make([]byte, req.Image.Size)
	_, err := req.Image.File.Read(imageBytes)
	if err != nil {
		return nil, ErrBadRequest
	}

	racket, err := h.RacketService.CreateRacket(ctx, &dto.CreateRacketReq{
		Brand:    req.Brand,
		Weight:   req.Weight,
		Balance:  req.Balance,
		HeadSize: req.HeadSize,
		Quantity: int(req.Quantity),
		Price:    req.Price,
		Image:    imageBytes,
	})
	if err != nil {
		return nil, ErrBadRequest
	}

	return &api.CreateRacketResponse{Racket: modelToAPIRacket(racket)}, nil
}

func (h *Handler) OrdersGetOrder(ctx context.Context, params api.OrdersGetOrderParams) (*api.GetOrderResponse, error) {
	order, err := h.OrderService.GetOrderByID(ctx, params.OrderID)
	if err != nil {
		return nil, ErrBadRequest
	}

	return &api.GetOrderResponse{Order: modelToAPIOrder(order)}, nil
}

func (h *Handler) OrdersGetOrders(ctx context.Context, params api.OrdersGetOrdersParams) (*api.GetOrdersResponse, error) {
	var orders []*model.Order
	var err error

	p := pagination(params.Pattern, params.Field, params.Sort)

	user := MustUserFromCtx(ctx)
	switch user.Role {
	case model.RoleAdmin:
		orders, err = h.OrderService.GetAllOrders(ctx, &dto.ListOrdersReq{Pagination: p})
	case model.RoleUser:
		orders, err = h.OrderService.GetMyOrders(ctx, &dto.ListOrdersReq{
			UserID:     common.MustUserIDFromCtx(ctx),
			Pagination: p,
		})
	default:
		return nil, ErrInternal
	}

	if err != nil {
		return nil, ErrBadRequest
	}

	return &api.GetOrdersResponse{Orders: modelToAPIOrders(orders)}, nil
}

func (h *Handler) OrdersUpdateOrderStatus(ctx context.Context, req *api.OrdersUpdateOrderStatusReq, params api.OrdersUpdateOrderStatusParams) error {
	if !req.GetStatus().IsSet() {
		return ErrNotModified
	}

	_, err := h.OrderService.UpdateOrderStatus(ctx, &dto.UpdateOrderReq{
		OrderID: params.OrderID,
		Status:  model.OrderStatus(req.Status.Value),
	})
	if err != nil {
		return ErrBadRequest
	}

	return nil
}

func (h *Handler) RacketsGetRacket(ctx context.Context, params api.RacketsGetRacketParams) (*api.GetRacketResponse, error) {
	racket, err := h.RacketService.GetRacketByID(ctx, params.RacketID)
	if err != nil {
		return nil, ErrBadRequest
	}

	return &api.GetRacketResponse{Racket: modelToAPIRacket(racket)}, nil
}

func (h *Handler) RacketsGetRacketFeedbacks(ctx context.Context, params api.RacketsGetRacketFeedbacksParams) (*api.GetRacketFeedbacksResponse, error) {
	feedbacks, err := h.FeedbackService.GetFeedbacksByRacketID(ctx, params.RacketID)
	if err != nil {
		return nil, ErrBadRequest
	}

	return &api.GetRacketFeedbacksResponse{Feedbacks: modelToAPIFeedbacks(feedbacks)}, nil
}

func (h *Handler) RacketsGetRackets(ctx context.Context, params api.RacketsGetRacketsParams) (*api.GetRacketsResponse, error) {
	rackets, err := h.RacketService.GetAllRackets(ctx, &dto.ListRacketsReq{
		Pagination: pagination(params.Pattern, params.Field, params.Sort),
	})
	if err != nil {
		return nil, ErrBadRequest
	}

	return &api.GetRacketsResponse{Rackets: modelToAPIRackets(rackets)}, nil
}

func (h *Handler) RacketsUpdateRacketQuantity(
	ctx context.Context,
	req *api.RacketsUpdateRacketQuantityReq,
	params api.RacketsUpdateRacketQuantityParams,
) error {
	if !req.GetQuantity().IsSet() {
		return ErrNotModified
	}

	err := h.RacketService.UpdateRacket(ctx, &dto.UpdateRacketReq{
		ID:       params.RacketID,
		Quantity: req.Quantity.Value,
	})
	if err != nil {
		return ErrBadRequest
	}

	return nil
}

func (h *Handler) UsersChangeUserRole(ctx context.Context, req *api.UsersChangeUserRoleReq, params api.UsersChangeUserRoleParams) error {
	if !req.GetRole().IsSet() {
		return ErrNotModified
	}

	_, err := h.UserService.UpdateRole(ctx, &dto.UpdateRoleReq{
		ID:   params.UserID,
		Role: model.Role(req.Role.Value),
	})
	if err != nil {
		return ErrBadRequest
	}

	return nil
}

func (h *Handler) UsersGetUser(ctx context.Context, params api.UsersGetUserParams) (*api.GetUserResponse, error) {
	user, err := h.UserService.GetUserByID(ctx, params.UserID)
	if err != nil {
		return nil, ErrBadRequest
	}

	return &api.GetUserResponse{User: modelToAPIUser(user)}, nil
}

func (h *Handler) UsersGetUsers(ctx context.Context) (*api.GetUsersResponse, error) {
	users, err := h.UserService.GetAllUsers(ctx)
	if err != nil {
		return nil, ErrBadRequest
	}

	return &api.GetUsersResponse{Users: modelToAPIUsers(users)}, nil
}

func (h *Handler) NewError(_ context.Context, err error) *api.ErrorResponseStatusCode {
	h.Logger.Errorf("unknown error: %s", err.Error())

	apiErr := api.ErrorResponseStatusCode{}
	if errors.Is(err, &apiErr) {
		return &apiErr
	}

	return ErrInternal
}

func pagination(pattern api.OptString, field api.OptString, sort api.OptString) *postgres.Pagination {
	return &postgres.Pagination{
		Filter: postgres.FilterOptions{
			Pattern: pattern.Or("*"),
			Column:  field.Or(""),
		},
		Sort: postgres.SortOptions{
			Direction: postgres.SortDirectionFromString(sort.Or(postgres.DESC.String())),
			Columns:   []string{field.Or("")},
		},
	}
}
