package handlers

import (
	"backend/internal/model"
	api "backend/internal/router/ogen"
)

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

func modelToApiOrderLines(orderLines []*model.OrderLine) []api.OrderLine {
	apiOrderLines := make([]api.OrderLine, 0, len(orderLines))
	for _, line := range orderLines {
		apiOrderLines = append(apiOrderLines, api.OrderLine{
			RacketID: line.RacketID,
			Quantity: line.Quantity,
		})
	}
	return apiOrderLines
}

func modelToApiOrder(order *model.Order) api.Order {
	return api.Order{
		ID:            order.ID,
		UserID:        order.UserID,
		CreationDate:  order.CreationDate,
		DeliveryDate:  order.DeliveryDate,
		Address:       order.Address,
		RecipientName: order.RecepientName,
		Status:        string(order.Status),
		Lines:         modelToApiOrderLines(order.Lines),
		TotalPrice:    order.TotalPrice,
	}
}

func modelToApiOrders(orders []*model.Order) []api.Order {
	apiOrders := make([]api.Order, 0, len(orders))
	for _, order := range orders {
		apiOrders = append(apiOrders, modelToApiOrder(order))
	}
	return apiOrders
}

func modelToApiRacket(racket *model.Racket) api.Racket {
	return api.Racket{
		ID:        racket.ID,
		Brand:     racket.Brand,
		Weight:    racket.Weight,
		Balance:   racket.Balance,
		HeadSize:  racket.HeadSize,
		Available: racket.Available,
		Quantity:  racket.Quantity,
		Price:     racket.Price,
		Image:     racket.Image,
	}
}

func modelToApiRackets(rackets []*model.Racket) []api.Racket {
	apiRackets := make([]api.Racket, 0, len(rackets))
	for _, racket := range rackets {
		apiRackets = append(apiRackets, modelToApiRacket(racket))
	}
	return apiRackets
}

func modelToApiUser(user *model.User) api.User {
	return api.User{
		Email:   user.Email,
		Name:    user.Name,
		Role:    string(user.Role),
		Surname: user.Surname,
	}
}

func modelToApiUsers(users []*model.User) []api.User {
	apiUsers := make([]api.User, 0, len(users))
	for _, user := range users {
		apiUsers = append(apiUsers, modelToApiUser(user))
	}
	return apiUsers
}
