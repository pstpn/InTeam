package handlers

import (
	"backend/internal/dto"
	"backend/internal/model"
	api "backend/internal/router/ogen"
	"backend/internal/service"
	"backend/pkg/common"
	"backend/pkg/logger"
	"context"
	"errors"
)

var _ api.Handler = (*Handler)(nil)

type Handler struct {
	Logger          logger.Interface
	AuthService     service.IAuthService
	CartService     service.ICartService
	UserService     service.IUserService
	FeedbackService service.IFeedbackService
	OrderService    service.IOrderService
}

func NewHandler(
	l logger.Interface,
	authService service.IAuthService,
	cartService service.ICartService,
	userService service.IUserService,
	feedbackService service.IFeedbackService,
	orderService service.IOrderService,
) *Handler {
	return &Handler{
		Logger:          l,
		AuthService:     authService,
		CartService:     cartService,
		UserService:     userService,
		FeedbackService: feedbackService,
		OrderService:    orderService,
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

func (h *Handler) UserAddRacket(ctx context.Context, req *api.UserAddRacketReq) (*api.AddRacketResponse, error) {
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
			Lines:      modelToApiCartLines(cart.Lines),
			Quantity:   cart.Quantity,
			TotalPrice: cart.TotalPrice,
		},
	}, nil
}

func (h *Handler) UserCreateFeedback(ctx context.Context, req *api.UserCreateFeedbackReq) (*api.CreateFeedbackResponse, error) {
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

func (h *Handler) UserCreateOrder(ctx context.Context, req *api.UserCreateOrderReq) error {
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

func (h *Handler) UserDeleteFeedback(ctx context.Context, params api.UserDeleteFeedbackParams) error {
	userID := common.MustUserIDFromCtx(ctx)

	if err := h.FeedbackService.RemoveFeedback(ctx, &dto.RemoveFeedbackReq{
		RacketID: params.RacketID,
		UserID:   userID,
	}); err != nil {
		return ErrBadRequest
	}

	return nil
}

func (h *Handler) UserDeleteRacket(ctx context.Context, params api.UserDeleteRacketParams) (*api.DeleteRacketResponse, error) {
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
			Lines:      modelToApiCartLines(cart.Lines),
			Quantity:   cart.Quantity,
			TotalPrice: cart.TotalPrice,
		},
	}, nil
}

func (h *Handler) UserGetCart(ctx context.Context) (*api.GetCartResponse, error) {
	userID := common.MustUserIDFromCtx(ctx)

	cart, err := h.CartService.GetCartByID(ctx, userID)
	if err != nil {
		return nil, ErrBadRequest
	}

	return &api.GetCartResponse{
		Cart: api.Cart{
			Lines:      modelToApiCartLines(cart.Lines),
			Quantity:   cart.Quantity,
			TotalPrice: cart.TotalPrice,
		},
	}, nil
}

func (h *Handler) UserGetFeedbacks(ctx context.Context) (*api.GetFeedbacksResponse, error) {
	userID := common.MustUserIDFromCtx(ctx)

	feedbacks, err := h.FeedbackService.GetFeedbacksByUserID(ctx, userID)
	if err != nil {
		return nil, ErrBadRequest
	}

	return &api.GetFeedbacksResponse{Feedbacks: modelToApiFeedbacks(feedbacks)}, nil
}

func (h *Handler) UserGetProfile(ctx context.Context) (*api.GetProfileResponse, error) {
	userID := common.MustUserIDFromCtx(ctx)

	user, err := h.UserService.GetUserByID(ctx, userID)
	if err != nil {
		return nil, ErrBadRequest
	}

	return &api.GetProfileResponse{
		User: api.User{
			Email:   user.Email,
			Name:    user.Name,
			Role:    string(user.Role),
			Surname: user.Surname,
		},
	}, nil
}

func (h *Handler) UserUpdateRacketsCount(ctx context.Context, req *api.UserUpdateRacketsCountReq, params api.UserUpdateRacketsCountParams) (*api.UpdateRacketsCountResponse, error) {
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
			Lines:      modelToApiCartLines(cart.Lines),
			Quantity:   cart.Quantity,
			TotalPrice: cart.TotalPrice,
		},
	}, nil
}

func (h *Handler) NewError(_ context.Context, err error) *api.ErrorResponseStatusCode {
	h.Logger.Errorf("unknown error: %s", err.Error())

	apiErr := api.ErrorResponseStatusCode{}
	if errors.Is(err, &apiErr) {
		return &apiErr
	}

	return ErrInternal
}

func modelToApiCartLines(cartLines []*model.CartLine) []api.CartLine {
	apiCartLines := make([]api.CartLine, 0, len(cartLines))
	for _, line := range cartLines {
		apiCartLines = append(apiCartLines, api.CartLine{
			RacketID: line.RacketID,
			Price:    line.Price,
			Quantity: line.Quantity,
		})
	}
	return apiCartLines
}

func modelToApiFeedbacks(feedbacks []*model.Feedback) []api.Feedback {
	apiFeedbacks := make([]api.Feedback, 0, len(feedbacks))
	for _, feedback := range feedbacks {
		apiFeedbacks = append(apiFeedbacks, api.Feedback{
			Date:     feedback.Date,
			Feedback: feedback.Feedback,
			RacketID: feedback.RacketID,
			Rating:   feedback.Rating,
		})
	}
	return apiFeedbacks
}
